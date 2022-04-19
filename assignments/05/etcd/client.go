package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
)

type Response string

type Key string

type Value string

type pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Client struct {
	client *clientv3.Client
}

func NewClient(addr string) (*Client, error) {
	client, err := clientv3.New(
		clientv3.Config{
			Endpoints:   []string{addr},
			DialTimeout: 2 * time.Second,
		},
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Get(key Key) (Response, error) {
	response, err := c.client.Get(context.Background(), string(key))
	if err != nil {
		return "", err
	}

	pairs := make([]pair, len(response.Kvs))
	for i, kv := range response.Kvs {
		pairs[i] = pair{
			Key:   string(kv.Key),
			Value: string(kv.Value),
		}
	}

	data, err := json.Marshal(pairs)
	if err != nil {
		return "", err
	}

	return Response(data), nil
}

func (c *Client) Post(key Key, value Value) (Response, error) {
	_, err := c.client.Put(context.Background(), string(key), string(value))

	return "ok", err
}

func (c *Client) Delete(key Key) (Response, error) {
	response, err := c.client.Delete(
		context.Background(),
		string(key),
	)
	if err != nil {
		return "", err
	}

	return Response(fmt.Sprintf("deleted: %d", response.Deleted)), nil
}
