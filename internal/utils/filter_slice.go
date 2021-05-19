package utils

// FilterSlice filters a slice by the criterion of the absence of an element.
func FilterSlice(slice []string, criterion []string) []string {
	dict := make(map[string]interface{}, len(criterion))
	for _, item := range criterion {
		dict[item] = nil
	}

	var result []string
	for _, item := range slice {
		if _, found := dict[item]; !found {
			result = append(result, item)
		}
	}

	return result
}
