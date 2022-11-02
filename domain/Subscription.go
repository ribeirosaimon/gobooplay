package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	ACTIVE   = "ACTIVE"
	DISABLED = "DISABLED"
	PAUSE    = "PAUSE"
	TRIAL    = "TRIAL"
)

type Subscription struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Owner     AccountRef         `json:"owner" bson:"owner"`
	Product   Product            `json:"product" bson:"product"`
	BegginAt  time.Time          `json:"startedAt" bson:"startedAt"`
	EndAt     time.Time          `json:"endAt" bson:"endAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
