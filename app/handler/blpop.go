package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) BLPop(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	return h.Serializer.BulkStringArray([]string{key, h.Storage.BLPOP(key, 0)})
}
