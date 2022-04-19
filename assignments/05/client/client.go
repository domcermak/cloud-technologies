package client

import (
	"context"
	"time"

	"domcermak/ctc/assignments/05/common"
	"domcermak/ctc/assignments/05/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client server.EtcdClient
}

func NewClient(addr string) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(
		ctx, addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: server.NewEtcdClient(conn),
	}, nil
}

func (c *Client) Get(key string) (string, error) {
	response, err := common.LogRequest("get", func() (interface{}, error) {
		return c.client.Get(
			context.Background(),
			&server.GetRequest{
				Key: key,
			},
		)
	})
	if err != nil {
		return "", err
	}
	return response.(*server.Response).Body, nil
}

func (c *Client) Post(key string, value string) (string, error) {
	response, err := common.LogRequest("post", func() (interface{}, error) {
		return c.client.Post(
			context.Background(),
			&server.PostRequest{
				Key:   key,
				Value: value,
			},
		)
	})
	if err != nil {
		return "", err
	}
	return response.(*server.Response).Body, nil
}

func (c *Client) Delete(key string) (string, error) {
	response, err := common.LogRequest("delete", func() (interface{}, error) {
		return c.client.Delete(
			context.Background(),
			&server.DeleteRequest{
				Key: key,
			},
		)
	})
	if err != nil {
		return "", err
	}
	return response.(*server.Response).Body, nil
}
