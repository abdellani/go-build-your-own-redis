package storage

func (s *Storage) RPush(key, value string) int {
	s.m.Lock()
	defer s.m.Unlock()
	valueObject := s.Map[key]
	valueObject.List = append(valueObject.List, value)
	s.Map[key] = valueObject
	s.UnblockWaitingWithoutLock(key)
	return len(valueObject.List)
}
