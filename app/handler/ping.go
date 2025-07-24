package handler

func (h *Handler) Ping() string {
	return h.Serializer.SimpleString("PONG")
}
