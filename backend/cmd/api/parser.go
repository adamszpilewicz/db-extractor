package main

import "fmt"

// Utility function to check if an interface holds a [16]uint8 type
func isUUID(value interface{}) bool {
	_, ok := value.([16]uint8)
	return ok
}

// Utility function to format [16]uint8 as UUID string
func formatUUID(value [16]uint8) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", value[0:4], value[4:6], value[6:8], value[8:10], value[10:])
}
