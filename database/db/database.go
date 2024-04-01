package db

import (
	db_constant "database/db/constant"
	filemanager "database/db/storageManager/fileManager"
	"database/db/structure"
	"errors"
	"fmt"
)

func CreateDatabase(name string, username string, password string) (*structure.Database, error) {
	// Create the database instance
	db := &structure.Database{
		Name:         name,
		Tables:       make([]*structure.Table, 0),
		Registry:     &structure.TableRegistry{Tables: make(map[string]*structure.Table)},
		MetadataPath: db_constant.DATABASE_PATH + name,
		TableNames:   []string{},
		Username:     username,
		Password:     password,
	}

	// Create the data file
	if err := filemanager.CreateDatabaseCollection(db); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *structure.Database, name string) (*structure.Table, error) {
	if name == "" {
		return nil, errors.New("table name cannot be empty")
	}
	// Check if the table with the same name already exists
	if _, exists := db.Registry.Tables[name]; exists {
		return nil, fmt.Errorf("table '%s' already exists", name)
	}

	// Create default metadata
	metadata := structure.TableMetadata{
		Name:         name,
		MetadataPath: db.MetadataPath,
	}
	// Create and return the Table instance
	newTable := &structure.Table{
		Metadata:   metadata,
		IndexTable: nil,
	}
	// Add the new table to the database
	db.Tables = append(db.Tables, newTable)

	// Register the new table in the TableRegistry
	db.Registry.Tables[name] = newTable

	db.TableNames = append(db.TableNames, newTable.Metadata.Name)

	// Create the metadata file
	if err := filemanager.CreateTableFile(newTable); err != nil {
		return nil, fmt.Errorf("failed to create table metadata file: %v", err)
	}
	if err := filemanager.UpdateDatabaseMetadataFile(db); err != nil {
		return nil, err
	}

	return newTable, nil
}

func GetTable(db *structure.Database, name string) (*structure.Table, error) {
	tbl, ok := db.Registry.Tables[name]
	if !ok {
		return nil, fmt.Errorf("table '%s' not found in the database", name)
	}
	return tbl, nil
}
