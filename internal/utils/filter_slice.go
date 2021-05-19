package utils

// FilterSlice filters a slice in-place by the criterion of the absence of an element.
func FilterSlice(slice []string, criterion []string) []string {
	dict := make(map[string]interface{}, len(criterion))
	for _, item := range criterion {
		dict[item] = nil
	}

	i, j := 0, len(slice)-1
	for i <= j {
		if _, found := dict[slice[i]]; found {
			if i != j {
				slice[i], slice[i+1], slice[j] = slice[i+1], slice[j], slice[i]
			}
			j--
		} else {
			i++
		}
	}

	return slice[:j+1]
}
