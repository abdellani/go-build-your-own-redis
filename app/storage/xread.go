package storage

import "github.com/abdellani/go-build-your-own-redis/app/storage/stream"

func (s *Storage) XRead(keys, ids []string) [][]stream.StreamItem {
	s.m.Lock()
	defer s.m.Unlock()
	streamItems := [][]stream.StreamItem{}
	for i, key := range keys {
		data := s.Map[key]
		id := ids[i]
		streamItems = append(streamItems, data.Stream.GetItem(id))

	}
	return streamItems
}
