package storage

func (s *Storage) XAdd(streamKey, id string) string {
	s.m.Lock()
	defer s.m.Unlock()
	_, ok := s.Map[streamKey]

	if !ok {
		value := Value{
			Type: STREAM_TYPE,
		}
		s.Map[streamKey] = value
	}
	return id
}
