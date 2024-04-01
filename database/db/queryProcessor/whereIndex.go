package queryProcessor

import (
	"database/db/lib"
	"database/db/structure"
	"errors"
	"fmt"
)

func (q *QueryManager) WhereWithIndex(conditions *[]structure.Condition) ([][]*structure.Row, error) {
	var mappingColumn map[string]int = make(map[string]int)
	for i, column := range q.CurrentColumns {
		mappingColumn[column.Name] = i
	}

	if q.CurrentRows == nil {
		q.CurrentIndexes = []*structure.Index{q.Table.IndexTable}
		q.CurrentRows = [][]*structure.Row{}
	}

	for _, cond := range *conditions {
		fmt.Println("test1")
		rowIndexes, _ := lib.CreateRefIndex(&q.CurrentRows, q.CurrentColumns)
		matchedRows, err := q.findUsingIndex(&cond, rowIndexes[mappingColumn[cond.ColumnName]])
		if err != nil {
			return nil, err
		}
		q.CurrentRows = matchedRows
	}

	return q.CurrentRows, nil
}

func (q *QueryManager) findUsingIndex(cond *structure.Condition, rowindex int) ([][]*structure.Row, error) {
	var matchedRows [][]*structure.Row
	value, ok := q.CurrentIndexes[rowindex].Rows[lib.MappingIndexKey(cond.ColumnName, cond.Value)]
	if ok {
		matchedRows = append(matchedRows, []*structure.Row{value})
		return matchedRows, nil
	}
	return nil, errors.New("the value is not found or is not in index")

}
