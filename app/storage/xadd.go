package storage

import (
	"errors"
	"strconv"
	"strings"
)

const (
	XADD_ERROR_ID_MUST_BE_GREATER_THAN_0_0        = "ERR The ID specified in XADD must be greater than 0-0"
	XADD_ERROR_ID_IS_EQUAL_OR_SMALL_TO_STREAM_TOP = "ERR The ID specified in XADD is equal or smaller than the target stream top item"
)

func (s *Storage) XAdd(streamKey, id string) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()

	if strings.Compare("0-0", id) == 0 {
		return "", errors.New(XADD_ERROR_ID_MUST_BE_GREATER_THAN_0_0)
	}
	millieSecondsTimeStr, sequenceStr, _ := parseXaddId(id)
	millieSecondsTime, _ := strconv.Atoi(millieSecondsTimeStr)
	data, exists := s.Map[streamKey]
	var sequence int
	if strings.Compare(sequenceStr, "*") != 0 {
		sequence, _ = strconv.Atoi(sequenceStr)
	} else {
		sequence = data.Stream.generateNextSeq(millieSecondsTime)
	}

	if !exists {
		// Define a new entry for the stream in the storage
		newData := NewDataStream()
		newData.Stream.AddItem(millieSecondsTime, sequence)
		s.Map[streamKey] = *newData
		return newData.Stream.GetTopId(), nil
	}

	if !data.Stream.isValidId(millieSecondsTime, sequence) {
		return "", errors.New(XADD_ERROR_ID_IS_EQUAL_OR_SMALL_TO_STREAM_TOP)
	}
	data.Stream.AddItem(millieSecondsTime, sequence)

	s.Map[streamKey] = data

	return data.Stream.GetTopId(), nil
}

func parseXaddId(id string) (string, string, error) {
	items := strings.Split(id, "-")
	if len(items) != 2 {
		return "", "", errors.New("")
	}
	return items[0], items[1], nil
}
