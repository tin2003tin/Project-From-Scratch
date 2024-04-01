package queryProcessor

import (
	"database/db/lib"
	filemanager "database/db/storageManager/fileManager"
	"database/db/structure"
	"fmt"
)

func (q *QueryManager) AddIndex(columnName string) error {
	// Create the index object
	if q.Table.IndexTable == nil {
		q.Table.IndexTable = &structure.Index{
			Name:       "default_index",
			Columns:    make(map[string]*structure.Column),
			Rows:       make(map[string]*structure.Row),
			Unique:     true,
			Comment:    "This is a default index",
			Include:    make([]string, 0),
			Predicate:  "",
			FillFactor: 0,
		}
	}
	// Populate the index columns
	found := false
	for _, column := range q.Table.Columns {
		if column.Name == columnName {
			q.Table.IndexTable.Columns[columnName] = &column
			q.Table.IndexTable.Include = append(q.Table.IndexTable.Include, columnName)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("column '%v' not found in the table", columnName)
	}
	// Update the index file (.tti) and write the index data
	if err := filemanager.UpdateIndexFile(q.Table); err != nil {
		return fmt.Errorf("failed to updated index file: %v", err)
	}

	return nil
}

func (q *QueryManager) AddRowToIndex(i *structure.Index, row structure.Row) error {
	// Extract relevant values from the new row based on the indexed columns
	for _, col := range i.Columns {
		indexKey := lib.MappingIndexKey(col.Name, row.Data[col.Index])
		i.Rows[indexKey] = &row
	}
	return nil
}
