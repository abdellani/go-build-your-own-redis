package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) Type(command *deserializer.Command) string {
	key := command.Arguments[0]
	valueType := h.Storage.Type(key)
	return h.Serializer.SimpleString(valueType)
}
