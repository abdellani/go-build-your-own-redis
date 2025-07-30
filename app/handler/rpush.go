package handler

import (
	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) RPush(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	value := cmd.Arguments[1]
	size := h.Storage.Push(key, value)
	return h.Serializer.Integer(size)
}
