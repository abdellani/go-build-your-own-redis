package storage

import "github.com/abdellani/go-build-your-own-redis/app/storage/stream"

func (s *Storage) XRead(key, id string) []stream.StreamItem {
	s.m.Lock()
	defer s.m.Unlock()
	data := s.Map[key]
	streamItem := data.Stream.GetItem(id)
	return streamItem
}
