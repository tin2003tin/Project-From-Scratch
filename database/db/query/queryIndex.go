package query

import (
	table "database/db/table"
	"errors"
	"fmt"
)

func QueryRowByIndex(index *table.Index, target interface{}) (*table.Row, error) {
	if index == nil {
		return nil, errors.New("index is nil")
	}
	key := fmt.Sprintf("%v", target) 
	indexKey := make(map[int]interface{})
	for columnIndex, column := range index.Columns {
		if column.Name == index.Name { 
			indexKey[columnIndex] = key 
			break 
		}
	}

	if len(indexKey) > 0 {
		row, ok := index.Rows[fmt.Sprintf("%v", indexKey)]
		if !ok {
			return nil, fmt.Errorf("row not found for key: %v", target)
		}
		return row, nil
	}

	return nil, errors.New("key column not found in the index")
}
