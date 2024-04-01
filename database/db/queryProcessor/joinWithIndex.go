package queryProcessor

import (
	"database/db/lib"
	"database/db/structure"
	"fmt"
)

func (q *QueryManager) JoinWithIndex(another *structure.Table, joinType structure.JoinType, on structure.On) (*[][]*structure.Row, error) {
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
		joinedData, err := performInnerJoinWithIndex(q, another, on)
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

func performInnerJoinWithIndex(self *QueryManager, another *structure.Table, on structure.On) (*[][]*structure.Row, error) {
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
	if _, ok := another.IndexTable.Columns[on.Another]; ok {
		_, err := JoinAllWithIndex(self, another, on)
		if err != nil {
			return nil, err
		}
	}
	// } else {
	// 	joinedRows, _ = t.JoinWithIndex(another, on)
	// }

	return &self.CurrentRows, nil
}

func JoinAllWithIndex(self *QueryManager, another *structure.Table, on structure.On) (*[][]*structure.Row, error) {
	var joinedRows [][]*structure.Row

	var mappingSelfCol map[string]int = make(map[string]int)
	for i, selfcol := range self.Table.Columns {
		mappingSelfCol[selfcol.Name] = i
	}

	for _, rows := range self.CurrentRows {
		for _, row_t := range rows {
			value := row_t.Data[mappingSelfCol[on.Self]]
			if ptr, ok := another.IndexTable.Rows[lib.MappingIndexKey(on.Another, value)]; ok {
				joinedRows = append(joinedRows, []*structure.Row{row_t, ptr})
				break
			}
		}
	}
	self.CurrentRows = joinedRows
	return &self.CurrentRows, nil
}
