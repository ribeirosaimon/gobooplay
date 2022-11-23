package domain

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Voucher struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Price       primitive.Decimal128 `json:"price" bson:"price"`
	Status      Status               `json:"status" bson:"status"`
	Quantity    float64              `json:"quantity" bson:"quantity"`
	CreatedAt   time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt" bson:"updatedAt"`
	UpdateBy    AccountRef           `bson:"updateBy" bson:"updateBy"`
}

type VoucherDTO struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	Quantity    float64         `json:"quantity"`
}
