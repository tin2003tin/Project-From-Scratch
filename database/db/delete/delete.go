package delete

import (
	"database/db/lib"
	"database/db/table"
	"fmt"
)

type Condition struct {
	ColumnName string      // Name of the column to check condition against
	Operator   string      // Operator for comparison (e.g., "=", ">", "<", etc.)
	Value      interface{} // Value to compare against
}

func DeleteRow(t *table.Table, conditions []Condition) error {
	if t == nil {
		return fmt.Errorf("table is nil")
	}
	if len(conditions) == 0 {
		return fmt.Errorf("no conditions provided")
	}

	// Iterate over each row in the table
	for i := len(t.Metadata.Rows) - 1; i >= 0; i-- {
		row := t.Metadata.Rows[i]
		if checkAllConditions(row, conditions) {
			// Remove the row from the table's rows slice
			t.Metadata.Rows = append(t.Metadata.Rows[:i], t.Metadata.Rows[i+1:]...)
			// Delete the row from the index using the specified column name
			if err := deleteFromIndex(t, row); err != nil {
				return err
			}
		}
	}

	return nil
}


func checkAllConditions(row table.Row, conditions []Condition) bool {
	for _, condition := range conditions {
		if !checkSingleCondition(row, condition) {
			return false
		}
	}
	return true
}


func checkSingleCondition(row table.Row, condition Condition) bool {
	value, ok := row.Data[condition.ColumnName]
	if !ok {
		return false
	}
	switch v := value.(type) {
	case int:
		return lib.CompareInt(v, condition.Operator, condition.Value)
	case float64:
		return lib.CompareFloat64(v, condition.Operator, condition.Value)
	case string:
		return lib.CompareString(v, condition.Operator, condition.Value)
	default:
		return false
	}
}

func deleteFromIndex(t *table.Table, row table.Row) error {
	for _,index := range t.IndexTable {
		key := fmt.Sprintf("%v", row.Data[index.Name])
		indexKey := make(map[string]interface{})
		for _, column := range index.Columns {
			if column.Name == index.Name { 
				indexKey[column.Name] = key 
				break 
			}
		}
		fmt.Println(indexKey)
		delete_row := fmt.Sprintf("%v", indexKey)
		delete(index.Rows, delete_row)
	}
	return nil;
}