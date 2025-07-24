package serializer

import "fmt"

type Serializer struct {
}

func (s *Serializer) SimpleString(text string) string {
	return fmt.Sprintf("+%s\r\n", text)
}

func (s *Serializer) BulkString(text string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(text), text)
}

func (s *Serializer) NullBulkString() string {
	return "$-1\r\n"
}

func (s *Serializer) BulkStringArray(texts []string) string {
	size := len(texts)
	result := fmt.Sprintf("*%d\r\n", size)
	for _, text := range texts {
		result += s.BulkString(text)
	}
	return result
}
