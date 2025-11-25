package storage

func (s *Storage) LLen(key string) int {
	s.m.Lock()
	defer s.m.Unlock()

	list, exists := s.Map[key]
	if !exists {
		return 0
	}
	return len(list.Values)
}
