package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShoppingCart struct {
	ID        primitive.ObjectID   `json:"ID" bson:"ID,omitempty"`
	Hash      string               `json:"Hash" bson:"Hash"`
	Product   Product              `json:"Product" bson:"Product"`
	Price     primitive.Decimal128 `json:"Price" bson:"Price"`
	Owner     AccountRef           `json:"Owner" bson:"Owner"`
	CreatedAt time.Time            `json:"CreatedAt" bson:"CreatedAt"`
	UpdateAt  time.Time            `json:"UpdateAt" bson:"UpdateAt"`
}
