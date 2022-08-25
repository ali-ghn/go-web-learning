package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	UserId    bson.ObjectId `json:"Id" bson:"Id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
}
