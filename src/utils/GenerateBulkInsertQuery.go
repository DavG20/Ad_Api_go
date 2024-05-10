package utils

import (
	"fmt"
	"strings"
)

func CreateManyRawQuery(values [][]string, tableName string, columns []string) string {

	// Start the INSERT statement
	query := "INSERT INTO " + tableName + " (" + strings.Join(columns, ", ") + ") VALUES "

	// Slice to hold value strings
	valueStrings := make([]string, 0, len(values))

	// Iterate over employees to create row strings
	for _, value := range values {
		row := fmt.Sprintf("(%s)", strings.Join(value, ", "))
		valueStrings = append(valueStrings, row)
	}

	// Concatenate the value strings
	query += strings.Join(valueStrings, ", ")

	return query
}
