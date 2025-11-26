package storage

import "github.com/abdellani/go-build-your-own-redis/app/storage/stream"

func (s *Storage) XRange(key, start, end string) []stream.StreamItem {
	s.m.Lock()
	defer s.m.Unlock()
	data := s.Map[key]
	return data.Stream.GetRange(start, end)
}
