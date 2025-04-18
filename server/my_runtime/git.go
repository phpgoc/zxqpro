package my_runtime

import "sync"

type stringSet struct {
	set sync.Map
}

func (s *stringSet) Add(value string) {
	s.set.Store(value, struct{}{})
}

func (s *stringSet) Remove(value string) {
	s.set.Delete(value)
}

func (s *stringSet) Each(f func(value string)) {
	s.set.Range(func(key, value interface{}) bool {
		f(key.(string))
		return true
	})
}

var GitPathList stringSet
