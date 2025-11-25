package storage

import (
	"fmt"
	"sync"
	"time"
)

type Storage struct {
	Map map[string]Data
	m   sync.Mutex
}

type Data struct {
	Value          string
	List           []string
	Stream         Stream
	ExpirationTime time.Time
	Blocked        []chan struct{}
	Type           Type
}

type StreamItem struct {
	MillieSecondTime int
	Sequence         int
	Value            map[string]string
}
type Stream struct {
	Items []StreamItem
}

func (s *Stream) isValidId(time, seq int) bool {
	if time == 0 && seq == 0 {
		return false
	}
	if len(s.Items) == 0 {
		return true
	}
	topItem := s.Items[len(s.Items)-1]
	topTime := topItem.MillieSecondTime
	topSeq := topItem.Sequence

	if (time < topTime) ||
		((time == topTime) && (seq <= topSeq)) {
		return false
	}
	return true
}

func (s *Stream) AddItem(time, seq int) {
	s.Items = append(s.Items, StreamItem{
		MillieSecondTime: time,
		Sequence:         seq,
	})
}

func (s *Stream) GetTop() *StreamItem {
	return &s.Items[len(s.Items)-1]
}
func (s *Stream) generateNextSeq(time int) int {
	if len(s.Items) == 0 {
		if time == 0 {
			return 1
		} else {
			return 0
		}
	}
	top := s.GetTop()
	if time == top.MillieSecondTime {
		return top.Sequence + 1
	}
	return 0
}

func (s *Stream) getMillieseconds() int {
	return int(time.Now().UnixMilli())

}

func (s *Stream) generateFullId() (int, int) {
	millieSecondsTime := s.getMillieseconds()
	sequence := s.generateNextSeq(millieSecondsTime)
	return millieSecondsTime, sequence
}

func (s *Stream) GetTopId() string {
	top := s.GetTop()
	return fmt.Sprintf("%d-%d", top.MillieSecondTime, top.Sequence)
}
func NewDataStream() *Data {
	return &Data{
		Type:   TYPE_STREAM,
		Stream: Stream{},
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
