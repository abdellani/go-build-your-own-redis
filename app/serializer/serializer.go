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
	result := []string{}
	for _, text := range texts {
		result = append(result, s.BulkString(text))
	}
	return s.Array(result)
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

		serializedItem := s.Array([]string{
			s.BulkString(id),
			s.BulkStringArray(values),
		})
		serializedItems = append(serializedItems, serializedItem)
	}
	return s.Array(serializedItems)
}

func (s *Serializer) XReadResult(keys []string, streamItems [][]stream.StreamItem) string {
	result := []string{}
	for i, key := range keys {
		streamItem := streamItems[i]
		result = append(result, s.Array([]string{
			s.BulkString(key),
			s.StreamItems(streamItem),
		}))
	}
	return s.Array(result)
}
