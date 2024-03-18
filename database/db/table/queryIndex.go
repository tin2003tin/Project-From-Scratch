package table

import (
	"errors"
	"fmt"
)

func (t *Table) QueryRowByIndex(columnName string, target interface{}) (*Row, error) {
	if columnName == "" {
		return nil, errors.New("index is empty")
	}
	key := fmt.Sprintf("%v", target)
	indexKey := make(map[string]interface{})
	for _, column := range t.IndexTable.Columns {
		if column.Name == columnName {
			indexKey[column.Name] = key
			break
		}
	}
	fmt.Println(indexKey)
	if len(indexKey) > 0 {
		row, ok := t.IndexTable.Rows[fmt.Sprintf("%v", indexKey)]
		if !ok {
			return nil, fmt.Errorf("row not found for key: %v", target)
		}
		return row, nil
	}

	return nil, errors.New("key column not found in the index")
}
