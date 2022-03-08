package server

type Pool interface {
	ListProducts() ([]Product, error)
	GetProduct(Id) (Product, error)
	UpdateProduct(Id, UpdateAttributes) error
	DeleteProduct(Id) error
	Close() error
}

type UpdateAttributes map[string]interface{}
