package models

import (
	"labix.org/v2/mgo/bson"
)

type Message struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Subject     string
	Content     string
	ContentHTML string
	Delay       string
}
