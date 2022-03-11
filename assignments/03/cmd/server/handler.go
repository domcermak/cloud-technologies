package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	pool Pool
}

func NewHandler(pool Pool) *Handler {
	return &Handler{
		pool: pool,
	}
}

func (h *Handler) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	data, statusCode := marshal(func() (interface{}, error) {
		return h.pool.ListProducts()
	}, http.StatusOK)

	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	data, statusCode := marshal(func() (interface{}, error) {
		id, err := parseId(r)
		if err != nil {
			return nil, err
		}

		return h.pool.GetProduct(id)
	}, http.StatusOK)

	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}
func (h *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	data, statusCode := marshal(func() (interface{}, error) {
		id, err := parseId(r)
		if err != nil {
			return nil, err
		}

		attrs := UpdateAttributes{}
		if err := json.NewDecoder(r.Body).Decode(&attrs); err != nil {
			return nil, err
		}

		return "", h.pool.UpdateProduct(id, attrs)
	}, http.StatusNoContent)

	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}
func (h *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	r.Close = true
	data, statusCode := marshal(func() (interface{}, error) {
		id, err := parseId(r)
		if err != nil {
			return nil, err
		}

		return "", h.pool.DeleteProduct(id)
	}, http.StatusNoContent)

	w.WriteHeader(statusCode)
	_, _ = w.Write(data)
}

func marshal(fn func() (interface{}, error), successStatusCode int) ([]byte, int) {
	data, err := fn()
	if err != nil {
		status, err := ErrorMapper(err)
		data, err := json.Marshal(map[string]interface{}{
			"error": err.Error(),
		})
		if err != nil {
			panic(err)
		}

		return data, status
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return jsonData, successStatusCode
}

func parseId(r *http.Request) (Id, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, ErrInvalidIdValue
	}

	return Id(id), nil
}
