package client

import (
	"encoding/json"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"

	cl "domcermak/ctc/assignments/03/cmd/client"
	"domcermak/ctc/assignments/03/cmd/server"
	"domcermak/ctc/assignments/03/tests/helpers"
)

func TestClient_ListProducts(t *testing.T) {
	for _, tc := range []struct {
		name             string
		statusCode       int
		expectedProducts []server.Product
	}{
		{
			name: "Returns a non-empty list of products",
			expectedProducts: []server.Product{
				{
					Id:     1,
					Name:   "Cabbage",
					Price:  1.00,
					Amount: 20,
				},
			},
			statusCode: http.StatusOK,
		},
		{
			name:             "Returns an empty list of products",
			expectedProducts: []server.Product{},
			statusCode:       http.StatusOK,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testServer := newTestServer(map[string]func(http.ResponseWriter, *http.Request){
				"list": func(writer http.ResponseWriter, _ *http.Request) {
					data, err := json.Marshal(tc.expectedProducts)
					if err != nil {
						t.Fatal(err)
					}

					writer.WriteHeader(tc.statusCode)
					_, _ = writer.Write(data)
				},
			})
			addr, quitChan := testServer.Run()
			defer func() {
				quitChan <- 0
			}()

			client := cl.NewClient(time.Second, addr)
			products, err := client.ListProducts()
			if err != nil {
				t.Fatal(err)
			}

			helpers.ExpectProducts(tc.expectedProducts, products, t)
		})
	}
}

func TestClient_GetProduct(t *testing.T) {
	for _, tc := range []struct {
		name            string
		statusCode      int
		expectedProduct server.Product
		expectedErr     error
	}{
		{
			name: "Returns the requested product",
			expectedProduct: server.Product{
				Id:     1,
				Name:   "Cabbage",
				Price:  1.00,
				Amount: 20,
			},
			expectedErr: nil,
			statusCode:  http.StatusOK,
		},
		{
			name:            "Returns not found error",
			expectedProduct: server.Product{},
			expectedErr:     server.ErrProductNotFound,
			statusCode:      http.StatusNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testServer := newTestServer(map[string]func(http.ResponseWriter, *http.Request){
				"get": func(writer http.ResponseWriter, _ *http.Request) {
					data := func() []byte {
						var toBeMarshaled interface{} = tc.expectedProduct
						if tc.expectedErr != nil {
							toBeMarshaled = map[string]string{
								"error": tc.expectedErr.Error(),
							}
						}

						data, err := json.Marshal(toBeMarshaled)
						if err != nil {
							t.Fatal(err)
						}

						return data
					}()

					writer.WriteHeader(tc.statusCode)
					_, _ = writer.Write(data)
				},
			})
			addr, quitChan := testServer.Run()
			defer func() {
				quitChan <- 0
			}()
			client := cl.NewClient(time.Second, addr)
			product, err := client.GetProduct(tc.expectedProduct.Id)
			helpers.Expect(tc.expectedErr, err, t)
			helpers.ExpectProduct(tc.expectedProduct, product, t)
		})
	}
}

func TestClient_UpdateProduct(t *testing.T) {
	for _, tc := range []struct {
		name             string
		id               server.Id
		statusCode       int
		updateAttributes server.UpdateAttributes
		expectedErr      error
	}{
		{
			name: "Updates the product",
			id:   1,
			updateAttributes: map[string]interface{}{
				"name":   "new name",
				"amount": 22,
				"price":  1.00,
			},
			statusCode:  http.StatusNoContent,
			expectedErr: nil,
		},
		{
			name:             "Returns not found error",
			id:               1,
			updateAttributes: map[string]interface{}{},
			statusCode:       http.StatusNotFound,
			expectedErr:      server.ErrProductNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testServer := newTestServer(map[string]func(http.ResponseWriter, *http.Request){
				"update": func(writer http.ResponseWriter, r *http.Request) {
					helpers.ExpectIdInRequestVars(tc.id, r, t)
					helpers.ExpectJsonRequestBody(tc.updateAttributes, r, t)

					data := func() []byte {
						var toBeMarshaled interface{} = ""
						if tc.expectedErr != nil {
							toBeMarshaled = map[string]string{
								"error": tc.expectedErr.Error(),
							}
						}

						data, err := json.Marshal(toBeMarshaled)
						if err != nil {
							t.Fatal(err)
						}

						return data
					}()

					writer.WriteHeader(tc.statusCode)
					_, _ = writer.Write(data)
				},
			})

			addr, quitChan := testServer.Run()
			defer func() {
				quitChan <- 0
			}()
			client := cl.NewClient(time.Second, addr)
			err := client.UpdateProduct(tc.id, tc.updateAttributes)
			helpers.Expect(tc.expectedErr, err, t)
		})
	}
}

