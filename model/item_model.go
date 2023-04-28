package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Stock       int                `json:"stock"`
	Seller      primitive.ObjectID `json:"seller"`
}
