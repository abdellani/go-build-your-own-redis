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
	Blocked        []chan struct{}
}

func New() *Storage {
	return &Storage{
		Map: map[string]Value{},
	}
}

func (s *Storage) UnblockWaitingWithoutLock(key string) {
	valueObject := s.Map[key]
	if len(valueObject.Blocked) == 0 {
		return
	}
	blocked := valueObject.Blocked[0]
	valueObject.Blocked = valueObject.Blocked[1:]
	s.Map[key] = valueObject
	blocked <- struct{}{}

}
