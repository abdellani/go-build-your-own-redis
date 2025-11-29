package handler

import (
	"math"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) XRead(command *deserializer.Command) string {
	keys := []string{}
	ids := []string{}
	middle := int(math.Floor(float64(len(command.Arguments)) / 2))
	for i := range middle {
		keyIdx := i + 1
		idIdx := i + middle
		keys = append(keys, command.Arguments[keyIdx])
		ids = append(ids, command.Arguments[idIdx])
	}
	results := h.Storage.XRead(keys, ids)
	return h.Serializer.XReadResult(keys, results)
}
