package utils

// SplitSlice splits a slice into uniform chunks.
func SplitSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string

	if chunkSize <= 0 {
		return chunks
	}

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
