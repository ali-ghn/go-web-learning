package models

import (
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	FirstName string        `json:"firstName" bson:"firstName`
	LastName  string        `json:"lastName bson:"lasttName"`
	UserId    bson.ObjectId `json:"Id"  bson:"Id"`
}
