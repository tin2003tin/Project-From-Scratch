package db

import (
	table "database/db/table"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

type Database struct {
	Name    string                  // Database name
	Tables  []*table.Table          // Slice of tables
	Registry *table.TableRegistry   // TableRegistry for managing tables by name
}

func CreateDatabase(name string) *Database {
	return &Database{
		Name:     name,
		Tables:   make([]*table.Table, 0),
		Registry: &table.TableRegistry{Tables: make(map[string]*table.Table)},
	}
}

func (db *Database) CreateTable(name string) (*table.Table, error) {
	// Validate input parameters
	if name == "" {
		return nil, errors.New("table name cannot be empty")
	}
	// Create default metadata
	metadata := table.TableMetadata{
		Name: name,
		Rows: make([]table.Row, 0),
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
	return newTable, nil
}

func (db *Database) GetTable(name string) (*table.Table, error) {
	tbl, ok := db.Registry.Tables[name]
	if !ok {
		return nil, fmt.Errorf("table '%s' not found in the database", name)
	}
	return tbl, nil
}

func (db *Database) SaveDatabase(filename string) error {
	file, err := os.Create(filename+".bin")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(db); err != nil {
		return err
	}

	fmt.Println("Database saved to", filename)
	return nil
}

func LoadDatabase(filename string) (*Database, error) {
	var db Database

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		return nil, err
	}

	fmt.Println("Database loaded from", filename)
	return &db, nil
}