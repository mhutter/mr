package mr

import "time"

// Model must be implemented by types that should be inserted into the
// Repository
type Model interface {
	getID() string
	generateID()
	setCreatedAt(time.Time)
	setUpdatedAt(time.Time)
}

// Repository interacts with the actual database
type Repository interface {
	C(name string) Collection
}

// Collection represents a MongoDB collection
type Collection interface {
	Insert(object Model) error
	FindAll(result interface{}) error
	Find(id, result interface{}) error
	FindBy(key string, value, result interface{}) error
	FindOneBy(key string, value, result interface{}) error
	Update(object Model) error
	Delete(object Model) error
}
