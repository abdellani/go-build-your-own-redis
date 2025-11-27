package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) XRead(command *deserializer.Command) string {
	key := command.Arguments[1]
	id := command.Arguments[2]
	result := h.Storage.XRead(key, id)
	return h.Serializer.XReadResult(key, result)
}
