package db

import (
	table "database/db/table"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetDataBase(name string) (*Database, error) {
	// Construct the metadata file path
	metaFilePath := fmt.Sprintf("./collection/%s/%s.meta", name, name)

	// Open the metadata file for reading
	metaFile, err := os.Open(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open metadata file: %v", err)
	}
	defer metaFile.Close()

	decoder := gob.NewDecoder(metaFile)

	// Decode the database metadata from the file
	var db Database
	if err := decoder.Decode(&db); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %v", err)
	}

	// Read individual table metadata files and add them to the database
	if err := db.readTable(metaFilePath); err != nil {
		return nil, fmt.Errorf("failed to read table metadata: %v", err)
	}

	if err := db.buildIndex(); err != nil {
		return nil, fmt.Errorf("failed to build index: %v", err)
	}

	return &db, nil
}

func (db *Database) readTable(metaFilePath string) error {
	tablesPath := filepath.Join(filepath.Dir(metaFilePath))
	tableFiles, err := filepath.Glob(filepath.Join(tablesPath, "*.ttm"))
	if err != nil {
		return fmt.Errorf("failed to list table metadata files: %v", err)
	}

	for _, tableFile := range tableFiles {
		// Open the table metadata file for reading
		tableMetaFile, err := os.Open(tableFile)
		if err != nil {
			return fmt.Errorf("failed to open table metadata file '%s': %v", tableFile, err)
		}
		defer tableMetaFile.Close()

		// Initialize a decoder for reading binary data
		tableDecoder := gob.NewDecoder(tableMetaFile)

		// Decode the table metadata from the file
		var tableMeta table.TableMetadata
		if err := tableDecoder.Decode(&tableMeta); err != nil {
			return fmt.Errorf("failed to decode table metadata from '%s': %v", tableFile, err)
		}

		// Create a new table instance
		newTable := &table.Table{Metadata: tableMeta}

		// Read the corresponding .ttr file (serialized rows) if it exists
		rawDataFilePath := strings.TrimSuffix(tableFile, filepath.Ext(tableFile)) + ".ttr"
		if _, err := os.Stat(rawDataFilePath); err == nil {
			err := db.readRawData(newTable)
			if err != nil {
				return err
			}
		}

		// Read the corresponding index file (.tti) if it exists
		indexFilePath := strings.TrimSuffix(tableFile, filepath.Ext(tableFile)) + ".tti"
		if _, err := os.Stat(indexFilePath); err == nil {
			indexTable, err := db.readIndexFile(indexFilePath)
			if err != nil {
				return err
			}
			newTable.IndexTable = indexTable
		}

		db.Tables = append(db.Tables, newTable)
		db.Registry.Tables[tableMeta.Name] = newTable
	}
	if err := db.buildForeignKey(db.Tables); err != nil {
		return err
	}

	return nil
}

func (db *Database) readRawData(t *table.Table) error {
	// Construct the path for the .ttr file
	ttrFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".ttr")

	// Open the .ttr file for reading
	ttrFile, err := os.Open(ttrFilePath)
	if err != nil {
		return fmt.Errorf("failed to open .ttr file: %v", err)
	}
	defer ttrFile.Close()

	// Initialize a decoder for reading binary data
	decoder := gob.NewDecoder(ttrFile)

	// Decode the rows from the .ttr file
	var rows []table.Row
	if err := decoder.Decode(&rows); err != nil {
		return fmt.Errorf("failed to decode rows from .ttr file: %v", err)
	}

	// Update the table's rows with the decoded rows
	t.Metadata.Rows = rows

	fmt.Printf("Rows read from '%s' successfully\n", ttrFilePath)
	return nil
}

func (db *Database) readIndexFile(indexFilePath string) (*table.Index, error) {
	// Open the index file for reading
	indexFile, err := os.Open(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open index file '%s': %v", indexFilePath, err)
	}
	defer indexFile.Close()

	// Initialize a decoder for reading binary data
	indexDecoder := gob.NewDecoder(indexFile)

	var indexTable *table.Index
	if err := indexDecoder.Decode(&indexTable); err != nil {
		return nil, fmt.Errorf("failed to decode index data from '%s': %v", indexFilePath, err)
	}

	return indexTable, nil
}

// buildIndex builds indexes for all tables in the database based on the specified keys.
func (db *Database) buildIndex() error {
	for _, t := range db.Tables {
		t.IndexTable.Rows = make(map[string]*table.Row)
		t.IndexTable.Columns = make(map[string]*table.Column)
		for _, column := range t.Metadata.Columns {
			for _, ci := range t.IndexTable.Include {
				if ci == column.Name {
					t.IndexTable.Columns[ci] = &column
					for _, row := range t.Metadata.Rows {
						data, ok := row.Data[ci]
						if !ok {
							return fmt.Errorf("key '%s' not found in row data", ci)
						}
						indexKey := make(map[string]interface{})
						indexKey[column.Name] = data
						t.IndexTable.Rows[fmt.Sprintf("%v", indexKey)] = &row
					}
				}
			}
		}
	}
	return nil
}

func (db *Database) buildForeignKey(tables []*table.Table) error {
	for _, t := range tables {
		for i := range t.Metadata.ForeignKeys {
			t.Metadata.ForeignKeys[i].RefTable = db.Registry.Tables[t.Metadata.ForeignKeys[i].RefTableName]
		}
	}

	return nil
}
