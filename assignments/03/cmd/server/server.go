package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewServer(addr string, pool Pool) *http.Server {
	router := mux.NewRouter()
	handler := NewHandler(pool)

	router.
		HandleFunc("/products", handler.ListProductsHandler).
		Methods(http.MethodGet)
	router.
		HandleFunc("/products/{id:[0-9]+}", handler.GetProductHandler).
		Methods(http.MethodGet)
	router.
		HandleFunc("/products/{id:[0-9]+}", handler.UpdateProductHandler).
		Methods(http.MethodPatch)
	router.
		HandleFunc("/products/{id:[0-9]+}", handler.DeleteProductHandler).
		Methods(http.MethodDelete)

	return &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}
}
