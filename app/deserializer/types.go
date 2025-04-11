package deserializer

import "strings"

func IsArray(symbol string) bool {
	return strings.Compare(symbol, "*") == 0
}

func IsString(symbol string) bool {
	return strings.Compare(symbol, "$") == 0
}
