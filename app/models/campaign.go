package models

import (
	"labix.org/v2/mgo/bson"
)

type Campaign struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string
	Messages     []Message
	TriggerEvent string
}
