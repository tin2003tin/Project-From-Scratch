package queryProcessor

import (
	"database/db/structure"
	"fmt"
)

// Join performs a join operation between the current table and another table
func (q *QueryManager) Join(another *structure.Table, joinType structure.JoinType, on structure.On) (*[][]*structure.Row, error) {
	// Validate input parameters
	if another == nil {
		return nil, fmt.Errorf("another table is nil")
	}
	if q.CurrentRows == nil {
		currRows, err := q.Where(&[]structure.Condition{})
		if err != nil {
			return nil, err
		}
		q.CurrentRows = currRows
	}

	// Check if both tables have data
	if len(q.CurrentRows[0]) == 0 || len(another.Rows) == 0 {
		return nil, fmt.Errorf("one of the tables is empty")
	}
	switch joinType {
	case structure.InnerJoin:
		joinedData, err := performInnerJoin(q, another, on)
		if err != nil {
			return nil, err
		}
		return joinedData, nil

	case structure.LeftJoin:
		// Perform left join
		// Implement your logic here
	case structure.RightJoin:
		// Perform right join
		// Implement your logic here
	case structure.FullJoin:
		// Perform full join
		// Implement your logic here
	default:
		return nil, fmt.Errorf("unsupported join type")
	}

	return nil, fmt.Errorf("join operation failed")
}

func performInnerJoin(self *QueryManager, another *structure.Table, on structure.On) (*[][]*structure.Row, error) {
	var exitedColumns map[string]int = map[string]int{}
	for _, selfCol := range self.CurrentColumns {
		exitedColumns[selfCol.Name] = 1
	}
	for _, col := range another.Columns {
		new_col := col
		if _, existed := exitedColumns[col.Name]; existed {
			exitedColumns[col.Name]++
			new_col.Name = fmt.Sprintf("%s_%d", col.Name, exitedColumns[col.Name])
		}
		self.CurrentColumns = append(self.CurrentColumns, new_col)
	}

	_, err := JoinAll(self, another, on)
	if err != nil {
		return nil, err
	}
	self.CurrentIndexes = append(self.CurrentIndexes, another.IndexTable)

	return &self.CurrentRows, nil
}

func JoinAll(self *QueryManager, another *structure.Table, on structure.On) (*[][]*structure.Row, error) {
	var joinedRows [][]*structure.Row

	var mappingSelfCol map[string]int = make(map[string]int)
	for i, selfcol := range self.CurrentColumns {
		mappingSelfCol[selfcol.Name] = i
	}
	var mappingAnotherCol map[string]int = make(map[string]int)
	for i, anotherCol := range another.Columns {
		mappingAnotherCol[anotherCol.Name] = i
	}

	for _, rows := range self.CurrentRows {
		for _, row_t := range rows {
			for _, anotherRow := range another.Rows {
				if checkOnCondition(*row_t, anotherRow, on, mappingSelfCol, mappingAnotherCol) {
					joinedRows = append(joinedRows, []*structure.Row{row_t, &anotherRow})
					break
				}
			}
		}
	}
	self.CurrentRows = joinedRows
	return &self.CurrentRows, nil
}

func checkOnCondition(row1, row2 structure.Row, on structure.On, mappingSelfCol map[string]int, mappingAnotherCol map[string]int) bool {
	switch on.Operator {
	case "=":
		return row1.Data[mappingSelfCol[on.Self]] == row2.Data[mappingAnotherCol[on.Another]]
	case "!=":
		return row1.Data[mappingSelfCol[on.Self]] != row2.Data[mappingAnotherCol[on.Another]]
	}
	return false
}
