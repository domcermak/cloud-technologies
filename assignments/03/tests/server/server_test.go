package server

import (
	"bytes"
	"domcermak/ctc/assignments/03/cmd/server"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestServer_ListProducts(t *testing.T) {
	client := client()

	for _, tc := range []struct {
		name             string
		expectedProducts []server.Product
	}{
		{
			name:             "Returns a non-empty list",
			expectedProducts: allProducts(),
		},
		{
			name:             "Returns an empty list of products",
			expectedProducts: []server.Product{},
		},
	} {
		poolMock.list = func() ([]server.Product, error) {
			return tc.expectedProducts, nil
		}

		res, err := client.Get(fmt.Sprintf("http://%s/products", addr))
		expect(nil, err, t)
		func() {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					t.Fatal(err)
				}
			}(res.Body)

			expect(http.StatusOK, res.StatusCode, t)

			var receivedProducts []server.Product
			err := json.NewDecoder(res.Body).Decode(&receivedProducts)
			expect(nil, err, t)
			expect(len(tc.expectedProducts), len(receivedProducts), t)

			for i, receivedProduct := range receivedProducts {
				expect(tc.expectedProducts[i].Id, receivedProduct.Id, t)
				expect(tc.expectedProducts[i].Name, receivedProduct.Name, t)
				expect(tc.expectedProducts[i].Price, receivedProduct.Price, t)
				expect(tc.expectedProducts[i].Amount, receivedProduct.Amount, t)
			}
		}()
	}
}

func TestServer_GetProduct(t *testing.T) {
	client := client()

	for _, tc := range []struct {
		name            string
		expectedProduct server.Product
		statusCode      int
		err             error
	}{
		{
			name:            "Successfully returns the product",
			expectedProduct: allProducts()[0],
			statusCode:      http.StatusOK,
			err:             nil,
		},
		{
			name:            "Returns not found error",
			expectedProduct: server.Product{Id: 20},
			statusCode:      http.StatusNotFound,
			err:             server.ErrProductNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			poolMock.get = func(id server.Id) (server.Product, error) {
				expect(tc.expectedProduct.Id, id, t)
				return tc.expectedProduct, tc.err
			}
			res, err := client.Get(fmt.Sprintf("http://%s/products/%d", addr, tc.expectedProduct.Id))
			expect(nil, err, t)
			func() {
				defer func(Body io.ReadCloser) {
					if err := Body.Close(); err != nil {
						t.Fatal(err)
					}
				}(res.Body)

				expect(tc.statusCode, res.StatusCode, t)
				decoder := json.NewDecoder(res.Body)

				if tc.err != nil {
					mapping := make(map[string]string)
					err := decoder.Decode(&mapping)
					expect(nil, err, t)
					expect(tc.err.Error(), mapping["error"], t)
					return
				}

				var receivedProduct server.Product
				err := decoder.Decode(&receivedProduct)
				expect(nil, err, t)

				expect(tc.expectedProduct.Id, receivedProduct.Id, t)
				expect(tc.expectedProduct.Name, receivedProduct.Name, t)
				expect(tc.expectedProduct.Price, receivedProduct.Price, t)
				expect(tc.expectedProduct.Amount, receivedProduct.Amount, t)
			}()
		})
	}
}

func TestServer_UpdateProduct(t *testing.T) {
	client := client()
	sampleProduct := server.Product{
		Id:   1,
		Name: "sample",
	}

	for _, tc := range []struct {
		name               string
		params             server.UpdateAttributes
		expectedStatusCode int
		expectedErr        error
	}{
		{
			name: "Updates the product",
			params: map[string]interface{}{
				"name": "new name",
				// price is missing - will not be updated
				"amount":        22,
				"unknown_param": 11,
			},
			expectedStatusCode: http.StatusNoContent,
			expectedErr:        nil,
		},
		{
			name:               "Returns not found error",
			expectedStatusCode: http.StatusNotFound,
			expectedErr:        server.ErrProductNotFound,
		},
	} {
		poolMock.get = func(id server.Id) (server.Product, error) {
			expect(sampleProduct.Id, id, t)
			return sampleProduct, tc.expectedErr
		}
		t.Run(tc.name, func(t *testing.T) {
			poolMock.update = func(id server.Id, params server.UpdateAttributes) error {
				expect(sampleProduct.Id, id, t)
				expect(tc.params, params, t)

				return tc.expectedErr
			}

			jsonData, err := json.Marshal(tc.params)
			expect(nil, err, t)

			req, err := http.NewRequest(
				http.MethodPatch,
				fmt.Sprintf("http://%s/products/%d", addr, sampleProduct.Id),
				bytes.NewBuffer(jsonData),
			)
			expect(nil, err, t)

			res, err := client.Do(req)
			expect(nil, err, t)

			func() {
				defer func(Body io.ReadCloser) {
					if err := Body.Close(); err != nil {
						t.Fatal(err)
					}
				}(res.Body)

				expect(tc.expectedStatusCode, res.StatusCode, t)
				if tc.expectedStatusCode == http.StatusNoContent {
					return // nothing to parse
				}

				decodedError := make(map[string]string)
				if err := json.NewDecoder(res.Body).Decode(&decodedError); err != nil {
					t.Fatal(err)
				}

				expect(tc.expectedErr, decodedError["error"], t)
			}()
		})
	}
}

func TestServer_DeleteProduct(t *testing.T) {
	client := client()

	for _, tc := range []struct {
		name               string
		id                 server.Id
		expectedErr        error
		expectedStatusCode int
	}{
		{
			name:               "Deletes the product",
			id:                 1,
			expectedErr:        nil,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "Returns not found error",
			id:                 1,
			expectedErr:        server.ErrProductNotFound,
			expectedStatusCode: http.StatusNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			poolMock.delete = func(id server.Id) error {
				expect(tc.id, id, t)
				return tc.expectedErr
			}

			req, err := http.NewRequest(
				http.MethodDelete,
				fmt.Sprintf("http://%s/products/%d", addr, tc.id),
				&bytes.Buffer{},
			)
			expect(nil, err, t)

			res, err := client.Do(req)
			expect(nil, err, t)
			func() {
				defer func(Body io.ReadCloser) {
					if err := Body.Close(); err != nil {
						t.Fatal(err)
					}
				}(res.Body)

				expect(tc.expectedStatusCode, res.StatusCode, t)
				if tc.expectedStatusCode == http.StatusNoContent {
					return // nothing to parse
				}

				decodedError := make(map[string]string)
				if err := json.NewDecoder(res.Body).Decode(&decodedError); err != nil {
					t.Fatal(err)
				}

				expect(tc.expectedErr, decodedError["error"], t)
			}()
		})
	}
}

func client() *http.Client {
	return &http.Client{Timeout: time.Second}
}
