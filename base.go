package mr

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Base contains common attributes that are common to all model structs that
// are to be stored in the database. It also implements the Model interface,
// which is used when interacting with the repository.
//
// See package documentation for usage.
type Base struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at,omitempty"`
}

func (b *Base) getID() bson.ObjectId { return b.ID }
func (b *Base) generateID()          { b.ID = bson.NewObjectId() }

func (b *Base) setCreatedAt(t time.Time) { b.CreatedAt = t }
func (b *Base) setUpdatedAt(t time.Time) { b.UpdatedAt = t }
