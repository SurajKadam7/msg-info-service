// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	msginfosrv "github.com/suraj.kadam7/msg-info-srv/msginfo_srv"
)

func Endpoints(svc msginfosrv.Service) EndpointsSet {
	return EndpointsSet{
		AddEndpoint:    AddEndpoint(svc),
		DeleteEndpoint: DeleteEndpoint(svc),
		GetEndpoint:    GetEndpoint(svc),
		UpdateEndpoint: UpdateEndpoint(svc),
	}
}

func AddEndpoint(svc msginfosrv.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*AddRequest)
		res0, res1 := svc.Add(arg0, req.Msg)
		return &AddResponse{Id: res0}, res1
	}
}

func DeleteEndpoint(svc msginfosrv.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*DeleteRequest)
		res0 := svc.Delete(arg0, req.UserId, req.MsgId)
		return &DeleteResponse{}, res0
	}
}

func GetEndpoint(svc msginfosrv.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetRequest)
		res0, res1 := svc.Get(arg0, req.UserId, req.Status)
		return &GetResponse{Msgs: res0}, res1
	}
}

func UpdateEndpoint(svc msginfosrv.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*UpdateRequest)
		res0 := svc.Update(arg0, req.UserId, req.MsgId, req.Status)
		return &UpdateResponse{}, res0
	}
}