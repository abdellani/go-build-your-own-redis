package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

type Handler struct {
	Storage *storage.Storage
}

func NewHandler(s *storage.Storage) *Handler {
	return &Handler{Storage: s}
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
	} else {
		return ""
	}
}

func (h *Handler) Ping() string {
	return "+PONG\r\n"
}

func (h *Handler) Echo(command *deserializer.Command) string {
	echo := command.Arguments[0]
	responseTemplate := "$%d\r\n%s\r\n"
	return fmt.Sprintf(responseTemplate, len(echo), echo)
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
	return "+OK\r\n"
}

func (h *Handler) Get(command *deserializer.Command) string {
	key := command.Arguments[0]
	value := h.Storage.Get(key)
	if len(value) == 0 {
		return "$-1\r\n"
	}
	template := "$%d\r\n%s\r\n"

	return fmt.Sprintf(template, len(value), value)
}
