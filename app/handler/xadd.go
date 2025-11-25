package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) XAdd(command *deserializer.Command) string {
	key := command.Arguments[0]
	id := command.Arguments[1]
	response, err := h.Storage.XAdd(key, id)
	if err != nil {
		return h.Serializer.SimpleError(err.Error())
	}
	return h.Serializer.BulkString(response)
}
