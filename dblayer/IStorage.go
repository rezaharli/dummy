package dblayer

type Saver interface {
	Save(interface{}) (int, error)
}
type Loader interface {
	Load(id interface{}) (interface{}, error)
}
type Deleter interface {
	Delete(id interface{}) error
}
type GetCaper interface {
	GetCap() int
}
type GetEmptyIndexer interface {
	GetEmptyIndex() int
}
type Iterator interface {
	Next() interface{}
	HasNext() bool
}
type GetIterator interface {
	GetIterator() Iterator
}
