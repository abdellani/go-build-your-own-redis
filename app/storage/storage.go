package storage

import (
	"sync"
	"time"
)

type Storage struct {
	Map map[string]string
	m   sync.Mutex
}

func New() *Storage {
	return &Storage{
		Map: map[string]string{},
	}
}
func (s *Storage) Set(key string, value string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.Map[key] = value
}

func (s *Storage) Get(key string) string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.Map[key]
}
func (s *Storage) Delete(key string) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.Map, key)
}

func (s *Storage) Keys() []string {
	s.m.Lock()
	defer s.m.Unlock()
	result := []string{}
	for k, _ := range s.Map {
		result = append(result, k)
	}
	return result
}
func (s *Storage) SetExpiration(key string, millieseconds int) {
	go func() {
		time.Sleep(time.Millisecond * time.Duration(millieseconds))
		s.Delete(key)
	}()
}
