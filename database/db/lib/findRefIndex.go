package lib

import (
	"database/db/structure"
)

func CreateRefIndex(rows *[][]*structure.Row, cols []structure.Column) ([]int, []int) {
	var rowIndex []int = []int{}
	var realindex []int = []int{}
	if len(*rows) == 0 {
		for j := range cols {
			realindex = append(realindex, j)
			rowIndex = append(rowIndex, 0)
		}
	} else {
		for i, row := range (*rows)[0] {
			for j := range (*row).Data {
				realindex = append(realindex, j)
				rowIndex = append(rowIndex, i)
			}
		}
	}

	return rowIndex, realindex
}
