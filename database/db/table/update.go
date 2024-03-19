package table

import (
	"fmt"
)

type Set struct {
	ColumnName string
	Value      interface{}
}

func (t *Table) Update(conditions []Condition, sets []Set) error {
	if t == nil {
		return fmt.Errorf("table is nil")
	}

	// Iterate over each row in the table
	for i := len(t.Metadata.Rows) - 1; i >= 0; i-- {
		row := t.Metadata.Rows[i]
		if checkAllConditions(row, conditions) {
			for _, set := range sets {
				if err := updateFromIndex(t, &row, set); err != nil {
					return err
				}
				if err := updateRow(t, &row, set); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func updateRow(t *Table, row *Row, set Set) error {
	for _, col := range t.Metadata.Columns {
		if set.ColumnName == col.Name {
			row.Data[set.ColumnName] = set.Value
			break
		}
	}
	return nil
}

func updateFromIndex(t *Table, row *Row, set Set) error {
	indexColumns := t.IndexTable.Columns
	indexRows := t.IndexTable.Rows
	for _, column := range indexColumns {
		if set.ColumnName == column.Name {
			newTemp := make(map[string]interface{})
			oldTemp := make(map[string]interface{})
			key := fmt.Sprintf("%v", row.Data[column.Name])
			oldTemp[column.Name] = key
			newTemp[column.Name] = set.Value
			oldkey := fmt.Sprintf("%v", oldTemp)
			newKey := fmt.Sprintf("%v", newTemp)

			indexRows[newKey] = indexRows[oldkey]
			delete(indexRows, oldkey)
		}
	}
	return nil
}
