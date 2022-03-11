package server

import (
	"encoding/json"

	"github.com/jackc/pgx"
)

type Postgres struct {
	conn *pgx.Conn
}

type queryRowInterface interface {
	QueryRow(string, ...interface{}) *pgx.Row
}

func NewPostgres(host string, port uint16, database, user, password string) (*Postgres, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     host,
		Port:     port,
		Database: database,
		User:     user,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &Postgres{
		conn: conn,
	}, nil
}

func (db *Postgres) ListProducts() ([]Product, error) {
	rows, err := db.conn.Query(
		"SELECT id, name, price, amount FROM public.products ORDER BY id ASC",
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []Product{}, nil
		}

		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Amount,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (db *Postgres) GetProduct(id Id) (Product, error) {
	return getProduct(db.conn, id)
}
func (db *Postgres) UpdateProduct(id Id, attrs UpdateAttributes) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	if err = updateProduct(tx, id, attrs); err != nil {
		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (db *Postgres) DeleteProduct(id Id) error {
	_, err := db.conn.Exec("DELETE FROM public.products WHERE id = $1", id)

	return err
}

func (db *Postgres) DeleteAllProducts() error {
	_, err := db.conn.Exec("DELETE FROM public.products")

	return err
}

func (db *Postgres) ReplaceAllWith(products []Product) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec("delete from products"); err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, product := range products {
		if _, err := tx.Exec(
			"INSERT INTO public.products (id, name, price, amount) VALUES ($1, $2, $3, $4)",
			product.Id,
			product.Name,
			product.Price,
			product.Amount,
		); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	defer func(tx *pgx.Tx) {
		_ = tx.Commit()
	}(tx)

	return err
}

func (db *Postgres) Close() error {
	return db.conn.Close()
}

func updateProduct(tx *pgx.Tx, id Id, attrs UpdateAttributes) error {
	product, err := getProduct(tx, id)
	if err != nil {
		return err
	}

	if err = copyToAttrsAndRemoveExtraAttrs(product, attrs); err != nil {
		return err
	}

	if _, err := tx.Exec(
		"UPDATE public.products SET name = $1, price = $2, amount = $3 WHERE id = $4",
		attrs["name"],
		attrs["price"],
		attrs["amount"],
		id,
	); err != nil {
		return err
	}

	return nil
}

func getProduct(executioner queryRowInterface, id Id) (Product, error) {
	product := Product{}
	err := executioner.QueryRow(
		"SELECT id, name, price, amount FROM public.products WHERE id = $1",
		id,
	).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Amount,
	)
	if err == pgx.ErrNoRows {
		return Product{}, ErrProductNotFound
	}

	return product, err
}

func copyToAttrsAndRemoveExtraAttrs(product Product, attrs UpdateAttributes) error {
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	mapping := make(map[string]interface{})
	if err := json.Unmarshal(data, &mapping); err != nil {
		return err
	}

	for key, value := range mapping {
		_, ok := attrs[key]
		if !ok {
			attrs[key] = value
		}
	}

	for key := range attrs {
		_, ok := mapping[key]
		if !ok {
			delete(attrs, key)
		}
	}

	return nil
}
