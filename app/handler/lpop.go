package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) LPop(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	return h.Serializer.BulkString(h.Storage.LPop(key))
}
