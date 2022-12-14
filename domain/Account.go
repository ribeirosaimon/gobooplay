package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Account struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name               string             `json:"name" bson:"name"`
	Password           string             `json:"password" bson:"password"`
	FamilyName         string             `json:"familyName" bson:"familyName"`
	Login              string             `json:"login" bson:"login"`
	Gender             Gender             `json:"gender" bson:"gender"`
	Status             Status             `json:"status" bson:"status"`
	Role               []Role             `json:"role" bson:"role"`
	LastLogin          time.Time          `json:"lastLogin" bson:"lastLogin"`
	LastLoginAttemp    time.Time          `json:"lastLoginAttemp" bson:"lastLoginAttemp"`
	PasswordErrorCount uint32             `json:"passwordErrorCount" bson:"passwordErrorCount"`
	LoginCount         uint32             `json:"login_count" bson:"loginCount"`
	CreatedAt          time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt          time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type AccountRef struct {
	Name   string `json:"name" bson:"name"`
	UserId string `json:"userId" bson:"userId"`
}

type AccountDTO struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	FamilyName string `json:"familyName"`
	Gender     Gender `json:"gender"`
}

type LoginDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoggedUser struct {
	Login  string `json:"username" bson:"username"`
	UserId string `json:"userId" bson:"userId"`
	Role   []Role `json:"role" bson:"role"`
}

type UserAccessToken struct {
	Token string `json:"access_token"`
}

func (a AccountDTO) cleanPassword() AccountDTO {
	a.Password = ""
	return a
}

func (a Account) MyRef() AccountRef {
	return AccountRef{Name: a.Name, UserId: a.ID.Hex()}
}

func (a Account) GetLoggedUser() LoggedUser {
	return LoggedUser{
		Login:  a.Login,
		UserId: a.ID.Hex(),
		Role:   a.Role,
	}
}
