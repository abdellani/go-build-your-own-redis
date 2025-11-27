package stream

import (
	"fmt"
	"time"
)

type Stream struct {
	Items []StreamItem
}

func (s *Stream) IsValidId(time, seq int) bool {
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

func (s *Stream) AddItem(time, seq int, values []string) {
	s.Items = append(s.Items, StreamItem{
		MillieSecondTime: time,
		Sequence:         seq,
		Values:           values,
	})
}

func (s *Stream) GetTop() *StreamItem {
	return &s.Items[len(s.Items)-1]
}

func (s *Stream) GenerateNextSeq(time int) int {
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

func (s *Stream) GetMillieseconds() int {
	return int(time.Now().UnixMilli())

}

func (s *Stream) GenerateFullId() (int, int) {
	millieSecondsTime := s.GetMillieseconds()
	sequence := s.GenerateNextSeq(millieSecondsTime)
	return millieSecondsTime, sequence
}

func (s *Stream) GetTopId() string {
	top := s.GetTop()
	return fmt.Sprintf("%d-%d", top.MillieSecondTime, top.Sequence)
}

func (s *Stream) GetRange(start, end string) []StreamItem {

	var result []StreamItem

	startTime, startSeq := ParseRangeStart(start)
	endTime, endSeq := ParseRangeEnd(end)

	for i := 0; i < len(s.Items); i++ {
		item := s.Items[i]
		inRange, timeToStop := InRange(startTime, startSeq, endTime, endSeq, item)
		if timeToStop {
			break
		}
		if !inRange {
			continue
		}
		result = append(result, item)
	}
	return result
}

func (s *Stream) GetItem(id string) []StreamItem {
	time, seq, _ := ConvertIdToIntegers(id)
	items := []StreamItem{}
	for _, item := range s.Items {
		if !IsGreater(item.MillieSecondTime, item.Sequence, time, seq) {
			continue
		}
		items = append(items, item)
	}
	return items
}
