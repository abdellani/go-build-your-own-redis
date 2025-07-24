package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) Get(command *deserializer.Command) string {
	key := command.Arguments[0]
	value := h.Storage.Get(key)
	if len(value) == 0 {
		return h.Serializer.NullBulkString()
	}
	return h.Serializer.BulkString(value)
}
