package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) Keys(command *deserializer.Command) string {
	return h.Serializer.BulkStringArray(h.Storage.Keys())
}
