package storage

import "time"

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
