// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import model "github.com/suraj.kadam7/msg-info-srv/model"

type (
	AddRequest struct {
		Msg model.MsgInfo `json:"msg"`
	}
	AddResponse struct {
		Id int `json:"id"`
	}

	DeleteRequest struct {
		UserId int `json:"user_id"`
		MsgId  int `json:"msg_id"`
	}
	// Formal exchange type, please do not delete.
	DeleteResponse struct{}

	GetRequest struct {
		UserId int          `json:"user_id"`
		Status model.Status `json:"status"`
	}
	GetResponse struct {
		Msgs []model.MsgInfo `json:"msgs"`
	}

	UpdateRequest struct {
		UserId int          `json:"user_id"`
		MsgId  int          `json:"msg_id"`
		Status model.Status `json:"status"`
	}
	// Formal exchange type, please do not delete.
	UpdateResponse struct{}
)