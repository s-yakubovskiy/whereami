package locator

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
	"github.com/s-yakubovskiy/whereami/internal/entity"
)

func CreateMapLocation(lat, long float64) string {
	mapPattern := "https://www.google.com/maps?q=%f,%f"

	return fmt.Sprintf(mapPattern, lat, long)
}

func handleDateField(fieldVal reflect.Value, cyanP func(format string, a ...interface{}), whiteP func(format string, a ...interface{})) {
	dateVal, isString := fieldVal.Interface().(string)
	if !isString || dateVal == "" {
		// If the date is not a string or is empty, use the current time
		dateVal = time.Now().Format("Monday, January 2, 2006, 15:04")
	} else {
		// Try to parse the existing date value
		parsedTime, err := time.Parse(time.RFC3339, dateVal)
		if err != nil {
			// If parsing fails, use the original string value
			cyanP("  date: ")
			whiteP("%v\n", dateVal)
			return
		}
		dateVal = parsedTime.Format("Monday, January 2, 2006, 15:04")
	}
	cyanP("  date: ")
	whiteP("%s\n", dateVal)
}

func handleMapField(fieldVal reflect.Value, cyanP func(format string, a ...interface{}), whiteP func(format string, a ...interface{}), lat, long float64) {
	mapVal, isString := fieldVal.Interface().(string)
	if !isString || mapVal == "" {
		mapVal = CreateMapLocation(lat, long)
	}
	cyanP("  map: ")
	whiteP("%s\n", mapVal)
}

// isZeroValue checks if the given reflect.Value is considered a zero value for its type.
func isZeroValue(v reflect.Value) bool {
	// Handle cases where the value is a pointer or interface
	if v.Kind() == reflect.Interface {
		return v.IsNil()
	}
	// Use reflect's Zero function to compare with a zero value of the same type
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// printStructFields handles printing each field of a struct value if they are not zero-valued.
func printStructFields(val reflect.Value, cyanP func(format string, a ...interface{}), whiteP func(format string, a ...interface{})) {
	lsTyp := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := lsTyp.Field(i)
		fieldValue := val.Field(i)
		// Check if the field should be skipped based on its zero value
		// if isZeroValue(fieldValue) {
		// 	continue // Skip this iteration if the value is zero
		// }
		cyanP("  %s: ", capitalizeFirst(field.Name))
		whiteP("%v\n", fieldValue.Interface())
	}
}

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	if len(rs) <= 3 {
		return strings.ToUpper(string(rs))
	}
	// return s
	return string(rs)
}

// findField locates a field by name within a struct, accounting for case insensitivity.
func findField(val reflect.Value, name string) (reflect.Value, bool) {
	fieldVal := val.FieldByNameFunc(func(n string) bool {
		return strings.EqualFold(n, name)
	})
	return fieldVal, fieldVal.IsValid()
}

// isEmpty checks if the given value is considered "empty" for skipping output.
func isEmpty(value interface{}) bool {
	// Check for nil directly
	if value == "scores" {
		return false
	}
	if value == nil {
		return true
	}
	// Convert to reflect.Value and check for zero values
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	}
	return false
}

// Output reworked to support ordered categories with conditional field output and correct struct handling
func (uc *UseCase) Output(location *entity.Location, categories map[string][]string, orderedCategories []string) {
	val := reflect.ValueOf(*location)

	// Prepare color outputs
	cyan := color.New(color.FgCyan).Add(color.Bold)
	white := color.New(color.FgWhite)                     // For field values
	magenta := color.New(color.FgMagenta).Add(color.Bold) // For categories
	cyanP := cyan.PrintfFunc()
	whiteP := white.PrintfFunc()
	magentaP := magenta.PrintfFunc() // Print function for category names in magenta

	// Iterate through each category in the ordered list
	for _, category := range orderedCategories {
		fields, exists := categories[category]
		if !exists {
			continue // Skip categories not defined in the map
		}
		magentaP("\n%s\n", capitalizeFirst(category)) // Print category name in magenta
		for _, field := range fields {
			fieldVal, found := findField(val, field)
			if !found {
				continue // Skip fields not found
			}
			if field == "date" {
				handleDateField(fieldVal, cyanP, whiteP)
				continue
			}

			if field == "map" {
				handleMapField(fieldVal, cyanP, whiteP, location.Latitude, location.Longitude)
			}
			if field == "gps" {
				printStructFields(fieldVal, cyanP, whiteP)
				continue
			}

			fieldValue := fieldVal.Interface()
			if reflect.TypeOf(fieldValue).Kind() == reflect.Struct && field == "scores" {
				// Handle struct types specifically for 'scores'
				if field != "scores" {
					cyanP("  %s:\n", capitalizeFirst(field))
				}
				printStructFields(fieldVal, cyanP, whiteP)
			} else if !isEmpty(fieldValue) {
				// Skip empty fields except for 'date', which is handled separately
				cyanP("  %s: ", capitalizeFirst(field))
				whiteP("%v\n", fieldValue)
			}
		}
	}
}
