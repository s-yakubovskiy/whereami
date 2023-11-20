package whereami

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func println(message, color string) {
	fmt.Println(color + message + C_RESET)
}

func errorln(message string) {
	println(message, C_RED)
	os.Exit(1)
}

func warnln(message string) {
	println(message, C_YELLOW)
	os.Exit(0)
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
		log.Fatal("Failed to generate JSON", err)
	}
	return string(prettyData)
}
