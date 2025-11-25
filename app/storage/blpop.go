package storage

import "time"

func (s *Storage) BLPOP(key string, waitTime int) (string, bool) {
	pop := func() (string, bool) {

		s.m.Lock()
		defer s.m.Unlock()
		objectValue, ok := s.Map[key]
		if !ok {
			return "", false
		}
		if len(objectValue.List) == 0 {
			return "", false
		}
		result := objectValue.List[0]
		objectValue.List = objectValue.List[1:]
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
