// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	"context"
	model "github.com/SurajKadam7/msg-info-service/model"
)

func (set EndpointsSet) Add(arg0 context.Context, arg1 model.MsgInfo) (res0 int, res1 error) {
	request := AddRequest{Msg: arg1}
	response, res1 := set.AddEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*AddResponse).Id, res1
}

func (set EndpointsSet) Delete(arg0 context.Context, arg1 int, arg2 int) (res0 error) {
	request := DeleteRequest{
		MsgId:  arg2,
		UserId: arg1,
	}
	_, res0 = set.DeleteEndpoint(arg0, &request)
	if res0 != nil {
		return
	}
	return res0
}

func (set EndpointsSet) GetAll(arg0 context.Context, arg1 int, arg2 model.Status) (res0 []model.MsgInfo, res1 error) {
	request := GetAllRequest{
		Status: arg2,
		UserId: arg1,
	}
	response, res1 := set.GetAllEndpoint(arg0, &request)
	if res1 != nil {
		return
	}
	return response.(*GetAllResponse).Msgs, res1
}

func (set EndpointsSet) Update(arg0 context.Context, arg1 int, arg2 int, arg3 model.Status) (res0 error) {
	request := UpdateRequest{
		MsgId:  arg2,
		Status: arg3,
		UserId: arg1,
	}
	_, res0 = set.UpdateEndpoint(arg0, &request)
	if res0 != nil {
		return
	}
	return res0
}
