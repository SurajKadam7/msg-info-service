package msginfosrv

import (
	"context"
	"errors"

	"github.com/SurajKadam7/msg-info-service/model"
	"github.com/SurajKadam7/msg-info-service/repos/msginfo"
)

// @microgen http, logging, middleware
type Service interface {
	// @http-method POST
	// @http-path msginfo/
	Add(ctx context.Context, msg model.MsgInfo) (id int, err error)
	// @http-method DELETE
	// @http-path msginfo/
	Delete(ctx context.Context, userId int, msgId int) (err error)
	// @http-method POST
	// @http-path msginfo/all/
	GetAll(ctx context.Context, userId int, status model.Status) (msgs []model.MsgInfo, err error)
	// @http-method PUT
	// @http-path msginfo/
	Update(ctx context.Context, userId int, msgId int, status model.Status) (err error)
}

type service struct {
	repo msginfo.Repository
}

func New(repo msginfo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func validateMsg(msg model.MsgInfo) (ok bool) {
	switch {
	case msg.From <= 0:
		return
	case msg.To <= 0:
		return
	case int(msg.Status) < 0:
		return
	default:
		ok = true
	}

	return
}

func (s *service) Add(ctx context.Context, msg model.MsgInfo) (id int, err error) {
	if !validateMsg(msg) {
		return 0, errors.New("not valid msg")
	}
	// TODO add some bussiness logic in case of error
	id, err = s.repo.Add(ctx, msg)
	return
}

func (s *service) Delete(ctx context.Context, userId int, msgId int) (err error) {
	if userId <= 0 || msgId <= 0 {
		return
	}
	err = s.Delete(ctx, userId, msgId)
	return
}

func (s *service) GetAll(ctx context.Context, userId int, status model.Status) (msgs []model.MsgInfo, err error) {
	if userId <= 0 || int(status) < 0 {
		return []model.MsgInfo{}, errors.New("wrong request")
	}
	msgs, err = s.repo.Get(ctx, userId, status)
	return
}

func (c *service) Update(ctx context.Context, userId int, msgId int, status model.Status) (err error) {
	if userId <= 0 || msgId <= 0 || int(status) < 0 {
		return errors.New("wrong request")
	}
	err = c.Update(ctx, userId, msgId, status)
	return
}
