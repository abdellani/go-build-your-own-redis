package serializer

import (
	"fmt"

	"github.com/abdellani/go-build-your-own-redis/app/storage/stream"
)

type Serializer struct {
}

func (s *Serializer) Integer(i int) string {
	return fmt.Sprintf(":%d\r\n", i)
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

func (s *Serializer) NullArray() string {
	return "*-1\r\n"
}

func (s *Serializer) BulkStringArray(texts []string) string {
	size := len(texts)
	result := fmt.Sprintf("*%d\r\n", size)
	for _, text := range texts {
		result += s.BulkString(text)
	}
	return result
}

func (s *Serializer) Array(texts []string) string {
	size := len(texts)
	result := fmt.Sprintf("*%d\r\n", size)
	for _, text := range texts {
		result += text
	}
	return result
}
func (s *Serializer) BulkStringArrayOrSingle(texts []string) string {
	if len(texts) == 1 {
		return s.BulkString(texts[0])
	}
	return s.BulkStringArray(texts)
}

func (s *Serializer) SimpleError(text string) string {
	return fmt.Sprintf("-%s\r\n", text)
}

func (s *Serializer) StreamItems(items []stream.StreamItem) string {
	var serializedItems []string = []string{}
	for i := range len(items) {
		item := items[i]
		id := item.IdString()
		values := item.Values

		serializedId := s.BulkString(id)
		serializedValue := s.BulkStringArray(values)
		serializedItem := s.Array([]string{serializedId, serializedValue})
		serializedItems = append(serializedItems, serializedItem)
	}
	return s.Array(serializedItems)
}

func (s *Serializer) XReadResult(key string, streamItems []stream.StreamItem) string {

	return s.Array(
		[]string{
			s.Array([]string{
				s.BulkString(key),
				s.StreamItems(streamItems),
			}),
		},
	)
}
