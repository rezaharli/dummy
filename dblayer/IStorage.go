package dblayer

import (
	"encoding/json"
)

type IStorage interface {
	Saver
	Loader
	Deleter
	GetCaper
	GetEmptyIndexer
	GetIterator
	Filterer
}
type Filter struct {
	Type       string
	FieldName  string
	FieldValue string
}

func (f *Filter) Evaluate(data interface{}) bool {
	if f.Type == "eq" {
		ii, _ := json.Marshal(data)
		kk := map[string]interface{}{}
		json.Unmarshal(ii, &kk)
		if kk[f.FieldName].(string) == f.FieldValue {
			return true
		}
		return false
	}
	return false
}

type IFilter interface {
	Evaluate(data interface{}) bool
}
type Filterer interface {
	Filter(IFilter) []interface{}
}
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
