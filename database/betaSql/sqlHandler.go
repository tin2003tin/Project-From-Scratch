package betasql

import (
	"fmt"
)

func empty(item [][]string) ([]string, error) {
	fmt.Println("Empty naja")
	return nil, nil
}
func (sql *Sql) loadColumn(item [][]string) ([]string, error) {
	return []string{item[0][len(item[len(item)-1])-1]}, nil
}

// func addColumn(item [][]string) ([]string, error) {
// 	var modified []string
// 	modified = append(modified, item[2]...)
// 	modified = append(modified, item[0]...)
// 	return modified, nil
// }

// func loadTable(item [][]string) ([]string, error) {
// 	return []string{item[len(item)-1][len(item[len(item)-1])-1]}, nil
// }

// func loadOp(item [][]string) ([]string, error) {
// 	operators := map[string]bool{
// 		">":  true,
// 		"<":  true,
// 		"=":  true,
// 		"!=": true,
// 		"<=": true,
// 		">=": true,
// 	}

// 	if len(item) == 0 || len(item[len(item)-1]) == 0 {
// 		return nil, errors.New("invalid input: empty item slice or empty sub-slice")
// 	}

// 	op := item[len(item)-1][len(item[len(item)-1])-1]
// 	if operators[op] {
// 		return []string{op}, nil
// 	}

// 	return nil, fmt.Errorf("%s is not an Operator", op)
// }

// func loadType(item [][]string) ([]string, error) {
// 	if len(item) == 0 || len(item[len(item)-1]) == 0 {
// 		return nil, errors.New("invalid input: empty item slice or empty sub-slice")
// 	}

// 	lastSubSlice := item[len(item)-1]
// 	lastElement := lastSubSlice[len(lastSubSlice)-1]

// 	if strings.HasPrefix(lastElement, "'") && strings.HasSuffix(lastElement, "'") {
// 		return []string{lastElement[1 : len(lastElement)-1], "string"}, nil
// 	}

// 	if _, err := strconv.Atoi(lastElement); err == nil {
// 		return []string{lastElement, "number"}, nil
// 	}

// 	return nil, fmt.Errorf("unsupported type: %s", lastElement)
// }

// func condition(item [][]string) ([]string, error) {
// 	if len(item) == 0 {
// 		return nil, fmt.Errorf("empty item slice")
// 	}
// 	var combinedConditions []string
// 	for i, j := 0, len(item)-1; i < j; i, j = i+1, j-1 {
// 		item[i], item[j] = item[j], item[i]
// 	}
// 	for _, subSlice := range item {
// 		if len(subSlice) == 0 {
// 			continue
// 		}
// 		combinedConditions = append(combinedConditions, strings.Join(subSlice, " "))
// 	}
// 	combinedString := strings.Join(combinedConditions, " ")
// 	fmt.Println(combinedString)
// 	return []string{combinedString}, nil
// }

// func loadCondition(item [][]string) ([]string, error) {
// 	if len(item) == 0 {
// 		return nil, errors.New("empty item slice")
// 	}

// 	firstElement := item[0]
// 	if len(firstElement) == 0 {
// 		return nil, errors.New("first sub-slice is empty")
// 	}

// 	result := firstElement[0]
// 	fmt.Println(result)
// 	return []string{result}, nil
// }

// func loadJoin(item [][]string) ([]string, error) {
// 	fmt.Println(item)
// 	return []string{item[4][0], item[2][0], item[0][0]}, nil
// }

// func loadSql(item [][]string) ([]string, error) {
// 	// _select := item[4]
// 	_where := item[0]
// 	_from := item[2]
// 	// _join := item[1]
// 	database, err := db.GetDataBase("H_market")
// 	if err != nil {
// 		return nil, err
// 	}
// 	t, err := database.GetTable(_from[0])
// 	if err != nil {
// 		return nil, err
// 	}
// 	if len(_where) > 0 {
// 		var conditions []table.Condition
// 		var value interface{}
// 		where := strings.Split(_where[0], " ")
// 		value = where[2]
// 		if where[3] == "number" {
// 			value, _ = strconv.ParseFloat(where[2], 64)
// 		}
// 		conditions = append(conditions, table.Condition{ColumnName: where[0], Operator: where[1], Value: value})
// 		fmt.Println(conditions)
// 		row, _ := t.Where(conditions)
// 		table.PrintAsTable(t.Metadata.Columns, row)
// 	} else {
// 		t.PrintAsTable()
// 	}
// 	return nil, nil
// }
