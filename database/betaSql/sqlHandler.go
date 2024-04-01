package betasql

import (
	"database/betaSql/lib"
	"database/db/queryProcessor"
	buffermanager "database/db/storageManager/bufferManager"
	"database/db/structure"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (sdb *SelectDatabase) loadColumn(item [][]string) ([][]string, error) {
	return item, nil
}

func (sdb *SelectDatabase) addColumn(item [][]string) ([][]string, error) {
	return [][]string{append(item[0], item[2][0])}, nil
}

func (sdb *SelectDatabase) loadTable(item [][]string) ([][]string, error) {
	for _, tableName := range sdb.DataBase.TableNames {
		if tableName == item[0][0] {
			return item, nil
		}
	}
	return nil, fmt.Errorf("cannot find table (" + item[0][0] + ")")
}

func (sdb *SelectDatabase) loadOp(item [][]string) ([][]string, error) {
	operators := map[string]bool{
		">":  true,
		"<":  true,
		"=":  true,
		"!=": true,
		"<=": true,
		">=": true,
	}

	if len(item) == 0 || len(item[len(item)-1]) == 0 {
		return nil, errors.New("invalid input: empty item slice or empty sub-slice")
	}

	op := item[len(item)-1][len(item[len(item)-1])-1]
	if operators[op] {
		return item, nil
	}

	return nil, fmt.Errorf("%s is not an Operator", op)
}

func (sdb *SelectDatabase) loadType(item [][]string) ([][]string, error) {
	if len(item) == 0 || len(item[len(item)-1]) == 0 {
		return nil, errors.New("invalid input: empty item slice or empty sub-slice")
	}

	lastSubSlice := item[len(item)-1]
	lastElement := lastSubSlice[len(lastSubSlice)-1]
	if strings.HasPrefix(lastElement, "'") && strings.HasSuffix(lastElement, "'") {
		return [][]string{{lastElement[1 : len(lastElement)-1], "string"}}, nil
	}

	if _, err := strconv.Atoi(lastElement); err == nil {
		return [][]string{{lastElement, "number"}}, nil
	}

	return nil, fmt.Errorf("unsupported type: %s", lastElement)
}

func (sdb *SelectDatabase) condition(item [][]string) ([][]string, error) {
	if len(item) == 0 {
		return nil, fmt.Errorf("empty item slice")
	}
	return item, nil
}

func (sdb *SelectDatabase) loadCondition(item [][]string) ([][]string, error) {
	if len(item) == 0 {
		return nil, errors.New("empty item slice")
	}
	modifiedItem := item[0]
	modifiedItem = append(modifiedItem, item[1][0], item[2][0], item[3][0])
	return [][]string{modifiedItem}, nil
}

func (sdb *SelectDatabase) loadJoin(item [][]string) ([][]string, error) {
	join := item[0]
	join = append(join, item[1][0], item[2][0], item[4][0])
	return [][]string{join}, nil

}

func (sdb *SelectDatabase) loadSql(item [][]string) ([][]string, error) {
	startLoadDatabase := time.Now()
	fmt.Println(item)
	_select := item[4]
	_from := item[2]
	_where := item[0]
	_join := item[1]
	var err error
	var wg sync.WaitGroup
	var table1 *structure.Table
	var table2 *structure.Table
	if _, ok := sdb.DataBase.Registry.Tables[_from[0]]; !ok {
		wg.Add(1)
		go func() {
			defer wg.Done()
			table1, err = buffermanager.LoadTableMetadata(sdb.DataBase, _from[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			err = buffermanager.LoadIndex(table1)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = buffermanager.LoadRawData(table1)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = buffermanager.BuildIndex(table1)
			if err != nil {
				fmt.Println(err)
				return
			}
			queryProcessor.NewQueryManager(table1)
		}()
	} else {
		table1 = sdb.DataBase.Registry.Tables[_from[0]]
	}

	if _join[0] != "nil" {
		if _, ok := sdb.DataBase.Registry.Tables[_join[3]]; !ok {
			wg.Add(1)
			go func() {
				defer wg.Done()
				table2, err = buffermanager.LoadTableMetadata(sdb.DataBase, _join[3])
				if err != nil {
					fmt.Println(err)
					return
				}

				err = buffermanager.LoadIndex(table2)
				if err != nil {
					fmt.Println(err)
					return
				}

				err = buffermanager.LoadRawData(table2)
				if err != nil {
					fmt.Println(err)
					return
				}

				err = buffermanager.BuildIndex(table2)
				if err != nil {
					fmt.Println(err)
					return
				}
			}()
		} else {
			table2 = sdb.DataBase.Registry.Tables[_join[3]]
		}
	}

	wg.Wait()
	queryTable1 := queryProcessor.NewQueryManager(table1)
	if _join[0] != "nil" {
		queryTable1.JoinWithIndex(table2, structure.InnerJoin, structure.On{Self: _join[0], Operator: _join[1], Another: _join[2]})
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
		_, err := queryTable1.Where(&[]structure.Condition{cond})
		if err != nil {
			fmt.Println("test")
			return nil, err
		}
		// queryTable1.WhereWithIndex(&[]structure.Condition{cond})
	}

	queryTable1.PrintAsTable()
	var selectedRows []map[string]interface{}
	if _select[0] == "*" {
		s, err := queryTable1.Select(true, nil)
		if err != nil {
			return nil, err
		}
		selectedRows = s
	} else {
		var selectedcolumn []string
		selectedcolumn = append(selectedcolumn, _select...)
		s, err := queryTable1.Select(false, selectedcolumn)
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
	sdb.Output = &json
	// fmt.Println(_select)
	// fmt.Println(_from)
	// fmt.Println(_where)
	loadDuration := time.Since(startLoadDatabase)
	fmt.Println("Execution Time for only Excute:", loadDuration)
	return nil, nil
}
