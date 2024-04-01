package filemanager

import (
	db_constant "database/db/constant"
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func CreateDatabaseCollection(db *structure.Database) error {
	dbFolderPath := filepath.Join(db_constant.DATABASE_PATH, db.Name)
	if err := os.Mkdir(dbFolderPath, 0755); err != nil {
		return fmt.Errorf("failed to create database folder: %v", err)
	}
	if err := createDatabaseMetadataFile(db); err != nil {
		return err
	}
	return nil
}

func createDatabaseMetadataFile(db *structure.Database) error {
	// Open or create the metadata file
	filePath := filepath.Join(db.MetadataPath, fmt.Sprintf("%s.meta", db.Name))
	metaFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer metaFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)
	new_db := structure.Database{
		Name:         db.Name,
		TableNames:   db.TableNames,
		MetadataPath: db.MetadataPath,
		Tables:       []*structure.Table{},
		Registry:     &structure.TableRegistry{make(map[string]*structure.Table)},
	}
	// Encode and write the database metadata to the file
	if err := encoder.Encode(new_db); err != nil {
		return fmt.Errorf("failed to encode metadata: %v", err)
	}

	fmt.Printf("Metadata file '%s' created successfully\n", filePath)
	return nil
}

func UpdateDatabaseMetadataFile(db *structure.Database) error {
	// Open or create the metadata file
	metaFile, err := os.Create(db.MetadataPath + "/" + db.Name + ".meta")
	if err != nil {
		return fmt.Errorf("failed to create metadata file: %v", err)
	}
	defer metaFile.Close()
	new_db := structure.Database{
		Name:         db.Name,
		TableNames:   db.TableNames,
		MetadataPath: db.MetadataPath,
		Tables:       []*structure.Table{},
		Registry:     &structure.TableRegistry{make(map[string]*structure.Table)},
		Username:     db.Username,
		Password:     db.Password,
	}
	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(metaFile)

	// Encode and write the updated database metadata to the file
	if err := encoder.Encode(new_db); err != nil {
		return fmt.Errorf("failed to encode metadata: %v", err)
	}

	fmt.Println("Metadata file updated successfully")
	return nil
}
