package mr_test

import (
	"testing"

	"github.com/mhutter/mr"
)

type Foo struct{}
type bar struct{}
type TestFoo struct{}
type testBar struct{}

var v = map[string]interface{}{
	"foos":      Foo{},
	"bars":      make([]bar, 0, 0),
	"test_foos": TestFoo{},
	"test_bars": []testBar{},
}

func TestCollectionNameFor(t *testing.T) {
	r := mr.MongoRepo{}
	for exp, it := range v {
		act := r.CollectionNameFor(it)
		if act != exp {
			t.Errorf("expected '%s', got '%s'", exp, act)
			return
		}
	}
}
