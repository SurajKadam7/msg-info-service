// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	"context"
	model "github.com/SurajKadam7/msg-info-service/model"
	service "github.com/SurajKadam7/msg-info-service/msginfo_srv"
	log "github.com/go-kit/log"
	"time"
)

// LoggingMiddleware writes params, results and working time of method call to provided logger after its execution.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (M loggingMiddleware) Add(arg0 context.Context, arg1 model.MsgInfo) (res0 int, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Add",
			"message", "Add called",
			"request", logAddRequest{Msg: arg1},
			"response", logAddResponse{Id: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Add(arg0, arg1)
}

func (M loggingMiddleware) Delete(arg0 context.Context, arg1 int, arg2 int) (res0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Delete",
			"message", "Delete called",
			"request", logDeleteRequest{
				MsgId:  arg2,
				UserId: arg1,
			},
			"err", res0,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Delete(arg0, arg1, arg2)
}

func (M loggingMiddleware) GetAll(arg0 context.Context, arg1 int, arg2 model.Status) (res0 []model.MsgInfo, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAll",
			"message", "GetAll called",
			"request", logGetAllRequest{
				Status: arg2,
				UserId: arg1,
			},
			"response", logGetAllResponse{Msgs: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAll(arg0, arg1, arg2)
}

func (M loggingMiddleware) Update(arg0 context.Context, arg1 int, arg2 int, arg3 model.Status) (res0 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Update",
			"message", "Update called",
			"request", logUpdateRequest{
				MsgId:  arg2,
				Status: arg3,
				UserId: arg1,
			},
			"err", res0,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Update(arg0, arg1, arg2, arg3)
}

type (
	logAddRequest struct {
		Msg model.MsgInfo
	}
	logAddResponse struct {
		Id int
	}
	logDeleteRequest struct {
		UserId int
		MsgId  int
	}
	logGetAllRequest struct {
		UserId int
		Status model.Status
	}
	logGetAllResponse struct {
		Msgs []model.MsgInfo
	}
	logUpdateRequest struct {
		UserId int
		MsgId  int
		Status model.Status
	}
)
