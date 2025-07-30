package handler

import (
	"strconv"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) LRange(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	start, _ := strconv.ParseInt(cmd.Arguments[1], 10, 32)
	end, _ := strconv.ParseInt(cmd.Arguments[2], 10, 32)
	return h.Serializer.BulkStringArray(h.Storage.Range(key, int(start), int(end)))
}
