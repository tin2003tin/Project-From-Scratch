package buffermanager

import (
	"database/db/lib"
	"database/db/structure"
)

func BuildIndex(t *structure.Table) error {
	t.IndexTable.Rows = make(map[string]*structure.Row)
	t.IndexTable.Columns = make(map[string]*structure.Column)
	for _, column := range t.Columns {
		for _, ci := range t.IndexTable.Include {
			if ci == column.Name {
				t.IndexTable.Columns[ci] = &column
				for _, row := range t.Rows {
					data := row.Data[column.Index]
					indexKey := lib.MappingIndexKey(column.Name, data)
					t.IndexTable.Rows[indexKey] = &row
				}
			}
		}
	}
	return nil
}
