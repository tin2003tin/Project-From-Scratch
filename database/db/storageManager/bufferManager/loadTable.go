package buffermanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
)

func LoadTableMetadata(db *structure.Database, name string) (*structure.Table, error) {
	found := false
	for _, tablename := range db.TableNames {
		if tablename == name {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("the table's name is not found. " + name)
	}
	tableFile := db.MetadataPath + "/" + name + ".ttm"

	// Open the table metadata file for reading
	tableMetaFile, err := os.Open(tableFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open table metadata file '%s': %v", tableFile, err)
	}
	defer tableMetaFile.Close()

	// Initialize a decoder for reading binary data
	tableDecoder := gob.NewDecoder(tableMetaFile)

	// Decode the table metadata from the file
	var table structure.Table
	if err := tableDecoder.Decode(&table); err != nil {
		return nil, fmt.Errorf("failed to decode table metadata from '%s': %v", tableFile, err)
	}
	db.Tables = append(db.Tables, &table)
	db.Registry.Tables[table.Metadata.Name] = &table
	return &table, nil
}
