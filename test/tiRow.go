package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type TableRow struct {
	ID     uint32 // ID of the record
	Offset int64  // Byte offset of the record in the data file
}

func createIndex(dataFile *os.File, indexFile *os.File) error {
	// Seek to the beginning of the data file
	_, err := dataFile.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	// Read records from the data file and create the index
	var id uint32
	var offset int64
	index := make(map[uint32]int64)

	for {
		// Read the ID and offset of the record
		err := binary.Read(dataFile, binary.LittleEndian, &id)
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file reached
			}
			return err
		}

		// Get the current offset in the data file
		offset, err = dataFile.Seek(0, os.SEEK_CUR)
		if err != nil {
			return err
		}

		// Store the ID and offset in the index map
		index[id] = offset - 4 // Subtract 4 bytes for the ID field itself
	}

	// Write the index entries to the index file
	for id, offset := range index {
		indexEntry := TableRow{ID: id, Offset: offset}
		if err := binary.Write(indexFile, binary.LittleEndian, indexEntry); err != nil {
			return err
		}
	}

	return nil
}

func readRecordByID(dataFile *os.File, indexFile *os.File, id uint32) ([]byte, error) {
	// Read the index entry from the index file
	var indexEntry TableRow
	err := binary.Read(indexFile, binary.LittleEndian, &indexEntry)
	if err != nil {
		return nil, err
	}

	// Seek to the offset of the record in the data file
	_, err = dataFile.Seek(indexEntry.Offset, os.SEEK_SET)
	if err != nil {
		return nil, err
	}

	// Read the record data from the data file
	record := make([]byte, 100) // Assuming fixed-size record length of 100 bytes
	_, err = dataFile.Read(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func main() {
	// Open data file and index file
	dataFile, err := os.OpenFile("table.td", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening data file:", err)
		return
	}
	defer dataFile.Close()

	indexFile, err := os.OpenFile("index.ti", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening index file:", err)
		return
	}
	defer indexFile.Close()

	// Create the index
	if err := createIndex(dataFile, indexFile); err != nil {
		fmt.Println("Error creating index:", err)
		return
	}
	fmt.Println("Index created successfully.")

	// Example: Read record by ID
	id := uint32(10)
	recordData, err := readRecordByID(dataFile, indexFile, id)
	if err != nil {
		fmt.Println("Error reading record:", err)
		return
	}
	fmt.Printf("Record with ID %d: %s\n", id, string(recordData))
}
