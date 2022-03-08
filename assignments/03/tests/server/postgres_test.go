package server

import (
	"domcermak/ctc/assignments/03/cmd/server"
	"fmt"
	"testing"
)

func TestPostgres_ListProducts_WithData(t *testing.T) {
	expected := allProducts()

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	expect(len(expected), len(products), t)
	for i, product := range products {
		expect(expected[i].Name, product.Name, t)
		expect(expected[i].Price, product.Price, t)
		expect(expected[i].Amount, product.Amount, t)
	}
}

func TestPostgres_ListProducts_WithoutData(t *testing.T) {
	t.Cleanup(cleanup)

	if err := postgres.DeleteAllProducts(); err != nil {
		t.Fatal(err)
	}

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	expect(0, len(products), t)
}

func TestPostgres_GetProduct(t *testing.T) {
	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name     string
		id       server.Id
		expected *server.Product
	}{
		{
			name: "Returns the product",
			id:   products[0].Id,
			expected: &server.Product{
				Id:     1,
				Name:   "clementine",
				Price:  1.38,
				Amount: 8,
			},
		},
		{
			name:     "Does not find the product",
			id:       products[len(products)-1].Id + 1,
			expected: nil,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			product, err := postgres.GetProduct(tc.id)
			if err != nil {
				if tc.expected == nil {
					expect(server.ErrProductNotFound, err, t)
					return
				}
				t.Fatal(err)
			}

			expect(tc.expected.Name, product.Name, t)
			expect(tc.expected.Price, product.Price, t)
			expect(tc.expected.Amount, product.Amount, t)
		})
	}
}

func TestPostgres_UpdateProduct(t *testing.T) {
	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name            string
		id              server.Id
		updateParams    server.UpdateAttributes
		err             error
		expectedProduct server.Product
	}{
		{
			name:            "No update attributes passed - product stays unchanged",
			id:              products[0].Id,
			updateParams:    make(server.UpdateAttributes),
			err:             nil,
			expectedProduct: products[0],
		},
		{
			name: "Changes product attributes",
			id:   products[0].Id,
			updateParams: map[string]interface{}{
				"name":   "something new",
				"price":  11111.11111,
				"amount": 11111,
			},
			err: nil,
			expectedProduct: server.Product{
				Id:     products[0].Id,
				Name:   "something new",
				Price:  11111.11111,
				Amount: 11111,
			},
		},
		{
			name:         "Does not find the product",
			id:           products[len(products)-1].Id + 1,
			updateParams: make(server.UpdateAttributes),
			err:          server.ErrProductNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := postgres.UpdateProduct(tc.id, tc.updateParams)
			if err != nil {
				expect(tc.err, err, t)
				return
			}

			updatedProduct, err := postgres.GetProduct(tc.id)
			if err != nil {
				t.Fatal(err)
			}

			expect(tc.expectedProduct.Id, updatedProduct.Id, t)
			expect(tc.expectedProduct.Name, updatedProduct.Name, t)
			expect(tc.expectedProduct.Price, updatedProduct.Price, t)
			expect(tc.expectedProduct.Amount, updatedProduct.Amount, t)
		})
	}
}

func TestPostgres_DeleteProduct(t *testing.T) {
	t.Cleanup(cleanup)

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name string
		id   server.Id
		err  error
	}{
		{
			name: "Deletes the product",
			id:   products[0].Id,
			err:  nil,
		},
		{
			name: "Does not find the product",
			id:   products[len(products)-1].Id + 1,
			err:  server.ErrProductNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := postgres.DeleteProduct(tc.id); err != nil {
				expect(tc.err, err, t)
				return
			}

			_, err := postgres.GetProduct(tc.id)
			expect(server.ErrProductNotFound, err, t)
			productsAfterOnDeleted, err := postgres.ListProducts()
			if err != nil {
				t.Fatal(err)
			}
			expect(len(products)-1, len(productsAfterOnDeleted), t)
		})
	}
}

func TestPostgres_DeleteAllProducts(t *testing.T) {
	t.Cleanup(cleanup)

	if err := postgres.DeleteAllProducts(); err != nil {
		t.Fatal(err)
	}

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}
	expect(0, len(products), t)
}

func expect(expected, actual interface{}, t *testing.T) {
	if fmt.Sprintf("%v", expected) != fmt.Sprintf("%v", actual) {
		t.Fatalf("Expected %v, but got %v", expected, actual)
	}
}

func allProducts() []server.Product {
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

func cleanup() {
	err := postgres.DeleteAllProducts()
	if err != nil {
		panic(err)
	}

	for _, product := range allProducts() {
		if err := postgres.InsertProduct(product); err != nil {
			panic(err)
		}
	}
}
