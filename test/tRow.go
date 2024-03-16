package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type Row struct {
	ID        uint32    // Example field
	Name      string    // Example field
	Age       uint8     // Example field
	CreatedAt int64     // Timestamp of row creation (UnixNano)
	UpdatedAt int64     // Timestamp of last update (UnixNano)
}

func main() {
	// Example rows data
	rows := []Row{
		{
			ID:        1,
			Name:      "John Doe",
			Age:       30,
			CreatedAt: time.Now().UnixNano(),
			UpdatedAt: time.Now().UnixNano(),
		},
		{
			ID:        2,
			Name:      "Jane Smith",
			Age:       35,
			CreatedAt: time.Now().UnixNano(),
			UpdatedAt: time.Now().UnixNano(),
		},
	}

	// Write rows data to .td file
	if err := WriteTableRows(rows, "table.td"); err != nil {
		fmt.Println("Error writing rows data to file:", err)
		return
	}
	fmt.Println("Rows data written to table.td")

	// Read rows data from .td file
	readRows, err := ReadTableRows("table.td")
	if err != nil {
		fmt.Println("Error reading rows data from file:", err)
		return
	}
	fmt.Println("Read Rows:", readRows)
}

// WriteTableRows writes the array of rows data to a .td file in binary format (one line per row)
func WriteTableRows(rows []Row, filename string) error {
	// Open the file for writing, create if not exists, truncate if exists
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Iterate through rows and write each row in binary format (serialized byte slice)
	for _, row := range rows {
		serializedRow := SerializeRow(row)
		if _, err := file.Write(serializedRow); err != nil {
			return err
		}
	}

	return nil
}

// SerializeRow serializes a single row into a byte slice
func SerializeRow(row Row) []byte {
	serialized := make([]byte, 0)

	// Serialize each field and append to the serialized byte slice
	serialized = append(serialized, SerializeUint32(row.ID)...)
	serialized = append(serialized, SerializeString(row.Name)...)
	serialized = append(serialized, SerializeUint8(row.Age)...)
	serialized = append(serialized, SerializeInt64(row.CreatedAt)...)
	serialized = append(serialized, SerializeInt64(row.UpdatedAt)...)

	return serialized
}

// Helper functions to serialize different data types

func SerializeUint32(value uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, value)
	return buf
}

func SerializeUint8(value uint8) []byte {
	return []byte{value}
}

func SerializeInt64(value int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(value))
	return buf
}

func SerializeString(value string) []byte {
	return append([]byte(value), 0) // Append null terminator for string
}

// ReadTableRows reads array of rows data from a .td file in binary format (one line per row)
func ReadTableRows(filename string) ([]Row, error) {
	var rows []Row

	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()

	// Read each line (row) and decode into Row struct
	for {
		row, err := ReadNextRow(file)
		if err != nil {
			break // Reached EOF or error
		}
		rows = append(rows, row)
	}

	return rows, nil
}

// ReadNextRow reads the next line (row) from the binary file and decodes it into a Row struct
func ReadNextRow(file *os.File) (Row, error) {
	var row Row

	// Read and decode each field from the serialized byte slice
	if err := binary.Read(file, binary.LittleEndian, &row.ID); err != nil {
		return row, err
	}
	nameBytes, err := ReadNullTerminatedString(file)
	if err != nil {
		return row, err
	}
	row.Name = string(nameBytes)
	if err := binary.Read(file, binary.LittleEndian, &row.Age); err != nil {
		return row, err
	}
	if err := binary.Read(file, binary.LittleEndian, &row.CreatedAt); err != nil {
		return row, err
	}
	if err := binary.Read(file, binary.LittleEndian, &row.UpdatedAt); err != nil {
		return row, err
	}

	return row, nil
}

// ReadNullTerminatedString reads a null-terminated string from the file
func ReadNullTerminatedString(file *os.File) ([]byte, error) {
	var buf []byte
	var b [1]byte
	for {
		if _, err := file.Read(b[:]); err != nil {
			return buf, err
		}
		if b[0] == 0 {
			break
		}
		buf = append(buf, b[0])
	}
	return buf, nil
}
