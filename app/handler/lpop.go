package handler

import (
	"strconv"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) LPop(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	num := int64(1)
	if len(cmd.Arguments) > 1 {
		num, _ = strconv.ParseInt(cmd.Arguments[1], 10, 32)
	}
	res := h.Storage.LPop(key, int(num))
	return h.Serializer.BulkStringArrayOrSingle(res)
}
