package queryProcessor

import (
	"database/db/structure"
	"fmt"
)

func (q *QueryManager) DeleteRow(conditions []structure.Condition) ([][]*structure.Row, error) {
	var delete_rows [][]*structure.Row = [][]*structure.Row{}
	if q.Table == nil {
		return nil, fmt.Errorf("table is nil")
	}
	var mappingColumn map[string]int = make(map[string]int)
	for i, column := range q.CurrentColumns {
		mappingColumn[column.Name] = i

	}
	for i := len(q.Table.Rows) - 1; i >= 0; i-- {
		row := q.Table.Rows[i]
		for _, cond := range conditions {
			if passed, _ := checkCondition(row.Data[mappingColumn[cond.ColumnName]], &cond); passed {
				delete_rows = append(delete_rows, []*structure.Row{&q.Table.Rows[i]})
				q.Table.Rows = append(q.Table.Rows[:i], q.Table.Rows[i+1:]...)
				if err := deleteFromIndex(q.Table, row); err != nil {
					return nil, err
				}
			}
		}
	}

	return delete_rows, nil
}

func deleteFromIndex(t *structure.Table, row structure.Row) error {
	for _, Column := range t.IndexTable.Columns {
		indexKey := make(map[string]interface{})
		key := fmt.Sprintf("%v", row.Data[Column.Index])
		indexKey[Column.Name] = key
		delete_row := fmt.Sprintf("%v", indexKey)
		delete(t.IndexTable.Rows, delete_row)
	}
	return nil
}
