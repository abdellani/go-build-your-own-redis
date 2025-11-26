package stream

import (
	"fmt"
)

type StreamItem struct {
	MillieSecondTime int
	Sequence         int
	Values           []string
}

func (s *StreamItem) IdString() string {
	return fmt.Sprintf("%d-%d", s.MillieSecondTime, s.Sequence)
}
