package lang

type Indexed interface {
	Counted

	Nth(i int, notFound interface{}) interface{}
}
