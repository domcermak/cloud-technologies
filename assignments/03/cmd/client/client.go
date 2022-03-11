package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"domcermak/ctc/assignments/03/cmd/server"
)

type ProductClientInterface interface {
	ListProducts() ([]server.Product, error)
	GetProduct(id server.Id) (server.Product, error)
	UpdateProduct(id server.Id) error
	DeleteProduct(id server.Id) error
}

type Client struct {
	client     *http.Client
	serverAddr string
}

func NewClient(timeout time.Duration, address string) *Client {
	// Transport necessary, otherwise causes EOF error
	// while sending request do a docker container
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Client{
		client: &http.Client{
			Timeout:   timeout,
			Transport: transport,
		},
		serverAddr: address,
	}
}

func (c *Client) ListProducts() ([]server.Product, error) {
	res, err := c.client.Get(c.listProductsUrl())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, parseErrorMessage(res.Body)
	}

	var products []server.Product
	err = json.NewDecoder(res.Body).Decode(&products)

	return products, err
}

func (c *Client) GetProduct(id server.Id) (server.Product, error) {
	res, err := c.client.Get(c.getProductUrl(id))
	if err != nil {
		return server.Product{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return server.Product{}, parseErrorMessage(res.Body)
	}

	var product server.Product
	err = json.NewDecoder(res.Body).Decode(&product)

	return product, err
}

func (c *Client) UpdateProduct(id server.Id, attributes server.UpdateAttributes) error {
	data, err := json.Marshal(attributes)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, c.updateProductUrl(id), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return parseErrorMessage(res.Body)
	}

	return nil
}

func (c *Client) DeleteProduct(id server.Id) error {
	req, err := http.NewRequest(http.MethodDelete, c.deleteProductUrl(id), bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return parseErrorMessage(res.Body)
	}

	return nil
}

func (c *Client) CommandExecutioners() []CommandExecutioner {
	return []CommandExecutioner{
		{
			CommandName: "list",
			Description: "Returns all existing products",
			Execute: func(_ map[string]interface{}) (string, error) {
				products, err := c.ListProducts()
				if err != nil {
					return "", err
				}

				return prettyJson(products)
			},
		},
		{
			CommandName: "get",
			Description: "Returns a product by given id",
			RequiredArgs: map[string]interface{}{
				"id": "<number>",
			},
			Execute: func(options map[string]interface{}) (string, error) {
				id, err := getInt(options, "id")
				if err != nil {
					return "", err
				}

				product, err := c.GetProduct(server.Id(id))
				if err != nil {
					return "", err
				}

				return prettyJson(product)
			},
		},
		{
			CommandName: "delete",
			Description: "Deletes a product by given id",
			RequiredArgs: map[string]interface{}{
				"id": "<number>",
			},
			Execute: func(options map[string]interface{}) (string, error) {
				id, err := getInt(options, "id")
				if err != nil {
					return "", err
				}

				return fmt.Sprintf(
					"Product with id %d deleted",
					id,
				), c.DeleteProduct(server.Id(id))
			},
		},
		{
			CommandName: "update",
			Description: "Updates a product by given id",
			RequiredArgs: map[string]interface{}{
				"id": "<number>",
			},
			OptionalArgs: map[string]interface{}{
				"name":   "<text>",
				"amount": "<number>",
				"price":  "<number>",
			},
			Execute: func(options map[string]interface{}) (string, error) {
				id, err := getInt(options, "id")
				if err != nil {
					return "", err
				}

				return fmt.Sprintf(
					"Product with id %d updated",
					id,
				), c.UpdateProduct(server.Id(id), options)
			},
		},
	}
}

func (c *Client) getProductUrl(id server.Id) string {
	return fmt.Sprintf("%s/%d", c.productsEndpoint(), id)
}

func (c *Client) deleteProductUrl(id server.Id) string {
	return c.getProductUrl(id)
}

func (c *Client) updateProductUrl(id server.Id) string {
	return c.getProductUrl(id)
}

func (c *Client) listProductsUrl() string {
	return c.productsEndpoint()
}

func (c *Client) productsEndpoint() string {
	return fmt.Sprintf("http://%s/products", c.serverAddr)
}

func parseErrorMessage(reader io.Reader) error {
	mapping := make(map[string]string)
	err := json.NewDecoder(reader).Decode(&mapping)
	if err != nil {
		return err
	}

	return errors.New(mapping["error"])
}
