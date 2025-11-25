package storage

// Possible types:
//
//	string, list, set, zset, hash, stream, and vectorset
func (s *Storage) Type(key string) string {
	s.m.Lock()
	defer s.m.Unlock()
	value, ok := s.Map[key]
	if !ok {
		return "none"
	}

	return value.Type.String()
}
