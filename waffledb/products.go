package waffledb

import (
	"context"
	"fmt"
	"solid-waffle/models"
	"strings"

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

func (ds WaffleDB) GetVariations(ctx context.Context) ([]models.Variation, error) {
	var variations []models.Variation
	err := pgxscan.Select(ctx, ds.conn, &variations, "SELECT * FROM variations")
	return variations, err
}

func (ds WaffleDB) InsertVariation(ctx context.Context, name string) (models.Variation, error) {
	var variation []models.Variation
	err := pgxscan.Select(ctx, ds.conn, &variation, fmt.Sprintf("INSERT INTO variations (created_at, updated_at, name) VALUES (now(), now(), '%s') RETURNING id, created_at, updated_at, name", name))
	return variation[0], err
}

func (ds WaffleDB) GetProductVariations(ctx context.Context) ([]models.ProductVariation, error) {
	var products []models.ProductVariation
	err := pgxscan.Select(ctx, ds.conn, &products, "SELECT * FROM product_variations")
	return products, err
}

func (ds WaffleDB) InsertProductVariation(ctx context.Context, variation models.ProductVariation) (models.ProductVariation, error) {
	stringifiedArray := strings.Trim(strings.Replace(fmt.Sprint(variation.VariationIDs), " ", ",", -1), "[]")
	var product []models.ProductVariation
	err := pgxscan.Select(ctx, ds.conn, &product, fmt.Sprintf("INSERT INTO product_variations (created_at, updated_at, product_id, variation_ids, price, is_available) VALUES (now(), now(), %d, '{%s}', %d, %t) RETURNING id, created_at, updated_at, product_id, variation_ids, price, is_available", variation.ProductID, stringifiedArray, variation.Price, variation.IsAvailable))
	return product[0], err
}
