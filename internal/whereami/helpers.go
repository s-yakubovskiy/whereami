package whereami

import "strings"

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
