package common

import (
	"encoding/json"
	"fmt"
	"log"
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

func println(message, color string) {
	fmt.Println(color + message + C_RESET)
}

func Errorln(message string) {
	println(message, C_RED)
}

func Warnln(message string) {
	println(message, C_YELLOW)
}

func Successln(message string) {
	println(message, C_GREEN)
}

func Infoln(message string) {
	println(message, C_BLUE)
}

func PrettyPrint(v interface{}) string {
	prettyData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Failed to generate JSON", err)
	}
	return string(prettyData)
}
