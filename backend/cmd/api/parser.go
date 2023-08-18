package main

import (
	"fmt"
	"github.com/jackc/pgx/v4"
)

// Utility function to check if an interface holds a [16]uint8 type
func isUUID(value interface{}) bool {
	_, ok := value.([16]uint8)
	return ok
}

// Utility function to format [16]uint8 as UUID string
func formatUUID(value [16]uint8) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", value[0:4], value[4:6], value[6:8], value[8:10], value[10:])
}

// ProcessRows takes a rows object and processes it into a slice of TableRow
func ProcessRows(rows pgx.Rows) ([]TableRow, error) {
	var results []TableRow

	// Iterate through the result rows
	for rows.Next() {
		row := make(TableRow)

		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		fieldDescriptions := rows.FieldDescriptions()

		for i, value := range values {
			columnName := string(fieldDescriptions[i].Name)
			// Check and format UUID
			if isUUID(value) {
				value = formatUUID(value.([16]uint8))
			}
			row[columnName] = value
		}

		results = append(results, row)
	}

	return results, nil
}
