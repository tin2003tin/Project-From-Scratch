package queryProcessor

import (
	"database/db/lib"
	"database/db/structure"
	"errors"
)

// QueryRows returns rows from the table that match the specified conditions
func (q *QueryManager) Where(conditions *[]structure.Condition) ([][]*structure.Row, error) {
	var mappingColumn map[string]int = make(map[string]int)
	for i, column := range q.CurrentColumns {
		mappingColumn[column.Name] = i
	}

	if q.CurrentRows == nil {
		table_matchedRows, err := q.findwithoutPointer(conditions, &q.Table.Rows, mappingColumn)
		if err != nil {
			return nil, err
		}
		q.CurrentIndexes = []*structure.Index{q.Table.IndexTable}
		q.CurrentRows = table_matchedRows
		return table_matchedRows, nil
	} else {
		for _, cond := range *conditions {
			rowIndexes, realIndexes := lib.CreateRefIndex(&q.CurrentRows, q.CurrentColumns)
			_, err := q.findwithPointer(&cond, &q.CurrentRows, rowIndexes[mappingColumn[cond.ColumnName]], realIndexes[mappingColumn[cond.ColumnName]])
			if err != nil {
				return nil, err
			}
		}
	}

	return q.CurrentRows, nil
}

func (q *QueryManager) findwithPointer(cond *structure.Condition, rows *[][]*structure.Row, rowindex int, realindex int) ([]*structure.Row, error) {
	var matchedRows [][]*structure.Row
	if !(q.hasColumnC(cond.ColumnName)) {
		return nil, errors.New(cond.ColumnName + " is not found in this table")
	}
	for _, t_row := range *rows {
		matched := true
		value := (t_row)[rowindex].Data[realindex]
		m, err := checkCondition(value, cond)
		if err != nil {
			return nil, err
		}

		matched = m
		if matched {
			matchedRows = append(matchedRows, t_row)
		}
	}
	q.CurrentRows = matchedRows
	return nil, nil
}

func (q *QueryManager) findwithoutPointer(conditions *[]structure.Condition, rows *[]structure.Row, mappingColumn map[string]int) ([][]*structure.Row, error) {
	var table_matchedRows [][]*structure.Row
	for _, row := range *rows {
		matched := true
		for _, cond := range *conditions {
			if !(q.hasColumnC(cond.ColumnName)) {
				return nil, errors.New(cond.ColumnName + " is not found in this table")
			}
			value := row.Data[mappingColumn[cond.ColumnName]]
			m, err := checkCondition(value, &cond)
			if err != nil {
				return nil, err
			}
			matched = m
		}

		if matched {
			table_matchedRows = append(table_matchedRows, []*structure.Row{&row})
		}
	}
	return table_matchedRows, nil
}
