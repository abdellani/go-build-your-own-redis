package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) XRange(command *deserializer.Command) string {
	key := command.Arguments[0]
	start := command.Arguments[1]
	end := command.Arguments[2]
	items := h.Storage.XRange(key, start, end)
	//serialization
	var serializedItems []string = []string{}
	for i := range len(items) {
		item := items[i]
		id := item.IdString()
		values := item.Values

		serializedId := h.Serializer.BulkString(id)
		serializedValue := h.Serializer.BulkStringArray(values)
		serializedItem := h.Serializer.Array([]string{serializedId, serializedValue})
		serializedItems = append(serializedItems, serializedItem)
	}
	return h.Serializer.Array(serializedItems)
}
