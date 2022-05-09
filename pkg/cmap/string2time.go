package cmap

import (
	"sync"
	"time"
)

type string2TimeMap struct {
	m sync.Map
}

func (s *string2TimeMap) Store(key string, t time.Time) {
	s.m.Store(key, t)
}

func (s *string2TimeMap) Load(key string) (value time.Time, ok bool) {
	v, ok := s.m.Load(key)
	if v != nil {
		value = v.(time.Time)
	}
	return
}

func (s *string2TimeMap) Range(f func(key string, value time.Time) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(string), value.(time.Time))
	}
	s.m.Range(f1)
}

func (s *string2TimeMap) Delete(key string) {
	s.m.Delete(key)
}
