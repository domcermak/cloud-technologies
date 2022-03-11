package helpers

import "domcermak/ctc/assignments/03/cmd/server"

func TestPostgresConnect() (*server.Postgres, error) {
	return server.NewPostgres(
		"localhost",
		1111,
		"postgres",
		"postgres",
		"postgres",
	)
}

func AllTestProducts() []server.Product {
	return []server.Product{
		{
			Id:     1,
			Name:   "clementine",
			Price:  1.38,
			Amount: 8,
		},
		{
			Id:     2,
			Name:   "apricot",
			Price:  12.3,
			Amount: 12,
		},
		{
			Id:     3,
			Name:   "peach",
			Price:  1.1,
			Amount: 1,
		},
		{
			Id:     4,
			Name:   "star fruit",
			Price:  1,
			Amount: 24,
		},
		{
			Id:     5,
			Name:   "huckleberry",
			Price:  33.9,
			Amount: 11,
		},
		{
			Id:     6,
			Name:   "jujube",
			Price:  12.89,
			Amount: 7,
		},
	}
}

func CleanupTestPostgres(pool *server.Postgres) func() {
	return func() {
		if err := pool.ReplaceAllWith(AllTestProducts()); err != nil {
			panic(err)
		}
	}
}
