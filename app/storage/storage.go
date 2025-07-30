package storage

import (
	"sync"
	"time"
)

type Storage struct {
	Map map[string]Value
	m   sync.Mutex
}

type Value struct {
	Value          string
	Values         []string
	ExpirationTime time.Time
}

func New() *Storage {
	return &Storage{
		Map: map[string]Value{},
	}
}
func (s *Storage) Set(key string, value Value) {
	s.m.Lock()
	defer s.m.Unlock()
	s.Map[key] = value
}

func (s *Storage) Get(key string) string {
	s.m.Lock()
	defer s.m.Unlock()
	if s.Map[key].ExpirationTime.IsZero() || s.Map[key].ExpirationTime.After(time.Now()) {
		return s.Map[key].Value
	}
	// expiration time is in past
	delete(s.Map, key)
	return ""
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

func (s *Storage) Push(key, value string) int {
	s.m.Lock()
	defer s.m.Unlock()
	valueObject := s.Map[key]
	valueObject.Values = append(valueObject.Values, value)
	s.Map[key] = valueObject
	return len(valueObject.Values)
}
