package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type DatabaseMetadata struct {
	Name         string    // Database name
	CreatedAt    time.Time // Timestamp of database creation
	LastModified time.Time // Timestamp of last modification
	Version      string    // Database version
	Comment      string    // Database comment or description
}

func main() {
	// Example DatabaseMetadata
	metadata := DatabaseMetadata{
		Name:         "MyDatabase",
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
		Version:      "1.0",
		Comment:      "Example database",
	}

	// Write metadata to .meta file
	if err := WriteDatabaseMetadata(metadata, "database.meta"); err != nil {
		fmt.Println("Error writing metadata to file:", err)
		return
	}
	fmt.Println("Metadata written to database.meta")

	// Read metadata from .meta file
	readMetadata, err := ReadDatabaseMetadata("database.meta")
	if err != nil {
		fmt.Println("Error reading metadata from file:", err)
		return
	}
	fmt.Println("Read Metadata:", readMetadata)
}

// WriteDatabaseMetadata writes the DatabaseMetadata to a .meta file
func WriteDatabaseMetadata(metadata DatabaseMetadata, filename string) error {
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
	if err := binary.Write(file, binary.LittleEndian, metadata.CreatedAt.UnixNano()); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, metadata.LastModified.UnixNano()); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, uint32(len(metadata.Version))); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, []byte(metadata.Version)); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, uint32(len(metadata.Comment))); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, []byte(metadata.Comment)); err != nil {
		return err
	}

	return nil
}

// ReadDatabaseMetadata reads DatabaseMetadata from a .meta file
func ReadDatabaseMetadata(filename string) (DatabaseMetadata, error) {
	var metadata DatabaseMetadata

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

	var createdAtNano int64
	if err := binary.Read(file, binary.LittleEndian, &createdAtNano); err != nil {
		return metadata, err
	}
	metadata.CreatedAt = time.Unix(0, createdAtNano)

	var lastModifiedNano int64
	if err := binary.Read(file, binary.LittleEndian, &lastModifiedNano); err != nil {
		return metadata, err
	}
	metadata.LastModified = time.Unix(0, lastModifiedNano)

	var versionLen uint32
	if err := binary.Read(file, binary.LittleEndian, &versionLen); err != nil {
		return metadata, err
	}
	versionBytes := make([]byte, versionLen)
	if err := binary.Read(file, binary.LittleEndian, &versionBytes); err != nil {
		return metadata, err
	}
	metadata.Version = string(versionBytes)

	var commentLen uint32
	if err := binary.Read(file, binary.LittleEndian, &commentLen); err != nil {
		return metadata, err
	}
	commentBytes := make([]byte, commentLen)
	if err := binary.Read(file, binary.LittleEndian, &commentBytes); err != nil {
		return metadata, err
	}
	metadata.Comment = string(commentBytes)

	return metadata, nil
}