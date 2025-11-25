package storage

func (s *Storage) LPush(key, value string) int {
	s.m.Lock()
	defer s.m.Unlock()
	prepend := []string{value}
	valueObject := s.Map[key]
	valueObject.List = append(prepend, valueObject.List...)
	s.Map[key] = valueObject
	return len(valueObject.List)
}
