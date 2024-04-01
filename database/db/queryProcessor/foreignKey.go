package queryProcessor

// func (q *queryManager) dataInForeignKeyExisted(row *structure.Row, columnName string) bool {
// 	for _, foreignKey := range q.Table.Metadata.ForeignKeys {
// 		if foreignKey.ColumnName == columnName {
// 			refTable := foreignKey.RefTable.IndexTable
// 			key := fmt.Sprintf("%v", row.Data[colIndex])
// 			indexKey := make(map[string]interface{})
// 			indexKey[foreignKey.RefColumnName] = key
// 			_, f := refTable.Rows[fmt.Sprintf("%v", indexKey)]
// 			return f
// 		}
// 	}
// 	return true
// }

// func (q *queryManager) CreateForeignKey(name string, columnName string, refTable *structure.Table, refColumnName string) error {
// 	if q.Table == nil {
// 		return errors.New("target table is nil")
// 	}
// 	if !q.hasColumn(columnName) {
// 		return errors.New("specified column does not exist in the table OR specified column is not unique or primaryKey")
// 	}

// 	refq := queryManager{Table: refTable}

// 	if !refq.hasColumn(refColumnName) {
// 		return errors.New("referenced column does not exist in the target table OR specified column is not unique or primaryKey")
// 	}
// 	var index int
// 	for i, col := range q.Table.Columns {
// 		if col.Name == columnName {
// 			index = i
// 			break
// 		}
// 	}
// 	q.Table.Columns[index].ForeignKey = true
// 	q.Table.Metadata.ForeignKeys = append(q.Table.Metadata.ForeignKeys, structure.ForeignKey{
// 		Name:          name,
// 		ColumnName:    columnName,
// 		RefTable:      refTable,
// 		RefColumnName: refColumnName,
// 		RefTableName:  refTable.Metadata.Name,
// 	})

// 	if err := filemanager.UpdateTableFile(q.Table); err != nil {
// 		return fmt.Errorf("failed to update metadata file: %v", err)
// 	}
// 	return nil
// }
