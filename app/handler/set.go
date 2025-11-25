package handler

import (
	"strconv"
	"strings"
	"time"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

func (h *Handler) Set(command *deserializer.Command) string {
	key := command.Arguments[0]
	value := storage.Data{
		Value: command.Arguments[1],
	}

	if len(command.Arguments) == 4 && strings.EqualFold(command.Arguments[2], "px") {
		duration, _ := strconv.Atoi(command.Arguments[3])
		value.ExpirationTime = convertDurationToTime(duration)
	}

	h.Storage.Set(key, value)
	return h.Serializer.SimpleString("OK")
}

func convertDurationToTime(millieseconds int) time.Time {
	return time.Now().Add(time.Millisecond * time.Duration(millieseconds))
}
