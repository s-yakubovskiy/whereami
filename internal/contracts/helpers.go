package contracts

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func createMapLocation(lat, long float64) string {
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
		mapVal = createMapLocation(lat, long)
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
		cyanP("  %s: ", field.Name)
		whiteP("%v\n", fieldValue.Interface())
	}
}

// // printStructFields handles printing each field of a struct value.
// func printStructFields(val reflect.Value, cyanP func(format string, a ...interface{}), whiteP func(format string, a ...interface{})) {
// 	lsTyp := val.Type()
// 	for i := 0; i < val.NumField(); i++ {
// 		field := lsTyp.Field(i)
// 		fieldValue := val.Field(i)
// 		cyanP("  %s: ", field.Name)
// 		whiteP("%v\n", fieldValue.Interface())
// 	}
// }

func capitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	rs[0] = unicode.ToUpper(rs[0])
	return s
	// return string(rs)
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
