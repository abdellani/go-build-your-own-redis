package handler

import (
	"strconv"
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/config"
	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
	"github.com/abdellani/go-build-your-own-redis/app/serializer"
	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

type Handler struct {
	Storage        *storage.Storage
	Serializer     serializer.Serializer
	Configurations config.Config
}

func NewHandler(s *storage.Storage, c *config.Config) *Handler {
	return &Handler{Storage: s, Configurations: *c}
}
func (h *Handler) Handle(command *deserializer.Command) string {
	if strings.Compare(command.Command, "PING") == 0 {
		return h.Ping()
	} else if strings.Compare(command.Command, "ECHO") == 0 {
		return h.Echo(command)
	} else if strings.Compare(command.Command, "SET") == 0 {
		return h.Set(command)
	} else if strings.Compare(command.Command, "GET") == 0 {
		return h.Get(command)
	} else if strings.Compare(command.Command, "CONFIG") == 0 {
		return h.Config(command)
	} else {
		return ""
	}
}

func (h *Handler) Ping() string {
	return h.Serializer.SimpleString("PONG")
}

func (h *Handler) Echo(command *deserializer.Command) string {
	echo := command.Arguments[0]
	return h.Serializer.BulkString(echo)
}

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

func (h *Handler) Get(command *deserializer.Command) string {
	key := command.Arguments[0]
	value := h.Storage.Get(key)
	if len(value) == 0 {
		return h.Serializer.NullBulkString()
	}
	return h.Serializer.BulkString(value)
}
