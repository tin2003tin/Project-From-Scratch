package db

import (
	table "database/db/table"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// createMetadataFile creates a metadata file and writes database metadata to it
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

func (db *Database) UpdateMetadataFile() error {
	// Open or create the metadata file
	fmt.Println(db.MetadataPath)
	metaFile, err := os.Create(db.MetadataPath + "/" + db.Name + ".meta" )
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

func GetDataBase(name string) (*Database, error) {
	// Construct the metadata file path
	metaFilePath := fmt.Sprintf("./collection/%s/%s.meta", name, name)

	// Open the metadata file for reading
	metaFile, err := os.Open(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize a decoder for reading binary data
	decoder := gob.NewDecoder(metaFile)

	// Decode the database metadata from the file
	var db Database
	if err := decoder.Decode(&db); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %v", err)
	}

	// Read individual table metadata files and add them to the database
	if err := readTable(&db, metaFilePath); err != nil {
		return nil, fmt.Errorf("failed to read table metadata: %v", err)
	}

	return &db, nil
}

func readIndexFile(indexFilePath string) (map[string]*table.Index, error) {
	// Open the index file for reading
	indexFile, err := os.Open(indexFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open index file '%s': %v", indexFilePath, err)
	}
	defer indexFile.Close()

	// Initialize a decoder for reading binary data
	indexDecoder := gob.NewDecoder(indexFile)

	// Decode the index data from the file
	var indexTable map[string]*table.Index
	if err := indexDecoder.Decode(&indexTable); err != nil {
		return nil, fmt.Errorf("failed to decode index data from '%s': %v", indexFilePath, err)
	}

	return indexTable, nil
}

func readTTRFile(t *table.Table) error {
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


func readTable(db *Database, metaFilePath string) error {
	tablesPath := filepath.Join(filepath.Dir(metaFilePath))
	tableFiles, err := filepath.Glob(filepath.Join(tablesPath, "*.ttm"))
	if err != nil {
		return fmt.Errorf("failed to list table metadata files: %v", err)
	}

	// Iterate over table metadata files and decode them
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
		newTable := &table.Table{Metadata: tableMeta, IndexTable: make(map[string]*table.Index)}

		// Read the corresponding index file (.tti) if it exists
		indexFilePath := strings.TrimSuffix(tableFile, filepath.Ext(tableFile)) + ".tti"
		if _, err := os.Stat(indexFilePath); err == nil {
			indexTable, err := readIndexFile(indexFilePath)
			if err != nil {
				return err
			}
			newTable.IndexTable = indexTable
		}

		// Read the corresponding .ttr file (serialized rows)
		if err := readTTRFile(newTable); err != nil {
			return err
		}

		db.Tables = append(db.Tables, newTable)
		db.Registry.Tables[tableMeta.Name] = newTable
	}

	return nil
}

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