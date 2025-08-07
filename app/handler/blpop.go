package handler

import (
	"strconv"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) BLPop(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	waitingTimeString := cmd.Arguments[1]
	waitingTimeFloat64, _ := strconv.ParseFloat(waitingTimeString, 64)
	waitingMillieseconds := int(waitingTimeFloat64 * 1000)
	result, found := h.Storage.BLPOP(key, waitingMillieseconds)

	if found {
		return h.Serializer.BulkStringArray([]string{key, result})
	} else {
		return h.Serializer.NullBulkString()
	}
}
