package handler

import (
	"strconv"
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) Set(command *deserializer.Command) string {
	key := command.Arguments[0]
	value := command.Arguments[1]
	h.Storage.Set(key, value)
	if len(command.Arguments) == 4 {
		if strings.EqualFold(command.Arguments[2], "px") {
			duration, _ := strconv.Atoi(command.Arguments[3])
			h.Storage.SetExpiration(key, duration)
		}
	}
	return h.Serializer.SimpleString("OK")
}
