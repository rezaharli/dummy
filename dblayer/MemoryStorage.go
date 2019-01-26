package dblayer

import (
	"errors"
	"fmt"
)

type MemStorage struct {
	Storage []interface{}
}
type MemStorageIterator struct {
	storage *[]interface{}
	idx     int
}

var (
	STORAGE_FULL_ERROR       = errors.New("No More Storage")
	INDEX_OUT_OF_BOUND_ERROR = errors.New("Index out of range")
	INDEX_IS_NOT_INTEGER     = errors.New("Index is not integer")
	INDEX_IS_EMPTY           = errors.New("Index is Empty")
)

func NewMemStorage(size int) MemStorage {
	storage := MemStorage{}
	storage.Storage = make([]interface{}, size)
	return storage
}

func (iter *MemStorageIterator) Next() interface{} {
	for i := iter.idx + 1; i < len(*iter.storage); i++ {
		if (*iter.storage)[i] != nil {
			iter.idx = i
			return (*iter.storage)[i]
		}
	}
	return nil
}
func (iter *MemStorageIterator) HasNext() bool {
	for i := iter.idx + 1; i < len(*iter.storage); i++ {
		if (*iter.storage)[i] != nil {
			return true
		}
	}
	return false
}
func (m *MemStorage) GetEmptyIndex() int {
	for idx, _ := range m.Storage {
		if m.Storage[idx] == nil {
			return idx
		}
	}
	return -1
}
func (m *MemStorage) Save(data interface{}) (int, error) {
	for idx, _ := range m.Storage {
		if m.Storage[idx] == nil {
			m.Storage[idx] = data
			return idx, nil
		}
	}
	return -1, STORAGE_FULL_ERROR
}
func (m *MemStorage) isIndexValid(id interface{}) (int, error) {
	if _, ok := id.(int); !ok {
		return -1, INDEX_IS_NOT_INTEGER
	}
	idx := id.(int)
	if len(m.Storage) <= idx {
		return -1, INDEX_OUT_OF_BOUND_ERROR
	}
	return idx, nil
}
func (m *MemStorage) Load(id interface{}) (interface{}, error) {
	idx, err := m.isIndexValid(id)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	data := m.Storage[idx]
	if data == nil {
		return nil, INDEX_IS_EMPTY
	}
	return data, nil
}
func (m *MemStorage) Delete(id interface{}) error {
	idx, err := m.isIndexValid(id)
	if err != nil {
		return err
	}
	m.Storage[idx] = nil
	return nil
}
