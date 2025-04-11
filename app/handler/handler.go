package handler

import (
	"fmt"
	"strings"

	"github.com/abdellani/go-build-your-own-redis/app/deserializer"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) Handle(command *deserializer.Command) string {
	if strings.Compare(command.Command, "PING") == 0 {
		return HandlePing()
	} else if strings.Compare(command.Command, "ECHO") == 0 {
		return handleEcho(command)
	}
	return ""
}

func HandlePing() string {
	return "+PONG\r\n"
}

func handleEcho(command *deserializer.Command) string {
	echo := command.Arguments[0]
	responseTemplate := "$%d\r\n%s\r\n"
	return fmt.Sprintf(responseTemplate, len(echo), echo)
}
