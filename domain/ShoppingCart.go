package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ShoppingCart struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Hash      string               `json:"hash" bson:"hash"`
	Product   Product              `json:"product" bson:"product"`
	Price     primitive.Decimal128 `json:"price" bson:"price"`
	Owner     AccountRef           `json:"owner" bson:"owner"`
	Voucher   Voucher              `json:"voucher" bson:"voucher"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time            `json:"updateAt" bson:"updateAt"`
}

func (s *ShoppingCart) getTotalPrice() primitive.Decimal128 {
	return s.Product.Price
}
