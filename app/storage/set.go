package storage

func (s *Storage) Set(key string, value Data) {
	s.m.Lock()
	defer s.m.Unlock()
	s.Map[key] = value
}
