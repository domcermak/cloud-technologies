package server

import (
	"domcermak/ctc/assignments/03/tests/helpers"
	"fmt"
	"testing"
	"time"

	"domcermak/ctc/assignments/03/cmd/server"
)

const (
	addr = "localhost:8080"
)

var (
	postgres *server.Postgres = nil
	poolMock *mock            = nil
)

func TestMain(m *testing.M) {
	var err error
	postgres, err = helpers.TestPostgresConnect()
	if err != nil {
		panic(fmt.Sprintf("postgres is offline: %v", err))
	}
	defer postgres.Close()

	poolMock = defaultMock()
	runServerInParallel(poolMock)
	m.Run()
}

func runServerInParallel(pool server.Pool) {
	go func() {
		err := server.NewServer(addr, pool).ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Millisecond * 10)
}

type mock struct {
	list   func() ([]server.Product, error)
	get    func(server.Id) (server.Product, error)
	update func(server.Id, server.UpdateAttributes) error
	delete func(server.Id) error
	close  func() error
}

func defaultMock() *mock {
	return &mock{
		list: func() ([]server.Product, error) {
			panic("not implemented")
		},
		get: func(id server.Id) (server.Product, error) {
			panic("not implemented")
		},
		update: func(id server.Id, attributes server.UpdateAttributes) error {
			panic("not implemented")
		},
		delete: func(id server.Id) error {
			panic("not implemented")
		},
		close: func() error {
			return nil
		},
	}
}

func (m mock) ListProducts() ([]server.Product, error) {
	return m.list()
}

func (m mock) GetProduct(id server.Id) (server.Product, error) {
	return m.get(id)
}

func (m mock) UpdateProduct(id server.Id, attributes server.UpdateAttributes) error {
	return m.update(id, attributes)
}

func (m mock) DeleteProduct(id server.Id) error {
	return m.delete(id)
}

func (m mock) Close() error {
	return m.close()
}
