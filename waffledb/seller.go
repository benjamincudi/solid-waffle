package waffledb

import (
	"context"
	"fmt"
	"solid-waffle/models"

	"github.com/georgysavva/scany/pgxscan"
)

func (ds WaffleDB) GetSellers(ctx context.Context) ([]models.Seller, error) {
	var sellers []models.Seller
	err := pgxscan.Select(ctx, ds.conn, &sellers, "SELECT * FROM sellers")
	return sellers, err
}

func (ds WaffleDB) InsertSeller(ctx context.Context, name string) (models.Seller, error) {
	var seller []models.Seller
	err := pgxscan.Select(ctx, ds.conn, &seller, fmt.Sprintf("INSERT INTO sellers (created_at, updated_at, name) VALUES (now(), now(), '%s') RETURNING id, created_at, updated_at, name", name))
	return seller[0], err
}
