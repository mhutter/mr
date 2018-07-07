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
	name := realTypeOf(obj).Name()
	name = reMiddleUpperCase.ReplaceAllString(name, "${1}_${2}")
	return strings.ToLower(name) + "s"
}

// CollectionFor returns the appropriate collection for the given object, based
// on the return value of "CollectionNameFor".
func (r MongoRepo) CollectionFor(obj interface{}) *mgo.Collection {
	return r.Database.C(r.CollectionNameFor(obj))
}

var kindsWithElem = []reflect.Kind{
	reflect.Array,
	reflect.Chan,
	reflect.Ptr,
	reflect.Slice,
}

func realTypeOf(obj interface{}) reflect.Type {
	typ := reflect.TypeOf(obj)
	return unwrapType(typ)
}
func unwrapType(typ reflect.Type) reflect.Type {
	for _, kind := range kindsWithElem {
		if typ.Kind() == kind {
			return unwrapType(typ.Elem())
		}
	}

	return typ
}
