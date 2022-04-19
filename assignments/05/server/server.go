package server

import (
	"context"
	"domcermak/ctc/assignments/05/common"

	"domcermak/ctc/assignments/05/etcd"
)

type Server struct {
	UnimplementedEtcdServer
	etcdClient *etcd.Client
}

func NewServer(client *etcd.Client) *Server {
	return &Server{
		etcdClient: client,
	}
}

func (s *Server) Get(_ context.Context, request *GetRequest) (*Response, error) {
	response, err := common.LogRequest("get", func() (interface{}, error) {
		return s.etcdClient.Get(etcd.Key(request.Key))
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Body: string(response.(etcd.Response)),
	}, nil
}

func (s *Server) Post(_ context.Context, request *PostRequest) (*Response, error) {
	response, err := common.LogRequest("post", func() (interface{}, error) {
		return s.etcdClient.Post(etcd.Key(request.Key), etcd.Value(request.Value))
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Body: string(response.(etcd.Response)),
	}, nil
}

func (s *Server) Delete(_ context.Context, request *DeleteRequest) (*Response, error) {
	response, err := common.LogRequest("delete", func() (interface{}, error) {
		return s.etcdClient.Delete(etcd.Key(request.Key))
	})
	if err != nil {
		return nil, err
	}

	return &Response{
		Body: string(response.(etcd.Response)),
	}, nil
}
