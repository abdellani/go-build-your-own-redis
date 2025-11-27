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
	if strings.Compare(end, "+") == 0 {
		return -1, -1
	}
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

func ConvertIdToIntegers(id string) (int, int, error) {
	//TODO: handle invalid ids
	timeStr, seqStr, err := SplitId(id)
	if err != nil {
		return -1, -1, err
	}
	time, _ := strconv.Atoi(timeStr)
	seq, _ := strconv.Atoi(seqStr)
	return time, seq, nil

}

func IsGreater(time1, seq1, time2, seq2 int) bool {
	return (time1 > time2) || (time1 == time2 && seq1 > seq2)
}

func InRange(startTime, startSeq, endTime, endSeq int, item StreamItem) (inRange bool, timeToStop bool) {
	if item.MillieSecondTime < startTime {
		return false, false
	}
	if item.MillieSecondTime == startTime && item.Sequence < startSeq {
		return false, false
	}
	//This means item comes after the start limit

	// when "+" is used, we want to continue to the end
	if endTime == -1 && endSeq == -1 {
		return true, false
	}

	// item's time comes after the end time limit
	if item.MillieSecondTime > endTime {
		return false, true
	}
	// seq == -1 means we want to take all the item in the milliesecondtime regradless of sequence time
	if item.MillieSecondTime == endTime && item.Sequence > endSeq && endSeq != -1 {
		return false, true
	}

	// item is in the targeted range
	return true, false
}
