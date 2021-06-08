// Package util implement a simple internal library for Ozon Code Platform Presentation API.
package util

// ReverseMap inverts the order of items
func ReverseMap(dict map[string]string) map[string]string {
	if dict == nil {
		return nil
	}

	reverseMap := make(map[string]string, len(dict))
	for key, value := range dict {
		reverseMap[value] = key
	}
	return reverseMap
}