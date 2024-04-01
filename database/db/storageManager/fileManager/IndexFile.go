package filemanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func UpdateIndexFile(t *structure.Table) error {
	indexFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".tti")

	// Open or create the index file
	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to update index file: %v", err)
	}
	defer indexFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(indexFile)

	indexTable := structure.Index{
		Name:       t.IndexTable.Name,
		Columns:    nil,
		Rows:       nil,
		Unique:     t.IndexTable.Unique,
		Using:      t.IndexTable.Using,
		Comment:    t.IndexTable.Comment,
		Tablespace: t.IndexTable.Tablespace,
		Include:    t.IndexTable.Include,
		Predicate:  t.IndexTable.Predicate,
		FillFactor: t.IndexTable.FillFactor,
	}
	// Encode and write the index data to the file
	if err := encoder.Encode(indexTable); err != nil {
		return fmt.Errorf("failed to encode index data: %v", err)
	}

	fmt.Println("Index file updated successfully:", indexFilePath)
	return nil
}
