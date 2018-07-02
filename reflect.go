package mr

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/globalsign/mgo"
)

var reMiddleUpperCase = regexp.MustCompile("(.)([A-Z])")

// CollectionNameFor determines the collection name for the given object. It
// naively converts the to lower case and appends an "s".
func (r MongoRepo) CollectionNameFor(obj interface{}) string {
	typ := reflect.TypeOf(obj)
	name := typ.Name()
	if typ.Kind() == reflect.Slice {
		name = typ.Elem().Name()
	}
	name = reMiddleUpperCase.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(name) + "s"
}

// CollectionFor returns the appropriate collection for the given object, based
// on the return value of "CollectionNameFor".
func (r MongoRepo) CollectionFor(obj interface{}) *mgo.Collection {
	return r.Database.C(r.CollectionNameFor(obj))
}
