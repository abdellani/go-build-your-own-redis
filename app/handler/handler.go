package handler

import (
	"log"
	"os"

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
	case "KEYS":
		return h.Keys(command)
	case "RPUSH":
		return h.RPush(command)
	case "LLEN":
		return h.LLen(command)
	case "LPUSH":
		return h.LPush(command)
	case "LRANGE":
		return h.LRange(command)
	case "LPOP":
		return h.LPop(command)
	case "BLPOP":
		return h.BLPop(command)
	default:
		log.Fatal("command not recognized")
		os.Exit(-1)
		return ""

	}
}
