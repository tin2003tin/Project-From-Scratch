package betasql

import (
	"database/betaSql/lib"
	"database/db/queryProcessor"
	"database/db/structure"
	"fmt"
	"strconv"
)

func (sdb *SelectDatabase) deleteRow(item [][]string) ([][]string, error) {
	_table := item[1]
	_where := item[0]

	var loadtables []string
	loadtables = append(loadtables, _table[0])
	tables, err := sdb.loadMultiFile(loadtables)
	if err != nil {
		return nil, err
	}

	queryTable := queryProcessor.NewQueryManager(tables[0])
	var delete_rows [][]*structure.Row
	if _where[0] != "nil" {
		var where_value interface{} = _where[0]
		switch _where[1] {
		case "number":
			numVal, err := strconv.Atoi(_where[0])
			if err != nil {
				return nil, err
			}
			where_value = numVal
		}
		cond := structure.Condition{ColumnName: _where[3], Operator: _where[2], Value: where_value}
		delete_rows, err = queryTable.DeleteRow([]structure.Condition{cond})
		if err != nil {

			return nil, err
		}
	} else {
		delete_rows, err = queryTable.DeleteRow([]structure.Condition{})
		if err != nil {
			return nil, err
		}
	}
	queryTable.Commit()
	queryTable.ResetCurrent()
	
	var mappedRows map[string]interface{} = map[string]interface{}{}
	if len(delete_rows) > 0 {
		fmt.Println(delete_rows[0][0])
		for i, col := range queryTable.CurrentColumns {
			mappedRows[col.Name] = delete_rows[0][0].Data[i]
		}
	}
	json, err := lib.ConvertToJson(mappedRows)
	if err != nil {
		return nil, err
	}
	sdb.Output = json

	return nil, nil
}
