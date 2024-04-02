package betasql

import (
	"database/betaSql/lib"
	"database/db/queryProcessor"
	"database/db/structure"
	"fmt"
	"strconv"
	"time"
)

func (sdb *SelectDatabase) loadSql(item [][]string) ([][]string, error) {
	startLoadDatabase := time.Now()
	fmt.Println(item)
	_select := item[4]
	_from := item[2]
	_where := item[0]
	_join := item[1]

	var loadtables []string
	loadtables = append(loadtables, _from[0])
	if _join[0] != "nil" {
		loadtables = append(loadtables, _join[3])
	}
	tables, err := sdb.loadMultiFile(loadtables)
	if err != nil {
		return nil, err
	}

	queryTable := queryProcessor.NewQueryManager(tables[0])
	if _join[0] != "nil" {
		queryTable.JoinWithIndex(tables[1], structure.InnerJoin, structure.On{Self: _join[0], Operator: _join[1], Another: _join[2]})
	}

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
		_, err := queryTable.Where(&[]structure.Condition{cond})
		if err != nil {

			return nil, err
		}
		// queryTable.WhereWithIndex(&[]structure.Condition{cond})
	} else {
		_, err := queryTable.Where(&[]structure.Condition{})
		if err != nil {
			return nil, err
		}
	}

	queryTable.PrintAsTable()
	var selectedRows []map[string]interface{}
	if _select[0] == "*" {
		s, err := queryTable.Select(true, nil)
		if err != nil {
			return nil, err
		}
		selectedRows = s
	} else {
		var selectedcolumn []string
		selectedcolumn = append(selectedcolumn, _select...)
		s, err := queryTable.Select(false, selectedcolumn)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		selectedRows = s
	}
	json, err := lib.ConvertToJson(selectedRows)
	if err != nil {
		return nil, err
	}
	sdb.Output = json
	loadDuration := time.Since(startLoadDatabase)
	fmt.Println("Execution Time for only Excute:", loadDuration)
	queryTable.ResetCurrent()
	return nil, nil
}
