package storage

func (s *Storage) Keys() []string {
	s.m.Lock()
	defer s.m.Unlock()
	result := []string{}
	for k, _ := range s.Map {
		result = append(result, k)
	}
	return result
}
