package storage

func (s *Storage) LPush(key, value string) int {
	s.m.Lock()
	defer s.m.Unlock()
	prepend := []string{value}
	valueObject := s.Map[key]
	valueObject.Values = append(prepend, valueObject.Values...)
	s.Map[key] = valueObject
	return len(valueObject.Values)
}
