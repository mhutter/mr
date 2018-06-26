package mr

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

const defaultURL = "mongodb://localhost/"

// MongoRepo interacts with MongoDB
type MongoRepo struct {
	*mgo.Database
}

// Autoconnect tries to determine its DB URI from looking at the MONGODB_URI
// environment variable. If it is not set or empty it falls back to connecting
// to localhost, using "fallbackDBName" as the Database to use.
func Autoconnect(fallbackDBName string) (Repository, error) {
	url := os.Getenv("MONGODB_URI")
	if len(url) < 1 {
		url = defaultURL + fallbackDBName
		log.Println("MONGODB_URI not set, connecting to fallback DB:", url)
	}
	return Connect(url)
}

// MustAutoconnect works similar to Autoconnect, but aborts the programm if
// connection fails.
func MustAutoconnect(fallbackDBName string) Repository {
	repo, err := Autoconnect(fallbackDBName)
	if err != nil {
		log.Fatalln("MongoRepo: Could not connect:", err.Error())
	}
	return repo
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
