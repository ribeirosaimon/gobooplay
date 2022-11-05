package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Subscription struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Owner     AccountRef         `json:"owner" bson:"owner"`
	Product   Product            `json:"product" bson:"product"`
	Status    Status             `json:"status" bson:"status"`
	BegginAt  time.Time          `json:"startedAt" bson:"startedAt"`
	EndAt     time.Time          `json:"endAt" bson:"endAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
