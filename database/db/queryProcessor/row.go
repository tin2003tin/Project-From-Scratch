package queryProcessor

import (
	"database/db/lib"
	"database/db/structure"
	"errors"
	"fmt"
	"time"
)

func (q *QueryManager) AddRow(columnValues []interface{}) error {
	// Check if the target table is nil

	if q.Table == nil {
		return errors.New("target table is nil")
	}
	createdAt := time.Now()
	newRow := structure.Row{
		Data:      columnValues,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
	// Validate primary key and non-nullable columns
	for _, column := range q.Table.Columns {
		if column.PrimaryKey && columnValues[column.Index] == nil {
			return errors.New("cannot add row, primary key column not provided: " + column.Name)
		}
		if newRow.Data[column.Index] != nil {
			newValue, err := lib.ConvertValue(newRow.Data[column.Index], column.DataType, column.Length)
			if err != nil {
				return fmt.Errorf("cannot add row, %v %v", column.Name, err)
			}
			newRow.Data[column.Index] = newValue
		}
		if column.PrimaryKey || column.Unique {
			temp := make(map[string]string, 0)
			temp[column.Name] = fmt.Sprintf("%v", columnValues[column.Index])
			_, found := q.Table.IndexTable.Rows[fmt.Sprintf("%v", temp)]
			if found {
				return errors.New("cannot add row, primary key or unique already exists: " + column.Name + fmt.Sprintf(" %v", columnValues[column.Index]))
			}
		}
		if !column.Nullable && columnValues[column.Index] == nil {
			if column.Default == nil {
				return errors.New("cannot add row, on-nullable column not provided: " + column.Name)
			} else {
				newRow.Data[column.Index] = column.Default
			}
		}

		// if column.ForeignKey {
		// 	if !q.dataInForeignKeyExisted(&newRow, column.Name) {
		// 		return fmt.Errorf("cannot add row, %v not found in foreign key column '%s'", newRow.Data[column.Index], column.Name)
		// 	}
		// }
	}
	q.Table.Rows = append(q.Table.Rows, newRow)

	// Update the index with the new values
	q.AddRowToIndex(q.Table.IndexTable, newRow)
	return nil
}
