package storage

func (s *Storage) Delete(key string) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.Map, key)
}