func TestClient_DeleteProduct(t *testing.T) {
	for _, tc := range []struct {
		name        string
		id          server.Id
		statusCode  int
		expectedErr error
	}{
		{
			name:        "Deletes the product",
			id:          1,
			statusCode:  http.StatusNoContent,
			expectedErr: nil,
		},
		{
			name:        "Returns not found error",
			id:          1,
			statusCode:  http.StatusNotFound,
			expectedErr: server.ErrProductNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			testServer := newTestServer(map[string]func(http.ResponseWriter, *http.Request){
				"delete": func(writer http.ResponseWriter, r *http.Request) {
					helpers.ExpectIdInRequestVars(tc.id, r, t)

					data := func() []byte {
						var toBeMarshaled interface{} = ""
						if tc.expectedErr != nil {
							toBeMarshaled = map[string]string{
								"error": tc.expectedErr.Error(),
							}
						}

						data, err := json.Marshal(toBeMarshaled)
						if err != nil {
							t.Fatal(err)
						}

						return data
					}()

					writer.WriteHeader(tc.statusCode)
					_, _ = writer.Write(data)
				},
			})

			addr, quitChan := testServer.Run()
			defer func() {
				quitChan <- 0
			}()
			client := cl.NewClient(time.Second, addr)
			err := client.DeleteProduct(tc.id)
			helpers.Expect(tc.expectedErr, err, t)
		})
	}
}

type testServer struct {
	*http.Server
}

func newTestServer(handlers map[string]func(http.ResponseWriter, *http.Request)) *testServer {
	addDefaultTestHandlers(handlers)

	router := &mux.Router{}
	router.
		HandleFunc("/products", handlers["list"]).
		Methods(http.MethodGet)
	router.
		HandleFunc("/products/{id:[0-9]+}", handlers["get"]).
		Methods(http.MethodGet)
	router.
		HandleFunc("/products/{id:[0-9]+}", handlers["update"]).
		Methods(http.MethodPatch)
	router.
		HandleFunc("/products/{id:[0-9]+}", handlers["delete"]).
		Methods(http.MethodDelete)

	return &testServer{
		Server: &http.Server{
			Handler:      router,
			WriteTimeout: time.Second,
			ReadTimeout:  time.Second,
		},
	}
}

func (s *testServer) Run() (string, chan<- interface{}) {
	quitChan := make(chan interface{})
	addr := make(chan string)

	go func() {
		go func() {
			listener, err := net.Listen("tcp", ":0") // find a free port
			if err != nil {
				panic(err)
			}
			addr <- listener.Addr().String()

			if err := s.Serve(listener); err != nil {
				panic(err)
			}
		}()

		// quitting this goroutine quits also the child
		// goroutine with the test server
		<-quitChan
	}()

	return <-addr, quitChan
}

func addDefaultTestHandlers(handlers map[string]func(http.ResponseWriter, *http.Request)) {
	for _, label := range []string{"list", "get", "update", "delete"} {
		_, ok := handlers[label]
		if !ok {
			handlers[label] = func(http.ResponseWriter, *http.Request) {}
		}
	}
}
