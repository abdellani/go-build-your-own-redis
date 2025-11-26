package stream

import (
	"errors"
	"strconv"
	"strings"
)

func ParseRangeStart(start string) (int, int) {
	if strings.Compare(start, "-") == 0 {
		return 0, 0
	}
	var startTime int
	var startSeq int
	if DoesIdIncludeSequenceNumber(start) {
		startTime, startSeq = ConvertTwoIdPartsToInt(start)
	} else {
		startTime, _ = strconv.Atoi(start)
		startSeq = 0
	}
	return startTime, startSeq
}

func ParseRangeEnd(end string) (int, int) {
	var endTime int
	var endSeq int
	if DoesIdIncludeSequenceNumber(end) {
		endTime, endSeq = ConvertTwoIdPartsToInt(end)
	} else {
		endTime, _ = strconv.Atoi(end)
		endSeq = -1
	}
	return endTime, endSeq
}
func ConvertTwoIdPartsToInt(id string) (int, int) {
	timeStr, seqStr, _ := SplitId(id)
	time, _ := strconv.Atoi(timeStr)
	seq, _ := strconv.Atoi(seqStr)
	return time, seq
}

func DoesIdIncludeSequenceNumber(id string) bool {
	return strings.Contains(id, "-")
}

func SplitId(id string) (string, string, error) {
	items := strings.Split(id, "-")
	if len(items) != 2 {
		return "", "", errors.New("")
	}
	return items[0], items[1], nil
}
