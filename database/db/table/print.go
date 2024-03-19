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
		for _, col := range columns {
			value := fmt.Sprintf("%v", row.Data[col.Name])
			formattedValue := fmt.Sprintf(" %-*v ", colWidths[colIndex(col, columns)], value)
			dataRow += formattedValue + "|"
		}
		fmt.Println(dataRow)
	}
	fmt.Println(separator)
}

// colIndex returns the index of the column in the columns slice
func colIndex(col Column, columns []Column) int {
	for i, c := range columns {
		if c.Name == col.Name {
			return i
		}
	}
	return -1 // Column not found
}

func (t *Table) PrintAsTable() {
	fmt.Println("Table :", t.Metadata.Name)
	PrintAsTable(t.Metadata.Columns, t.Metadata.Rows)
}

func (t *Table) ListColumn() []string {
	columns := make([]string, 0)
	for _, column := range t.Metadata.Columns {
		text := column.Name + ":" + column.DataType
		columns = append(columns, text)
	}
	return columns
}

func ListColumn(t *Table) []string {
	columns := make([]string, 0)
	for _, column := range t.Metadata.Columns {
		text := column.Name + ":" + column.DataType
		columns = append(columns, text)
	}
	return columns
}
