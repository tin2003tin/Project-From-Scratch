package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type TableMetadata struct {
	Name         string         // Table name
	Columns      []Column       // List of columns
	Rows         []Row          // List of rows in the table
	PrimaryKeys  []string       // List of primary key column names
	ForeignKeys  []ForeignKey   // List of foreign key constraints
	Check        string         // Table-level check constraint expression
	Comment      string         // Table comment or description
	Tablespace   string         // Tablespace where the table is stored
	Inheritance  []string       // List of inherited tables (for table inheritance)
	Owner        string         // Owner or creator of the table
	LastModified time.Time      // Last modification timestamp
}

type Column struct {
	Name         string      // Column name
	DataType     string      // Data type (e.g., int, string)
	Length       int         // Length of the column (for variable-length types)
	Precision    int         // Precision of the column (for numeric types)
	Scale        int         // Scale of the column (for numeric types)
	PrimaryKey   bool        // Indicates if the column is part of the primary key
	ForeignKey   bool        // Indicates if the column is a foreign key
	Unique       bool        // Indicates if the column values must be unique
	Nullable     bool        // Indicates if the column allows NULL values
	Default      interface{} // Default value for the column
	Check        string      // Check constraint expression
	Comment      string      // Column comment or description
}

type Row struct {
	Data       map[string]interface{} // Map to store column name-value pairs
	CreatedAt  time.Time               // Timestamp of row creation
	UpdatedAt  time.Time               // Timestamp of last update
}

type ForeignKey struct {
	Name         string     // Constraint name
	Columns      []string   // Columns in the current table
	RefTable     string     // Referenced table
	RefColumns   []string   // Referenced columns in the referenced table
	OnUpdate     string     // Action to perform on update (e.g., CASCADE, SET NULL)
	OnDelete     string     // Action to perform on delete (e.g., CASCADE, SET NULL)
}

func main() {
	// Example TableMetadata
	metadata := TableMetadata{
		Name:         "MyTable",
		Columns:      []Column{{Name: "ID", DataType: "int", PrimaryKey: true}},
		Rows:         []Row{{Data: map[string]interface{}{"ID": 1}, CreatedAt: time.Now()}},
		PrimaryKeys:  []string{"ID"},
		ForeignKeys:  nil,
		Check:        "",
		Comment:      "Example table",
		Tablespace:   "",
		Inheritance:  nil,
		Owner:        "Admin",
		LastModified: time.Now(),
	}

	// Write metadata to .tm file
	if err := WriteTableMetadata(metadata, "table.tm"); err != nil {
		fmt.Println("Error writing metadata to file:", err)
		return
	}
	fmt.Println("Metadata written to table.tm")

	// Read metadata from .tm file
	readMetadata, err := ReadTableMetadata("table.tm")
	if err != nil {
		fmt.Println("Error reading metadata from file:", err)
		return
	}
	fmt.Println("Read Metadata:", readMetadata)
}

// WriteTableMetadata writes the TableMetadata to a .tm file
func WriteTableMetadata(metadata TableMetadata, filename string) error {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write metadata fields to the file
	if err := binary.Write(file, binary.LittleEndian, uint32(len(metadata.Name))); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, []byte(metadata.Name)); err != nil {
		return err
	}
	// Write other metadata fields similarly...

	return nil
}

// ReadTableMetadata reads TableMetadata from a .tm file
func ReadTableMetadata(filename string) (TableMetadata, error) {
	var metadata TableMetadata

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return metadata, err
	}
	defer file.Close()

	// Read metadata fields from the file
	var nameLen uint32
	if err := binary.Read(file, binary.LittleEndian, &nameLen); err != nil {
		return metadata, err
	}
	nameBytes := make([]byte, nameLen)
	if err := binary.Read(file, binary.LittleEndian, &nameBytes); err != nil {
		return metadata, err
	}
	metadata.Name = string(nameBytes)

	// Read other metadata fields similarly...

	return metadata, nil
}
