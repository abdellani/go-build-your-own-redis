package handler

import (
	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) RPush(cmd *deserializer.Command) string {
	key := cmd.Arguments[0]
	items := len(cmd.Arguments)
	size := 0
	for i := 1; i < items; i++ {
		value := cmd.Arguments[i]
		size = h.Storage.Push(key, value)
	}
	return h.Serializer.Integer(size)
}
