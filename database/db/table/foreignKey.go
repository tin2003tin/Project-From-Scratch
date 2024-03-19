package table

import (
	"errors"
	"fmt"
)

func (t *Table) dataInForeignKeyExisted(row *Row, columnName string) bool {
	for _, foreignKey := range t.Metadata.ForeignKeys {
		if foreignKey.ColumnName == columnName {
			refTable := foreignKey.RefTable.IndexTable
			key := fmt.Sprintf("%v", row.Data[columnName])
			indexKey := make(map[string]interface{})
			indexKey[foreignKey.RefColumnName] = key
			_, f := refTable.Rows[fmt.Sprintf("%v", indexKey)]
			return f
		}
	}
	return true
}

func (t *Table) CreateForeignKey(name string, columnName string, refTable *Table, refColumnName string) error {
	if t == nil {
		return errors.New("target table is nil")
	}
	if !t.hasColumn(columnName) {
		return errors.New("specified column does not exist in the table OR specified column is not unique or primaryKey")
	}

	if !refTable.hasColumn(refColumnName) {
		return errors.New("referenced column does not exist in the target table OR specified column is not unique or primaryKey")
	}
	var index int
	for i, col := range t.Metadata.Columns {
		if col.Name == columnName {
			index = i
			break
		}
	}
	t.Metadata.Columns[index].ForeignKey = true
	t.Metadata.ForeignKeys = append(t.Metadata.ForeignKeys, ForeignKey{
		Name:          name,
		ColumnName:    columnName,
		RefTable:      refTable,
		RefColumnName: refColumnName,
		RefTableName:  refTable.Metadata.Name,
	})
	t.updateMetadataFile()
	return nil
}

func (t *Table) hasColumn(columnName string) bool {
	for _, col := range t.Metadata.Columns {
		if col.Name == columnName {
			return true
		}
	}
	return false
}
