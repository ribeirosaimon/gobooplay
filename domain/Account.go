package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	MALE gender = iota
	FEMALE
)

type gender int

type Account struct {
	Id                 primitive.ObjectID `json:"id" bson:"_id"`
	Name               string             `json:"name" bson:"name"`
	Password           string             `json:"password" bson:"password"`
	FamilyName         string             `json:"familyName" bson:"familyName"`
	Login              string             `json:"login"bson:"login"`
	Gender             gender             `json:"gender"bson:"gender"`
	Role               []string           `json:"role" bson:"role"`
	LastLogin          time.Time          `json:"lastLogin" bson:"lastLogin"`
	LastLoginAttemp    time.Time          `json:"lastLoginAttemp" bson:"lastLoginAttemp"`
	PasswordErrorCount uint32             `json:"passwordErrorCount" bson:"passwordErrorCount"`
	LoginCount         uint32             `json:"login_count" bson:"loginCount"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt"`
}
