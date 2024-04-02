package betasql

import (
	"database/betaSql/lib"
	"database/db/queryProcessor"
	"fmt"
	"strconv"
	"strings"
)

func (sdb *SelectDatabase) loadValue(item [][]string) ([][]string, error) {
	return item, nil
}

func (sdb *SelectDatabase) addValue(item [][]string) ([][]string, error) {
	return [][]string{append(item[0], item[2][0])}, nil
}

func (sdb *SelectDatabase) loadColumns(item [][]string) ([][]string, error) {
	return item, nil
}

func (sdb *SelectDatabase) loadInsert(item [][]string) ([][]string, error) {
	_table := item[5]
	_values := item[1]

	var loadtables []string
	loadtables = append(loadtables, _table[0])
	tables, err := sdb.loadMultiFile(loadtables)
	if err != nil {
		return nil, err
	}

	queryTable1 := queryProcessor.NewQueryManager(tables[0])

	var row []interface{}
	for _, val := range _values {
		if strings.HasPrefix(val, "'") {
			row = append(row, val[1:len(val)-1])
		} else {
			intValue, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("cannot convert value to int")
			}
			row = append(row, intValue)
		}

	}

	for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
		row[i], row[j] = row[j], row[i]
	}

	err = queryTable1.AddRow(row)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	queryTable1.Commit()
	queryTable1.ResetCurrent()

	var addedRow map[string]interface{} = map[string]interface{}{}
	for i, col := range queryTable1.CurrentColumns {
			addedRow[col.Name] = row[i]
	}
	json, err := lib.ConvertToJson(addedRow)
	if err != nil {
		return nil, err
	}
	sdb.Output = json

	return nil, nil
}
