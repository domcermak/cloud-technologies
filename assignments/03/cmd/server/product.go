package server

type Id uint64

type Product struct {
	Id     Id      `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount uint64  `json:"amount"`
}
