package utils

// ReverseMap inverts the order of items
func ReverseMap(dict map[string]string) map[string]string {
	reverseMap := make(map[string]string, len(dict))
	for key, value := range dict {
		reverseMap[value] = key
	}
	return reverseMap
}
