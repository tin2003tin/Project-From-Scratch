package table

import (
	"fmt"
	"strings"
)

func PrintAsTable(columns []Column, rows []Row) {
	// Calculate column widths
	colWidths := make([]int, len(columns))
	for i, col := range columns {
		colWidths[i] = len(col.Name)
	}
	for _, row := range rows {
		for i, col := range columns {
			valueWidth := len(fmt.Sprintf("%v", row.Data[col.Name]))
			if valueWidth > colWidths[i] {
				colWidths[i] = valueWidth
			}
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
		dataRow := "|"
		for i, col := range columns {
			dataRow += fmt.Sprintf(" %-*v |", colWidths[i], row.Data[col.Name])
		}
		fmt.Println(dataRow)
	}

	// Print bottom separator line
	fmt.Println(separator)
}

func (t *Table) PrintAsTable(){
	fmt.Println("Table :",t.Metadata.Name)
	PrintAsTable(t.Metadata.Columns,t.Metadata.Rows)
}