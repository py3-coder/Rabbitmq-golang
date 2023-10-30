package models

import "time"

type User struct {
	UserID    int       `bson:"user_id" json:"user_id"`
	Name      string    `bson:"name" json:"name"`
	Mobile    int       `bson:"mobile" json:"mobile"`
	Latitude  float64   `bson:"latitude" json:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
