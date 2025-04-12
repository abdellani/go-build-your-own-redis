package handler

import (
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

func (h *Handler) Config(command *deserializer.Command) string {
	subCommand := command.Arguments[0]
	if strings.Compare(subCommand, "GET") == 0 {
		return h.ConfigGet(command)
	}
	return ""
}

func (h *Handler) ConfigGet(command *deserializer.Command) string {
	param := command.Arguments[1]
	if strings.Compare(param, "dir") == 0 {
		dir := h.Configurations.RDB.Dir
		items := []string{
			h.Serializer.BulkString("dir"),
			h.Serializer.BulkString(dir),
		}
		return h.Serializer.BulkStringArray(items)
	} else if strings.Compare(param, "dbfilename") == 0 {
		dbFileName := h.Configurations.RDB.FileName
		items := []string{
			h.Serializer.BulkString("dbfilename"),
			h.Serializer.BulkString(dbFileName),
		}
		return h.Serializer.BulkStringArray(items)
	}
	return ""
}
