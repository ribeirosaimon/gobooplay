package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/text/currency"
	"time"
)

type Product struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	Price            currency.Amount    `json:"price" bson:"price"`
	SubscriptionTime time.Time          `json:"subscriptionTime" bson:"subscriptionTime"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
