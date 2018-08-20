package relationships

type Relationships interface {
	SetConnect(tag string, v interface{}) interface{}
	QueryConnect(tag string) interface{}
}
