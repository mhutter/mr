package mr

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoCollection interacts with collections
type MongoCollection struct {
	*mgo.Collection
}

// Insert generates an ID for `object`, updates date fields and inserts it into
// `coll`
func (c *MongoCollection) Insert(object Model) error {
	object.generateID()
	now := time.Now()
	object.setCreatedAt(now)
	object.setUpdatedAt(now)
	return c.Collection.Insert(object)
}

// FindAll return all objects in `coll`
func (c *MongoCollection) FindAll(result interface{}) error {
	return c.Collection.Find(nil).All(result)
}

// Find returns one record by its ObjectID. Returns an ErrNoObjectID if "id" is
// not a valid ObjectID.
func (c *MongoCollection) Find(id string, result interface{}) error {
	if !bson.IsObjectIdHex(id) {
		return ErrNoObjectID(id)
	}
	return c.Collection.FindId(bson.ObjectIdHex(id)).One(result)
}

// FindBy returns documents where `key` is set to `value`
func (c *MongoCollection) FindBy(key string, value, result interface{}) error {
	q := bson.M{}
	q[key] = value
	return c.Collection.Find(q).All(result)
}

// FindOneBy returns the first document where `key` is set to `value`
func (c *MongoCollection) FindOneBy(key string, value, result interface{}) error {
	q := bson.M{}
	q[key] = value
	return c.Collection.Find(q).One(result)
}

// Update updates the document with the same ID as the given one
func (c *MongoCollection) Update(object Model) error {
	object.setUpdatedAt(time.Now())
	return c.Collection.UpdateId(object.getID(), object)
}

// Delete removes the document with the same ID as the given one
func (c *MongoCollection) Delete(object Model) error {
	return c.Collection.RemoveId(object.getID())
}
