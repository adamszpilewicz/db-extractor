package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
)

func ConvertRowsToCSV(rows []TableRow) ([]byte, error) {
	// Use a buffer to write our CSV to
	var b bytes.Buffer
	writer := csv.NewWriter(&b)

	// If rows are empty, return an empty CSV
	if len(rows) == 0 {
		return b.Bytes(), nil
	}

	// Extract column headers for the CSV.
	// Assuming all rows have the same columns
	headers := make([]string, 0, len(rows[0]))
	for column := range rows[0] {
		headers = append(headers, column)
	}

	// Write the headers to the CSV
	if err := writer.Write(headers); err != nil {
		return nil, err
	}

	// Write each row to the CSV
	for _, row := range rows {
		record := make([]string, 0, len(row))
		for _, header := range headers {
			value := fmt.Sprint(row[header]) // Convert interface{} to string
			record = append(record, value)
		}
		if err := writer.Write(record); err != nil {
			return nil, err
		}
	}

	// Flush any buffered data to the writer
	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
