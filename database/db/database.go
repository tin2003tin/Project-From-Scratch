package db

import (
	table "database/db/table"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Database struct {
	Name          string                // Database name
	Tables        []*table.Table        // Slice of tables
	Registry      *table.TableRegistry  // TableRegistry for managing tables by name
	MetadataPath  string                // Path to the metadata file
}



func CreateDatabase(name string) (*Database, error) {
	// Create the database instance
	db := &Database{
		Name:     name,
		Tables:   make([]*table.Table, 0),
		Registry: &table.TableRegistry{Tables: make(map[string]*table.Table)},
		MetadataPath: "./collection/" + name,
	}

	// Create a folder for the database
	dbFolderPath := filepath.Join("./collection/", name)
	if err := os.Mkdir(dbFolderPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create database folder: %v", err)
	}

	// Create and write metadata to the metadata file
	metaFilePath := filepath.Join(dbFolderPath, fmt.Sprintf("%s.meta", name))
	if err := createMetadataFile(db, metaFilePath); err != nil {
		return nil, fmt.Errorf("failed to create metadata file: %v", err)
	}

	return db, nil
}

func (db *Database) CreateTable(name string) (*table.Table, error) {
	// Validate input parameters
	if name == "" {
		return nil, errors.New("table name cannot be empty")
	}
	// Check if the table with the same name already exists
	if _, exists := db.Registry.Tables[name]; exists {
		return nil, fmt.Errorf("table '%s' already exists", name)
	}

	// Create default metadata
	metadata := table.TableMetadata{
		Name: name,
		Rows: make([]table.Row, 0),
		MetadataPath: db.MetadataPath,
	}
	// Create the table index table
	indexTable := make(map[string]*table.Index)
	// Create and return the Table instance
	newTable := &table.Table{
		Metadata:    metadata,
		IndexTable:  indexTable,
	}
	// Add the new table to the database
	db.Tables = append(db.Tables, newTable)
	// Register the new table in the TableRegistry
	db.Registry.Tables[name] = newTable

	// Update the metadata file
	tableMetaFilePath := filepath.Join(db.MetadataPath, fmt.Sprintf("%s.ttm", name))
	if err := createTableMetadataFile(newTable, tableMetaFilePath); err != nil {
		return nil, fmt.Errorf("failed to create table metadata file: %v", err)
	}

	return newTable, nil
}

func (db *Database) GetTable(name string) (*table.Table, error) {
	tbl, ok := db.Registry.Tables[name]
	if !ok {
		return nil, fmt.Errorf("table '%s' not found in the database", name)
	}
	return tbl, nil
}