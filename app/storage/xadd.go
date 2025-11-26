package storage

import (
	"errors"
	"strconv"
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/storage/stream"
)

const (
	XADD_ERROR_ID_MUST_BE_GREATER_THAN_0_0        = "ERR The ID specified in XADD must be greater than 0-0"
	XADD_ERROR_ID_IS_EQUAL_OR_SMALL_TO_STREAM_TOP = "ERR The ID specified in XADD is equal or smaller than the target stream top item"
)

func (s *Storage) XAdd(streamKey, id string, values []string) (string, error) {
	s.m.Lock()
	defer s.m.Unlock()
	//TODO: isolate the id parsing logic to the stream structure
	if strings.Compare("0-0", id) == 0 {
		return "", errors.New(XADD_ERROR_ID_MUST_BE_GREATER_THAN_0_0)
	}
	data, exists := s.Map[streamKey]
	var millieSecondsTime int
	var sequence int

	if strings.Compare("*", id) == 0 {
		millieSecondsTime, sequence = data.Stream.GenerateFullId()
	} else {
		millieSecondsTimeStr, sequenceStr, _ := stream.SplitId(id)
		millieSecondsTime, _ = strconv.Atoi(millieSecondsTimeStr)
		if strings.Compare(sequenceStr, "*") != 0 {
			sequence, _ = strconv.Atoi(sequenceStr)
		} else {
			sequence = data.Stream.GenerateNextSeq(millieSecondsTime)
		}
	}

	if !exists {
		// Define a new entry for the stream in the storage
		newData := NewDataStream()
		newData.Stream.AddItem(millieSecondsTime, sequence, values)
		s.Map[streamKey] = *newData
		return newData.Stream.GetTopId(), nil
	}

	if !data.Stream.IsValidId(millieSecondsTime, sequence) {
		return "", errors.New(XADD_ERROR_ID_IS_EQUAL_OR_SMALL_TO_STREAM_TOP)
	}
	data.Stream.AddItem(millieSecondsTime, sequence, values)

	s.Map[streamKey] = data

	return data.Stream.GetTopId(), nil
}
