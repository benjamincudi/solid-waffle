package waffledb

import (
	"context"
	"fmt"
	"solid-waffle/models"

	"github.com/georgysavva/scany/pgxscan"
)

func (ds WaffleDB) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := pgxscan.Select(ctx, ds.conn, &products, "SELECT * FROM products")
	return products, err
}

func (ds WaffleDB) InsertProduct(ctx context.Context, sellerId int, name string) (models.Product, error) {
	var product []models.Product
	err := pgxscan.Select(ctx, ds.conn, &product, fmt.Sprintf("INSERT INTO products (created_at, updated_at, seller_id, name) VALUES (now(), now(), %d, '%s') RETURNING id, created_at, updated_at, seller_id, name", sellerId, name))
	return product[0], err
}
