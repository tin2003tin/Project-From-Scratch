package table

import (
	"database/db/lib"
	"fmt"
)

func (t *Table) DeleteRow(conditions []Condition) error {
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

func checkAllConditions(row Row, conditions []Condition) bool {
	for _, condition := range conditions {
		if !checkSingleCondition(row, condition) {
			return false
		}
	}
	return true
}

func checkSingleCondition(row Row, condition Condition) bool {
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

func deleteFromIndex(t *Table, row Row) error {
	key := fmt.Sprintf("%v", row.Data[t.IndexTable.Name])
	indexKey := make(map[string]interface{})
	for _, column := range t.IndexTable.Columns {
		if column.Name == t.IndexTable.Name {
			indexKey[column.Name] = key
			break
		}
	}
	delete_row := fmt.Sprintf("%v", indexKey)
	delete(t.IndexTable.Rows, delete_row)
	return nil
}
