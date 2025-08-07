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

func (s *Storage) RPush(key, value string) int {
	s.m.Lock()
	valueObject := s.Map[key]
	valueObject.Values = append(valueObject.Values, value)
	s.Map[key] = valueObject
	s.m.Unlock()
	go s.UnblockWaiting(key)
	return len(valueObject.Values)
}

func (s *Storage) UnblockWaiting(key string) {
	s.m.Lock()
	defer s.m.Unlock()
	valueObject := s.Map[key]
	if len(valueObject.Blocked) == 0 {
		return
	}
	blocked := valueObject.Blocked[0]
	valueObject.Blocked = valueObject.Blocked[1:]
	s.Map[key] = valueObject
	blocked <- struct{}{}

}

func (s *Storage) LPush(key, value string) int {
	s.m.Lock()
	defer s.m.Unlock()
	prepend := []string{value}
	valueObject := s.Map[key]
	valueObject.Values = append(prepend, valueObject.Values...)
	s.Map[key] = valueObject
	return len(valueObject.Values)
}

func (s *Storage) Range(key string, start, end int) []string {
	s.m.Lock()
	defer s.m.Unlock()
	list := s.Map[key].Values
	result := []string{}
	if start < 0 {
		if -start > len(list) {
			start = 0
		} else {

			start += len(list)
		}
	}
	if end < 0 {
		end += len(list)
	}
	for i := start; i <= end && i < len(list); i++ {
		result = append(result, list[i])
	}
	return result
}

func (s *Storage) LLen(key string) int {
	s.m.Lock()
	defer s.m.Unlock()

	list, exists := s.Map[key]
	if !exists {
		return 0
	}
	return len(list.Values)
}

func (s *Storage) LPop(key string, num int) []string {
	s.m.Lock()
	defer s.m.Unlock()
	objectValue := s.Map[key]
	result := objectValue.Values[0:num]
	objectValue.Values = objectValue.Values[num:]
	s.Map[key] = objectValue
	return result
}

func (s *Storage) BLPOP(key string, waitTime int) (string, bool) {
	pop := func() (string, bool) {

		s.m.Lock()
		defer s.m.Unlock()
		objectValue, ok := s.Map[key]
		if !ok {
			return "", false
		}
		if len(objectValue.Values) == 0 {
			return "", false
		}
		result := objectValue.Values[0]
		objectValue.Values = objectValue.Values[1:]
		s.Map[key] = objectValue
		return result, true
	}
	wait := func() <-chan struct{} {
		s.m.Lock()
		object := s.Map[key]
		pause := make(chan struct{})
		object.Blocked = append(object.Blocked, pause)
		s.Map[key] = object
		s.m.Unlock()
		return pause
	}
	timer := time.After(time.Duration(waitTime) * time.Millisecond)
	if waitTime == 0 {
		timer = time.After(time.Duration(1000000) * time.Hour)

	}
	for {
		result, ok := pop()
		if ok {
			return result, true
		}
		select {
		case <-wait():
			continue
		case <-timer:
			return "", false
		}

	}

}
