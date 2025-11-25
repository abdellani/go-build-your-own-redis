package storage

func (s *Storage) LPop(key string, num int) []string {
	s.m.Lock()
	defer s.m.Unlock()
	objectValue := s.Map[key]
	result := objectValue.List[0:num]
	objectValue.List = objectValue.List[num:]
	s.Map[key] = objectValue
	return result
}
