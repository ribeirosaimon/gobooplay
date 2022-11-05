package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID               primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name             string               `json:"name" bson:"name"`
	Price            primitive.Decimal128 `json:"price" bson:"price"`
	SubscriptionTime uint8                `json:"subscriptionTime" bson:"subscriptionTime"`
	Status           Status               `json:"status" bson:"status"`
	CreatedAt        time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time            `json:"updatedAt" bson:"updatedAt"`
	UpdateBy         LoggedUser           `json:"updateBy" bson:"updateBy"`
}

type ProductDTO struct {
	ID               string `json:"id,omitempty" bson:"_id,omitempty"`
	Name             string `json:"name" bson:"name"`
	Price            string `json:"price" bson:"price"`
	SubscriptionTime uint8  `json:"subscriptionTime" bson:"subscriptionTime"`
}
