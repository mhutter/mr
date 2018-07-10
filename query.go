package mr

import (
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Insert generates an ID for `object`, updates date fields and inserts it into
// `coll`
func (r MongoRepo) Insert(object Model) error {
	object.generateID()
	now := time.Now()
	object.setCreatedAt(now)
	object.setUpdatedAt(now)
	return r.CollectionFor(object).Insert(object)
}

// FindAll return all objects in `coll`
func (r MongoRepo) FindAll(result interface{}) error {
	return r.CollectionFor(result).Find(nil).All(result)
}

// Find all documents based on query
func (r MongoRepo) Find(query bson.M, result interface{}) error {
	return r.CollectionFor(result).Find(query).All(result)
}

// FindOne returns the first item selected by "query"
func (r MongoRepo) FindOne(query bson.M, result interface{}) error {
	return r.CollectionFor(result).Find(query).One(result)
}

// FindID returns one record by its ObjectID. Returns an ErrNoObjectID if "id"
// is not a valid ObjectID.
func (r MongoRepo) FindID(id string, result interface{}) error {
	if !bson.IsObjectIdHex(id) {
		return ErrNoObjectID(id)
	}
	return r.CollectionFor(result).FindId(bson.ObjectIdHex(id)).One(result)
}

// Update updates the document with the same ID as the given one
func (r MongoRepo) Update(object Model) error {
	object.setUpdatedAt(time.Now())
	return r.CollectionFor(object).UpdateId(object.getID(), object)
}

// UpdateID updates the document with the given ID. It makes sure the ID is not
// changed and the "updated_at" is updated.
//
// Returns an ErrNoObjectID if "id" is not a valid ObjectID.
func (r MongoRepo) UpdateID(id string, object Model) error {
	if !bson.IsObjectIdHex(id) {
		return ErrNoObjectID(id)
	}
	object.setID(bson.ObjectIdHex(id))
	return r.Update(object)
}

// Delete removes the document with the same ID as the given one.
//
// Note that there is no "DeleteID" method as we need the object to determine
// the collection name anyway...
func (r MongoRepo) Delete(object Model) error {
	return r.CollectionFor(object).RemoveId(object.getID())
}

// EnsureUnique creates unique keys for the given Model
func (r MongoRepo) EnsureUnique(object Model, keys []string) {
	coll := r.CollectionFor(object)
	if err := coll.EnsureIndex(mgo.Index{
		Key:        keys,
		Unique:     true,
		DropDups:   true,
		Background: true,
	}); err != nil {
		log.Printf(
			"ERROR creating unique index in '%s' for '%s': %s\n",
			coll.Name,
			keys,
			err,
		)
	}
}
