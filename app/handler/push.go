package handler

import "github.com/abdellani/go-build-your-own-redis/app/deserializer"

func (h *Handler) LPush(cmd *deserializer.Command) string {
	return h.Push(cmd, h.Storage.LPush)
}

func (h *Handler) RPush(cmd *deserializer.Command) string {
	return h.Push(cmd, h.Storage.RPush)
}

func (h *Handler) Push(cmd *deserializer.Command, pushFunc func(string, string) int) string {
	key := cmd.Arguments[0]
	items := len(cmd.Arguments)
	size := 0
	for i := 1; i < items; i++ {
		value := cmd.Arguments[i]
		size = pushFunc(key, value)
	}
	return h.Serializer.Integer(size)

}
