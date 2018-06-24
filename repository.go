package mr

import (
	"github.com/globalsign/mgo"
)

// MongoRepo interacts with MongoDB
type MongoRepo struct {
	*mgo.Database
}

// Connect dials MongoDB and returns the configured Repository
func Connect(url string) (Repository, error) {
	info, err := mgo.ParseURL(url)
	if err != nil {
		return nil, ErrInvalidURL(url)
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}

	repo := &MongoRepo{session.DB(info.Database)}
	return repo, nil
}

// C returns a Collection
func (r *MongoRepo) C(name string) Collection {
	return &MongoCollection{r.Database.C(name)}
}
