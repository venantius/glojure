package lang

// This is a file for things that don't natively exist in Go but do in Java. For the time being I've created stub implementations of the Java version until I have a chance to figure out an idiomatic equivalent in Go.

type Collection struct{}

func (c *Collection) IsEmpty() bool {
	return true
}

// TODO: Remove this
type Iterator struct{}
type Iterable interface{}

// TODO: All of this is tied to java.util.list ;; is there an analog for this in Go that we'd actually use?
type List struct{}

func (l *List) Size() int {
	return 0
}

func (l *List) ToArray() []interface{} {
	return make([]interface{}, 5)
}

func (l *List) Get(i int) interface{} {
	return 5
}
