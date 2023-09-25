package msginfo

import (
	"context"

	"github.com/suraj.kadam7/msg-info-srv/model"
)

// @microgen logging, middleware
type Repository interface {
	Add(ctx context.Context, msg model.MsgInfo) (id int, err error)
	Delete(ctx context.Context, userId int, msgId int) (err error)
	Get(ctx context.Context, userId int, status model.Status) (msgs []model.MsgInfo, err error)
	Update(ctx context.Context, userId int, msgId int, status model.Status) (err error)
}
