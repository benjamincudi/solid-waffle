package models

import "time"

type DefaultDBFields struct {
	ID        int       `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Seller struct {
	DefaultDBFields `bulkprocess:"substruct"`
	Name            string `db:"name" json:"name"`
}

type Product struct {
	DefaultDBFields `bulkprocess:"substruct"`
	Name            string `db:"name" json:"name"`
	SellerID        int    `db:"seller_id" json:"seller_id"`
}

type Variation struct {
	DefaultDBFields `bulkprocess:"substruct"`
	Name            string `db:"name" json:"name"`
}

type ProductVariation struct {
	DefaultDBFields `bulkprocess:"substruct"`
	ProductID       int   `db:"product_id" json:"product_id"`
	VariationIDs    []int `db:"variation_ids" json:"variation_ids"`
	IsAvailable     bool  `db:"is_available" json:"is_available"`
	Price           int   `db:"price" json:"price"`
}
