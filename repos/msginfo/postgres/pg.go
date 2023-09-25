package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/suraj.kadam7/msg-info-srv/model"
	"github.com/suraj.kadam7/msg-info-srv/repos/msginfo"
)

type pg struct {
	client       *pgxpool.Pool
	msgInfoTable string
}

func New(client *pgxpool.Pool, msgInfoTable string) msginfo.Repository {
	return &pg{
		client:       client,
		msgInfoTable: msgInfoTable,
	}
}

func (p *pg) Add(ctx context.Context, msg model.MsgInfo) (id int, err error) {
	query := "INSERT INTO $1 (from, to, msg, status) VALUES ($2, $3, $4, $5)"
	args := []any{p.msgInfoTable, msg.From, msg.To, msg.Msg, msg.Status}
	row := p.client.QueryRow(ctx, query, args...)

	msgInfo := model.MsgInfo{}
	if err = row.Scan(&msgInfo); err != nil {
		return
	}

	id = msg.Id
	return
}

// TODO NEED TO DELETE THE MSGS AS WELL
func (p *pg) Delete(ctx context.Context, userId int, msgId int) (err error) {
	query := "DELETE FROM $1 WHERE id=$2 AND from=$3"
	args := []any{p.msgInfoTable, msgId, userId}
	row := p.client.QueryRow(ctx, query, args...)

	err = row.Scan()

	switch err {
	case pgx.ErrNoRows:
		return errors.New("wrong msg delete request")
	default:
		return
	}
}

// TODO NEED TO CONVERT IN BATCH
func (p *pg) Get(ctx context.Context, userId int, status model.Status) (msgs []model.MsgInfo, err error) {
	query := "SELECT id, from, to, msg, status FROM $1 WHERE to=$2 AND status=$3"
	args := []any{p.msgInfoTable, userId, status}
	rows, err := p.client.Query(ctx, query, args...)
	if err != nil {
		return []model.MsgInfo{}, err
	}

	defer rows.Close()

	for {
		msg := model.MsgInfo{}
		rows.Scan(&msg)
		msgs = append(msgs, msg)
		if !rows.Next() {
			break
		}
	}

	err = rows.Err()
	return
}

func (p *pg) Update(ctx context.Context, userId int, msgId int, status model.Status) (err error) {
	query := "UPDATE $1 SET status=$2 WHERE id=$3 AND from=$4"
	args := []any{p.msgInfoTable, status, msgId, userId}
	row := p.client.QueryRow(ctx, query, args...)

	err = row.Scan()
	switch err {
	case pgx.ErrNoRows:
		return errors.New("wrong msg update request")
	default:
		return
	}
}
