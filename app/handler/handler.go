package handler

import (
	"log"
	"os"
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
	switch command.Command {
	case "PING":
		return h.Ping()
	case "ECHO":
		return h.Echo(command)
	case "SET":
		return h.Set(command)
	case "GET":
		return h.Get(command)
	case "CONFIG":
		return h.Config(command)
	default:
		log.Fatal("command not recognized")
		os.Exit(-1)
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
