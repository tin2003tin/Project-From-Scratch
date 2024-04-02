package betasql

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
