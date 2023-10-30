package models

import "time"

type Product struct {
	UserID                  int       `bson:"user_id" json:"user_id"`
	ProductName             string    `bson:"product_name" json:"product_name"`
	ProductDescription      string    `bson:"product_description" json:"product_description"`
	ProductImages           []string  `bson:"product_images" json:"product_images"`
	ProductPrice            int       `bson:"product_price" json:"product_price"`
	CompressedProductImages []string  `bson:"compressed_product_images" json:"compressed_product_images"`
	CreatedAt               time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt               time.Time `bson:"updated_at" json:"updated_at"`
}

type ProductResponse struct {
	StatusCode int    `json:"statuscode"`
	Status     string `json:"status"`
}

type ProductDetails struct {
	UserID                  int       `bson:"user_id" json:"user_id"`
	ProductName             string    `bson:"product_name" json:"product_name"`
	ProductDescription      string    `bson:"product_description" json:"product_description"`
	ProductPrice            int       `bson:"product_price" json:"product_price"`
	ProductID               int       `bson:"product_id" json:"product_id"`
	CompressedProductImages []string  `bson:"compressed_product_images" json:"compressed_product_images"`
	ProductImages           []string  `bson:"product_images" json:"product_images"`
	CreatedAt               time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt               time.Time `bson:"updated_at" json:"updated_at"`
}
