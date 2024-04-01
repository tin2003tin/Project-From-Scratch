package filemanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func CreateTableFile(tbl *structure.Table) error {
	filePath := filepath.Join(tbl.Metadata.MetadataPath, fmt.Sprintf("%s.ttm", tbl.Metadata.Name))
	// Open or create the table metadata file
	metaFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create table metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)
	tb := structure.Table{
		Metadata: tbl.Metadata,
		Columns:  []structure.Column{},
		Rows:     []structure.Row{},
	}

	// Encode and write the table metadata to the file
	if err := encoder.Encode(tb); err != nil {
		return fmt.Errorf("failed to encode table metadata: %v", err)
	}

	fmt.Printf("Table metadata file '%s' created successfully\n", filePath)
	return nil
}

func UpdateTableFile(tbl *structure.Table) error {
	filePath := filepath.Join(tbl.Metadata.MetadataPath, fmt.Sprintf("%s.ttm", tbl.Metadata.Name))
	// Open or create the table metadata file
	metaFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create table metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	new_tm := structure.Table{
		Metadata: tbl.Metadata,
		Columns:  tbl.Columns,
		Rows:     []structure.Row{},
	}

	for _, foreignKey := range new_tm.Metadata.ForeignKeys {
		foreignKey.RefTable = nil
	}

	// Encode and write the table metadata to the file
	if err := encoder.Encode(new_tm); err != nil {
		return fmt.Errorf("failed to encode table metadata: %v", err)
	}

	fmt.Printf("Table metadata file '%s' created successfully\n", filePath)
	return nil
}
