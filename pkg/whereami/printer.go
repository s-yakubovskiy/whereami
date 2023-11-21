package whereami

import (
	"encoding/json"
	"fmt"
	"log"
)

func println(message, color string) {
	fmt.Println(color + message + C_RESET)
}

func errorln(message string) {
	println(message, C_RED)
}

func warnln(message string) {
	println(message, C_YELLOW)
}

func successln(message string) {
	println(message, C_GREEN)
}

func infoln(message string) {
	println(message, C_BLUE)
}

func PrettyPrint(v interface{}) string {
	prettyData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Failed to generate JSON", err)
	}
	return string(prettyData)
}
