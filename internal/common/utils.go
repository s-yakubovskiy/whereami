package common

import (
	"fmt"
	"reflect"
	"strings"
)

// CountryCodeToEmoji converts a two-letter country code to its corresponding flag emoji.
func CountryCodeToEmoji(code string) string {
	const offset = 127397 // Offset between uppercase ASCII and regional indicator symbols
	var emoji strings.Builder

	for _, r := range strings.ToUpper(code) {
		if r < 'A' || r > 'Z' {
			return "" // Invalid code
		}
		emoji.WriteRune(rune(r) + offset)
	}

	return emoji.String()
}

func printStruct(s interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
