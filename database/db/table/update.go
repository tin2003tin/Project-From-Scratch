package table

import (
	"database/db/lib"
	"errors"
	"fmt"
)

type Set struct {
	ColumnName string
	Value      interface{}
}

func (t *Table) UpdateRow(conditions []Condition, sets []Set) error {
	if t == nil {
		return fmt.Errorf("table is nil")
	}

	// Iterate over each row in the table

	for i := len(t.Metadata.Rows) - 1; i >= 0; i-- {
		row := t.Metadata.Rows[i]
		if checkAllConditions(row, conditions) {
			for _, set := range sets {
				for _, column := range t.Metadata.Columns {
					if column.PrimaryKey || column.Unique {
						temp := make(map[string]string, 0)
						temp[set.ColumnName] = fmt.Sprintf("%v", set.Value)
						_, found := t.IndexTable.Rows[fmt.Sprintf("%v", temp)]
						if found {
							return errors.New("cannot update row, primary key or unique already exists: " + set.ColumnName + fmt.Sprintf(" %v", set.Value))
						}
					}
					if column.ForeignKey {
						if !t.dataInForeignKeyExisted(&Row{Data: map[string]interface{}{set.ColumnName: set.Value}}, column.Name) {
							return fmt.Errorf("cannot add row, %v not found in foreign key column '%s'", set.Value, column.Name)
						}
					}
					if column.Name == set.ColumnName {
						newValue, err := lib.ConvertValue(set.Value, column.DataType, column.Length)
						if err != nil {
							return fmt.Errorf("cannot update, %v %v", column.Name, err)
						}
						set.Value = newValue
					}
				}

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
