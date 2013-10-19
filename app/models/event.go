package models

import (
	"labix.org/v2/mgo/bson"
)

type Event struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Date time.Time
	Type string
}
