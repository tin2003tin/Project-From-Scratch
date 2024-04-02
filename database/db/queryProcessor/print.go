package queryProcessor

import (
	"database/db/structure"
	"fmt"
	"strings"
)

func (q *QueryManager) PrintAsTable() {
	// if q.CurrentRows == nil {
	// 	q.Where(&[]structure.Condition{})
	// }
	// Calculate column widths
	columns := q.CurrentColumns
	rows := q.CurrentRows
	colWidths := make([]int, len(columns))
	for i, col := range columns {
		colWidths[i] = len(col.Name)
	}
	for _, row := range rows {
		index := 0
		for _, t_row := range row {
			for i, col := range (columns)[index : index+len(t_row.Data)] {
				valueWidth := len(fmt.Sprintf("%v", (t_row).Data[col.Index]))
				if valueWidth > colWidths[i] {
					colWidths[i] = valueWidth
				}
			}
			index += len(t_row.Data)
		}
	}

	// Print column headers
	headerRow := "|"
	for i, col := range columns {
		headerRow += fmt.Sprintf(" %-*s |", colWidths[i], col.Name)
	}
	fmt.Println(headerRow)

	// Print separator line
	separator := "+" + strings.Repeat("-", len(headerRow)-2) + "+"
	fmt.Println(separator)

	// Print rows

	for _, row := range rows {
		index := 0
		dataRow := "|"
		for _, t_row := range row {
			if index < len(columns) {
				for _, col := range (columns)[index : index+len(t_row.Data)] {
					value := fmt.Sprintf("%v", t_row.Data[col.Index])
					formattedValue := fmt.Sprintf(" %-*v ", colWidths[colIndex(col, columns)], value)
					dataRow += formattedValue + "|"
				}
				index += len(t_row.Data)
			}
		}

		fmt.Println(dataRow)
	}
	fmt.Println(separator)
}

// colIndex returns the index of the column in the columns slice
func colIndex(col structure.Column, columns []structure.Column) int {
	for i, c := range columns {
		if c.Name == col.Name {
			return i
		}
	}
	return -1 // Column not found
}
