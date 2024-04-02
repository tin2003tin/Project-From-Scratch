package queryProcessor

import (
	"database/db/lib"
	"fmt"
)

func (q *QueryManager) Select(all bool, colNames []string) ([]map[string]interface{}, error) {
	// if q.CurrentRows == nil {
	// 	currRows, err := q.Where(&[]structure.Condition{})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	q.CurrentRows = currRows
	// }
	var selectedIndex []int = []int{}

	if all {
		for i, _ := range q.CurrentColumns {
			selectedIndex = append(selectedIndex, i)
		}
	}
	for _, name := range colNames {
		existed := false
		for i, col := range q.CurrentColumns {
			if col.Name == name {
				existed = true
				selectedIndex = append(selectedIndex, i)
				break
			}
		}
		if !existed {
			return nil, fmt.Errorf("column %s not found in columnNames", name)
		}
	}
	rowIndexes, realIndexes := lib.CreateRefIndex(&q.CurrentRows, q.CurrentColumns)
	var selectedRows []map[string]interface{}
	for _, t_row := range q.CurrentRows {
		var selectedIndexInRow map[string]interface{} = make(map[string]interface{})
		for _, index := range selectedIndex {
			value := (t_row)[rowIndexes[index]].Data[realIndexes[index]]
			selectedIndexInRow[q.CurrentColumns[index].Name] = value
		}
		selectedRows = append(selectedRows, selectedIndexInRow)
	}
	return selectedRows, nil
}
