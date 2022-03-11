package server

import (
	"testing"

	"domcermak/ctc/assignments/03/cmd/server"
	"domcermak/ctc/assignments/03/tests/helpers"
)

func TestPostgres_ListProducts_WithData(t *testing.T) {
	helpers.CleanupTestPostgres(postgres)()
	expected := helpers.AllTestProducts()

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	helpers.ExpectProducts(expected, products, t)
}

func TestPostgres_ListProducts_WithoutData(t *testing.T) {
	t.Cleanup(helpers.CleanupTestPostgres(postgres))

	if err := postgres.DeleteAllProducts(); err != nil {
		t.Fatal(err)
	}

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}

	helpers.Expect(0, len(products), t)
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
					helpers.Expect(server.ErrProductNotFound, err, t)
					return
				}
				t.Fatal(err)
			}

			helpers.Expect(tc.expected.Name, product.Name, t)
			helpers.Expect(tc.expected.Price, product.Price, t)
			helpers.Expect(tc.expected.Amount, product.Amount, t)
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
			name: "Changes product with string attributes",
			id:   products[0].Id,
			updateParams: map[string]interface{}{
				"name":   "something new",
				"price":  "11111.11111",
				"amount": "11111",
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
				helpers.Expect(tc.err, err, t)
				return
			}

			updatedProduct, err := postgres.GetProduct(tc.id)
			if err != nil {
				t.Fatal(err)
			}

			helpers.Expect(tc.expectedProduct.Id, updatedProduct.Id, t)
			helpers.Expect(tc.expectedProduct.Name, updatedProduct.Name, t)
			helpers.Expect(tc.expectedProduct.Price, updatedProduct.Price, t)
			helpers.Expect(tc.expectedProduct.Amount, updatedProduct.Amount, t)
		})
	}
}

func TestPostgres_DeleteProduct(t *testing.T) {
	t.Cleanup(helpers.CleanupTestPostgres(postgres))

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
				helpers.Expect(tc.err, err, t)
				return
			}

			_, err := postgres.GetProduct(tc.id)
			helpers.Expect(server.ErrProductNotFound, err, t)
			productsAfterOnDeleted, err := postgres.ListProducts()
			if err != nil {
				t.Fatal(err)
			}
			helpers.Expect(len(products)-1, len(productsAfterOnDeleted), t)
		})
	}
}

func TestPostgres_DeleteAllProducts(t *testing.T) {
	t.Cleanup(helpers.CleanupTestPostgres(postgres))

	if err := postgres.DeleteAllProducts(); err != nil {
		t.Fatal(err)
	}

	products, err := postgres.ListProducts()
	if err != nil {
		t.Fatal(err)
	}
	helpers.Expect(0, len(products), t)
}
