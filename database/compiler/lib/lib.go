package lib

import (
	"strings"
)

func AddUnique(element string, array *[]string) bool {
	if !IsElement(element, *array) {
		*array = append(*array, element)
		return true
	}
	return false
}

func AddUniqueMap(element string, dictionary map[string][]string, key string) bool {
	if !IsElement(element, dictionary[key]) {
		dictionary[key] = append(dictionary[key], element)
		return true
	}
	return false
}
func IsElement(element string, array []string) bool {
	for _, val := range array {
		if element == val {
			return true
		}
	}
	return false
}

func TrimElements(elements []string) []string {
	trimmed := make([]string, len(elements))
	for i, e := range elements {
		trimmed[i] = strings.TrimSpace(e)
	}
	return trimmed
}
