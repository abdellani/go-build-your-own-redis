package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) XAdd(command *deserializer.Command) string {
	key := command.Arguments[0]
	id := command.Arguments[1]
	return h.Serializer.BulkString(h.Storage.XAdd(key, id))
}
