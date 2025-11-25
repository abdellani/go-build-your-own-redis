package storage

func (s *Storage) Range(key string, start, end int) []string {
	s.m.Lock()
	defer s.m.Unlock()
	list := s.Map[key].List
	result := []string{}
	if start < 0 {
		if -start > len(list) {
			start = 0
		} else {

			start += len(list)
		}
	}
	if end < 0 {
		end += len(list)
	}
	for i := start; i <= end && i < len(list); i++ {
		result = append(result, list[i])
	}
	return result
}
