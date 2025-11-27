package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) XRange(command *deserializer.Command) string {
	key := command.Arguments[0]
	start := command.Arguments[1]
	end := command.Arguments[2]
	items := h.Storage.XRange(key, start, end)
	return h.Serializer.StreamItems(items)
}
