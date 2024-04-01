package queryProcessor

import (
	"database/db/lib"
	"database/db/structure"
	"errors"
	"fmt"
	"reflect"
)

type Set struct {
	ColumnName string
	Value      interface{}
}

func (q *QueryManager) UpdateRow(conditions []structure.Condition, sets []Set) error {
	if q.Table == nil {
		return fmt.Errorf("table is nil")
	}

	var mappingColumn map[string]int = make(map[string]int)
	for i, column := range q.CurrentColumns {
		mappingColumn[column.Name] = i

	}
	for i := len(q.Table.Rows) - 1; i >= 0; i-- {
		row := q.Table.Rows[i]
		for _, cond := range conditions {
			if passed, _ := checkCondition(row.Data[mappingColumn[cond.ColumnName]], &cond); passed {
				for _, set := range sets {
					for _, column := range q.Table.Columns {
						if column.PrimaryKey || column.Unique {
							_, found := q.Table.IndexTable.Rows[lib.MappingIndexKey(set.ColumnName, set.Value)]
							if found {
								return errors.New("cannot update row, primary key or unique already exists: " + set.ColumnName + fmt.Sprintf(" %v", set.Value))
							}
						}
						// if column.ForeignKey {
						// 	if !t.dataInForeignKeyExisted(&struc Row{Data: map[string]interface{}{set.ColumnName: set.Value}}, column.Name) {
						// 		return fmt.Errorf("cannot update row, %v not found in foreign key column '%s'", set.Value, column.Name)
						// 	}
						// }
						if column.Name == set.ColumnName {
							newValue, err := lib.ConvertValue(set.Value, column.DataType, column.Length)
							if err != nil {
								return fmt.Errorf("cannot update, %v %v", column.Name, err)
							}
							set.Value = newValue
						}
					}

					if err := updateFromIndex(q.Table, &row, set); err != nil {
						return err
					}
					if err := updateRow(q.Table, &row, set); err != nil {
						return err
					}
				}
			}

		}
	}

	return nil
}

func updateRow(t *structure.Table, row *structure.Row, set Set) error {
	for i := range t.Columns {
		if set.ColumnName == t.Columns[i].Name {
			row.Data[i] = set.Value
			break
		}
	}
	return nil
}

func updateFromIndex(t *structure.Table, row *structure.Row, set Set) error {
	indexColumns := t.IndexTable.Columns
	indexRows := t.IndexTable.Rows
	for _, column := range indexColumns {
		if set.ColumnName == column.Name {
			newTemp := make(map[string]interface{})
			oldTemp := make(map[string]interface{})
			key := fmt.Sprintf("%v", row.Data[column.Index])
			oldTemp[column.Name] = key
			newTemp[column.Name] = set.Value
			oldkey := fmt.Sprintf("%v", oldTemp)
			newKey := fmt.Sprintf("%v", newTemp)

			indexRows[newKey] = indexRows[oldkey]
			delete(indexRows, oldkey)
		}
	}
	return nil
}

func checkCondition(value interface{}, cond *structure.Condition) (bool, error) {
	switch v := value.(type) {
	case int:
		if !lib.CompareInt(v, cond.Operator, cond.Value) {
			return false, nil
		}
		return true, nil
	case float64:
		if !lib.CompareFloat64(v, cond.Operator, cond.Value) {
			return false, nil
		}
		return true, nil

	case string:
		if !lib.CompareString(v, cond.Operator, cond.Value) {
			return false, nil
		}
		return true, nil
	// Add more cases for other types as needed
	default:
		return false, errors.New("unsupported data type for comparison: " + reflect.TypeOf(value).String())
	}
}
