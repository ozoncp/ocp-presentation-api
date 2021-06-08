// Package util implement a simple internal library for Ozon Code Platform Presentation API.
package util

import "github.com/ozoncp/ocp-presentation-api/internal/model"

// FilterSlice filters a slice by the criterion of the absence of keys.
func FilterSlice(arr []model.Presentation, keys []model.Presentation) []model.Presentation {
	if arr == nil || keys == nil {
		return arr
	}

	dict := make(map[uint64]struct{}, len(keys))
	for _, key := range keys {
		dict[key.ID] = struct{}{}
	}

	var result []model.Presentation
	for _, item := range arr {
		if _, found := dict[item.ID]; !found {
			result = append(result, item)
		}
	}

	return result
}
