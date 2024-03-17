package table

import (
	"fmt"
)

// JoinType represents the type of join operation
type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullJoin
)

type On struct {
	Self string // Name of the column to check condition against
	Operator string // Operator for comparison (e.g., "=", ">", "<", etc.)
	Another string
}

// Join performs a join operation between the current table and another table
func (t *Table) Join(another *Table, joinType JoinType, conditions []On) (*Table, error) {
	// Validate input parameters
	if another == nil {
		return nil, fmt.Errorf("another table is nil")
	}
	if len(conditions) == 0 {
		return nil, fmt.Errorf("no join conditions provided")
	}

	// Check if both tables have data
	if len(t.Metadata.Rows) == 0 || len(another.Metadata.Rows) == 0 {
		return nil, fmt.Errorf("one of the tables is empty")
	}
	switch joinType {
	case InnerJoin:
		joinedData,err := t.performInnerJoin(another, conditions)
		if (err != nil) {
			return nil, err
		}
		return joinedData, nil
		
	case LeftJoin:
		// Perform left join
		// Implement your logic here
	case RightJoin:
		// Perform right join
		// Implement your logic here
	case FullJoin:
		// Perform full join
		// Implement your logic here
	default:
		return nil, fmt.Errorf("unsupported join type")
	}

	return nil, fmt.Errorf("join operation failed")
}

func (t *Table) performInnerJoin(another *Table, conditions []On) (*Table, error) {
	joinedRows := make([]Row, 0)
	joinedColumns := make([]Column, 0) 

	joinedColumns  = append(joinedColumns , t.Metadata.Columns...)
	joinedColumns  = append(joinedColumns , another.Metadata.Columns...)

	for _, row := range t.Metadata.Rows {
		for _, anotherRow := range another.Metadata.Rows {
			if t.checkAllConditions(row, anotherRow, conditions) {
				joinedRow := Row{
					Data: make(map[string]interface{}),
				}
				for key, value := range row.Data {
					joinedRow.Data[key] = value
				}

				for key, value := range anotherRow.Data {
					joinedRow.Data[key] = value
				}

				joinedRows = append(joinedRows, joinedRow)
			}
		}
	}
	return &Table{Metadata: TableMetadata{Rows: joinedRows,Columns: joinedColumns }}, nil
}

func (t *Table) checkAllConditions(row1, row2 Row, conditions []On) bool {
	for _, condition := range conditions {
		if !t.checkSingleCondition(row1, row2, condition) {
			return false
		}
	}
	return true
}

func (t *Table) checkSingleCondition(row1, row2 Row, condition On) bool {
	switch (condition.Operator) {
		case "=" :
			return row1.Data[condition.Self] == row2.Data[condition.Another]
		case "!=" :
			return row1.Data[condition.Self] != row2.Data[condition.Another]
	}
	return false;
}