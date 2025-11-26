package storage

import (
	"sync"
	"time"

	"github.com/abdellani/go-build-your-own-redis/app/storage/stream"
)

type Storage struct {
	Map map[string]Data
	m   sync.Mutex
}

type Data struct {
	Value          string
	List           []string
	Stream         stream.Stream
	ExpirationTime time.Time
	Blocked        []chan struct{}
	Type           Type
}

func NewDataStream() *Data {
	return &Data{
		Type:   TYPE_STREAM,
		Stream: stream.Stream{},
	}
}

type Type int

const (
	TYPE_UNKOWN Type = iota - 1
	TYPE_STRING
	TYPE_STREAM
)

func (t Type) String() string {
	switch t {
	case TYPE_STRING:
		return "string"
	case TYPE_STREAM:
		return "stream"
	default:
		return "none"
	}
}

func New() *Storage {
	return &Storage{
		Map: map[string]Data{},
	}
}

func (s *Storage) UnblockWaitingWithoutLock(key string) {
	valueObject := s.Map[key]
	if len(valueObject.Blocked) == 0 {
		return
	}
	blocked := valueObject.Blocked[0]
	valueObject.Blocked = valueObject.Blocked[1:]
	s.Map[key] = valueObject
	blocked <- struct{}{}

}
