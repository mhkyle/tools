package cmap

import (
	"testing"
	"time"
)

func Test_string2TimeMap_Load(t *testing.T) {
	var m string2TimeMap
	m.Store("1", time.Now())
	val, ok := m.Load("1")
	if ok {
		t.Logf("val is %v\n", val)
	}
}

func Test_string2TimeMap_Store(t *testing.T) {
	var m string2TimeMap
	m.Store("1", time.Now())
	val, ok := m.Load("1")
	if ok {
		t.Logf("val is %v\n", val)
	}
}

func Test_string2TimeMap_Delete(t *testing.T) {
	var m string2TimeMap
	m.Store("1", time.Now())
	val, ok := m.Load("1")
	if ok {
		t.Logf("val is %v\n", val)
	}

	m.Delete("1")

	val, ok = m.Load("1")
	if ok {
		t.Fatalf("1 deleted, it should not exist %v", val)
	} else {
		t.Logf("1 deleted, not existed any more")
	}
}

func Test_string2TimeMap_Range(t *testing.T) {
	var m string2TimeMap
	m.Store("1", time.Now())
	time.Sleep(time.Second)
	m.Store("2", time.Now())

	m.Range(func(key string, val time.Time) bool {
		t.Logf("item: ")
		t.Logf("key: %s, value: %v\n", key, val)
		return true
	})

}

