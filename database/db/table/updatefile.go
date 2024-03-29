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
		return fmt.Errorf("failed to update table metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	refTables := make([]*Table, len(t.Metadata.ForeignKeys))
	for i := range t.Metadata.ForeignKeys {
		refTables[i] = t.Metadata.ForeignKeys[i].RefTable
		t.Metadata.ForeignKeys[i].RefTable = nil
	}

	// Encode and write the updated database metadata to the file
	if err := encoder.Encode(t.Metadata); err != nil {
		return fmt.Errorf("failed to encode table's metadata: %v", err)
	}

	for i := range t.Metadata.ForeignKeys {
		t.Metadata.ForeignKeys[i].RefTable = refTables[i]
	}

	fmt.Println("Table metadata file updated successfully")
	return nil
}

func (t *Table) SerializeRows() error {
	// Update the index file
	if err := t.updateIndexFile(); err != nil {
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

func (t *Table) updateIndexFile() error {
	indexFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".tti")

	// Open or create the index file
	indexFile, err := os.Create(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to update index file: %v", err)
	}
	defer indexFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(indexFile)
	var indexTable Index

	indexTable = Index{
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
