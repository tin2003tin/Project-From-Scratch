package lib

import "fmt"

func MappingIndexKey(columnName string, value interface{}) string {
	indexKey := make(map[string]interface{})
	indexKey[columnName] = value
	return fmt.Sprintf("%v", indexKey)
}
