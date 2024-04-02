package betasql

import (
	"database/betaSql/lib"
	"database/db/queryProcessor"
	"database/db/structure"
	"fmt"
	"strconv"
)

func (sdb *SelectDatabase) loadSet(item [][]string) ([][]string, error) {
	fmt.Println("loadSet")
	items := []string{}
	for it := range item {
		items = append(items, item[it]...)
	}
	return [][]string{items}, nil
}

func (sdb *SelectDatabase) loadSets(item [][]string) ([][]string, error) {
	fmt.Println("load Sets")
	return item, nil
}

func (sdb *SelectDatabase) addSet(item [][]string) ([][]string, error) {
	items := []string{}
	for it := range item {
		if item[it][0] != "," {
			items = append(items, item[it]...)
		}
	}

	return [][]string{items}, nil
}

func (sdb *SelectDatabase) updateRow(item [][]string) ([][]string, error) {
	fmt.Println(item)
	_where := item[0]
	_sets := item[1]
	_table := item[3]

	var loadtables []string
	loadtables = append(loadtables, _table[0])
	tables, err := sdb.loadMultiFile(loadtables)
	if err != nil {
		return nil, err
	}

	queryTable := queryProcessor.NewQueryManager(tables[0])
	var condition structure.Condition
	if _where[0] != "nil" {
		var value interface{}
		value = _where[0]
		if _where[1] == "number" {
			intValue, err := strconv.Atoi(_where[0])
			if err != nil {
				return nil, fmt.Errorf("cannot convert value to int")
			}
			value = intValue
		}
		condition = structure.Condition{ColumnName: _where[3], Operator: _where[2], Value: value}
	}
	var sets []queryProcessor.Set = []queryProcessor.Set{}
	fmt.Println(len(_sets))
	for i := 0; i < len(_sets); i += 4 {
		if i%4 == 0 {
			col := _sets[i+3]
			var value interface{}
			value = _sets[i]
			if _sets[i+1] == "number" {
				intValue, err := strconv.Atoi(_sets[i])
				if err != nil {
					return nil, fmt.Errorf("cannot convert value to int")
				}
				value = intValue
			}
			sets = append(sets, queryProcessor.Set{ColumnName: col, Value: value})
		}

	}
	fmt.Println(condition)
	fmt.Println(sets)
	err = queryTable.UpdateRow([]structure.Condition{condition}, sets)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	queryTable.Commit()
	queryTable.ResetCurrent()

	var updated map[string]interface{} = make(map[string]interface{})
	for _, set := range sets {
		updated[set.ColumnName] = set.Value
	}

	json, err := lib.ConvertToJson(updated)
	if err != nil {
		return nil, err
	}

	sdb.Output = json

	return nil, nil
}
