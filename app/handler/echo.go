package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) Echo(command *deserializer.Command) string {
	echo := command.Arguments[0]
	return h.Serializer.BulkString(echo)
}
