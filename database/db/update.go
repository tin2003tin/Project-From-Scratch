package db

import (
	table "database/db/table"
	"encoding/gob"
	"fmt"
	"os"
)

func createTableMetadataFile(tbl *table.Table, filePath string) error {
	// Open or create the table metadata file
	metaFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create table metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	// Encode and write the table metadata to the file
	if err := encoder.Encode(tbl.Metadata); err != nil {
		return fmt.Errorf("failed to encode table metadata: %v", err)
	}

	fmt.Printf("Table metadata file '%s' created successfully\n", filePath)
	return nil
}

func createMetadataFile(db *Database, filePath string) error {
	// Open or create the metadata file
	metaFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	// Encode and write the database metadata to the file
	if err := encoder.Encode(db); err != nil {
		return fmt.Errorf("failed to encode metadata: %v", err)
	}

	fmt.Printf("Metadata file '%s' created successfully\n", filePath)
	return nil
}

func (db *Database) updateMetadataFile() error {
	// Open or create the metadata file
	metaFile, err := os.Create(db.MetadataPath + "/" + db.Name + ".meta")
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	// Encode and write the updated database metadata to the file
	if err := encoder.Encode(db); err != nil {
		return fmt.Errorf("failed to encode metadata: %v", err)
	}

	fmt.Println("Metadata file updated successfully")
	return nil
}
