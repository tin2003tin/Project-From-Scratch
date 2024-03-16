package table

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func (t *Table) updateMetadataFile() error {
	// Construct the metadata file path
	metaFilePath := fmt.Sprintf("%s/%s.ttm", t.Metadata.MetadataPath, t.Metadata.Name)

	// Open or create the metadata file
	metaFile, err := os.Create(metaFilePath)
	if err != nil {
		return fmt.Errorf("failed to create table's metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	// Encode and write the updated database metadata to the file
	if err := encoder.Encode(t.Metadata); err != nil {
		return fmt.Errorf("failed to encode table's metadata: %v", err)
	}

	fmt.Println("Metadata file updated successfully")
	return nil
}

func createIndexFile(table *Table, metadataPath string) error {
	indexFilePath := filepath.Join(metadataPath, table.Metadata.Name+".tti")

	// Open or create the index file
	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}
	defer indexFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(indexFile)

	// Encode and write the index data to the file
	if err := encoder.Encode(table.IndexTable); err != nil {
		return fmt.Errorf("failed to encode index data: %v", err)
	}

	fmt.Println("Index file created successfully:", indexFilePath)
	return nil
}

func (t *Table) SerializeRows() error {
	// Create the index file
	if err := createIndexFile(t, t.Metadata.MetadataPath); err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}

	// Construct the path for the .ttr file
	ttrFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".ttr")

	// Create or open the .ttr file
	ttrFile, err := os.Create(ttrFilePath)
	if err != nil {
		return fmt.Errorf("failed to create .ttr file: %v", err)
	}
	defer ttrFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(ttrFile)

	// Encode and write the rows to the .ttr file
	if err := encoder.Encode(t.Metadata.Rows); err != nil {
		return fmt.Errorf("failed to encode rows: %v", err)
	}

	fmt.Printf("Rows serialized to '%s' successfully\n", ttrFilePath)
	return nil
}