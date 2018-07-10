package mr

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Model must be implemented by types that should be inserted into the
// Repository
type Model interface {
	getID() bson.ObjectId
	setID(bson.ObjectId)
	generateID()
	setCreatedAt(time.Time)
	setUpdatedAt(time.Time)
}

// Repository interacts with the actual database
type Repository interface {
	Insert(object Model) error
	FindAll(result interface{}) error
	Find(query bson.M, result interface{}) error
	FindOne(query bson.M, result interface{}) error
	FindID(id string, result interface{}) error
	Update(object Model) error
	UpdateID(id string, object Model) error
	Delete(object Model) error
	EnsureUnique(object Model, keys []string)
}
