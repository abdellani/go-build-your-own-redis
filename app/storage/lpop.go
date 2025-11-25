package storage

func (s *Storage) LPop(key string, num int) []string {
	s.m.Lock()
	defer s.m.Unlock()
	objectValue := s.Map[key]
	result := objectValue.Values[0:num]
	objectValue.Values = objectValue.Values[num:]
	s.Map[key] = objectValue
	return result
}
