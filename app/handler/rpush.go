package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) RPush(cmd *deserializer.Command) string {
	return h.Serializer.Integer(1)
}
