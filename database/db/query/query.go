package query

import (
	"database/db/lib"
	table "database/db/table"
	"errors"
	"reflect"
)

type Condition struct {
	ColumnName string      // Name of the column to check condition against
	Operator   string      // Operator for comparison (e.g., "=", ">", "<", etc.)
	Value      interface{} // Value to compare against
}

// QueryRows returns rows from the table that match the specified conditions
func QueryRows(t *table.Table , conditions []Condition) ([]table.Row, error) {
	var matchedRows []table.Row

	// Iterate through each row in the table
	for _, row := range t.Metadata.Rows {
		// Check if the row matches all conditions
		matched := true
		for _, cond := range conditions {
			// Check if the column exists in the row's data
			value, ok := row.Data[cond.ColumnName]
			if !ok {
				matched = false
				break
			}

			// Perform type assertion to handle comparison based on the value's type
			switch v := value.(type) {
			case int:
				if !lib.CompareInt(v, cond.Operator, cond.Value) {
					matched = false
					break
				}
			case float64:
				if !lib.CompareFloat64(v, cond.Operator, cond.Value) {
					matched = false
					break
				}
			case string:
				if !lib.CompareString(v, cond.Operator, cond.Value) {
					matched = false
					break
				}
			// Add more cases for other types as needed
			default:
				return nil, errors.New("unsupported data type for comparison: " + reflect.TypeOf(value).String())
			}
		}
		// If the row matches all conditions, add it to the matchedRows slice
		if matched {
			matchedRows = append(matchedRows, row)
		}
	}

	return matchedRows, nil
}

