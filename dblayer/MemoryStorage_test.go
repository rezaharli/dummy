package dblayer

import (
	"testing"
)

type DummyStruct struct {
	Id int
}

func TestSave(t *testing.T) {
	base := NewMemStorage(2)
	ii, err := base.Save(DummyStruct{0})
	if err != nil {
		t.Fail()
	}
	if ii != 0 {
		t.Log("Index salah")
		t.Fail()
	}
	ii, err = base.Save(DummyStruct{1})
	if err != nil {
		t.Fail()
	}
	if ii != 1 {
		t.Log("Index salah")
		t.Fail()
	}
	ii, err = base.Save(DummyStruct{3})
	if err == nil {
		t.Log("Seharusnya error")
		t.Fail()
	}

}
func TestLoad(t *testing.T) {
	base := NewMemStorage(2)
	ii, err := base.Save(DummyStruct{0})
	if err != nil {
		t.Log("Gagal Save", err.Error())
		t.Fail()
	}
	if ii != 0 {
		t.Log("Index salah")
		t.Fail()
	}
	ii, err = base.Save(DummyStruct{1})
	if err != nil {
		t.Log("Gagal Save", err.Error())
		t.Fail()
	}
	if ii != 1 {
		t.Log("Index salah")
		t.Fail()
	}
	returned, err := base.Load(0)
	if err != nil {
		t.Fail()
	}
	if (returned.(DummyStruct)).Id != 0 {
		t.Log("Gagal")
		t.Fail()
	}
	returned, err = base.Load(1)
	if err != nil {
		t.Fail()
	}
	if (returned.(DummyStruct)).Id != 1 {
		t.Log("Gagal")
		t.Fail()
	}

}
func TestDelete(t *testing.T) {
	base := NewMemStorage(2)
	ii, err := base.Save(DummyStruct{0})
	if err != nil {
		t.Log("Gagal Save", err.Error())
		t.Fail()
	}
	if ii != 0 {
		t.Log("Index salah")
		t.Fail()
	}
	ii, err = base.Save(DummyStruct{1})
	if err != nil {
		t.Log("Gagal Save", err.Error())
		t.Fail()
	}
	if ii != 1 {
		t.Log("Index salah")
		t.Fail()
	}
	err = base.Delete(0)
	if err != nil {
		t.Log("Gagal Delete", err.Error())
		t.Fail()
	}
	emp := base.GetEmptyIndex()
	if emp != 0 {
		t.Log("salah")
		t.Fail()
	}
	ii, err = base.Save(DummyStruct{2})
	if err != nil {
		t.Log("Gagal Save", err.Error())
		t.Fail()
	}
	if ii != 0 {
		t.Log("Salah index")
		t.Fail()
	}
	returned, err := base.Load(0)
	if err != nil {
		t.Log("Gagal Get", err.Error())
		t.Fail()
	}
	if (returned.(DummyStruct)).Id != 2 {
		t.Log("Gagal")
		t.Fail()
	}
}
func TestGetEmptyOnFull(t *testing.T) {
	base := NewMemStorage(2)
	base.Save(DummyStruct{0})
	base.Save(DummyStruct{1})
	ii := base.GetEmptyIndex()
	if ii != -1 {
		t.Fail()
	}
}
func TestIterator(t *testing.T) {
	base := NewMemStorage(2)
	cap1 := base.GetCap()
	if cap1 != 2 {
		t.Fail()
	}
	base.Save(DummyStruct{0})
	base.Save(DummyStruct{1})
	iterator := base.GetIterator()
	counter := 0
	for iterator.HasNext() {
		item1 := iterator.Next()
		if item1.(DummyStruct).Id != counter {
			t.Log("Salah counter")
			t.Fail()
		}
		counter++
	}
	base.Delete(0)
	iterator = base.GetIterator()
	counter = 1
	for iterator.HasNext() {
		item1 := iterator.Next()
		if item1.(DummyStruct).Id != counter {
			t.Log("Salah counter")
			t.Fail()
		}
		counter++
	}
}
