package buffermanager

import (
	"database/db/structure"
	"errors"
)

func BuildForeignKey(db *structure.Database, t *structure.Table) error {
	for i := range t.Metadata.ForeignKeys {
		refTable, ok := db.Registry.Tables[t.Metadata.ForeignKeys[i].RefTableName]
		if !ok {
			return errors.New("not found the Reftable " + t.Metadata.ForeignKeys[i].RefTableName)
		}
		t.Metadata.ForeignKeys[i].RefTable = refTable
	}
	return nil
}
