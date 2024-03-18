package table

import (
	"errors"
	"fmt"
	"time"
)

func (t *Table) AddIdColumn() error {
	err := t.AddColumn("id", "int", 0, 0, 0, true, true, false, nil, "", "")
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) AddColumn(name string, dataType string, length int, precision int, scale int, primaryKey bool, unique bool, nullable bool, defaultValue interface{}, check string, comment string) error {
	// Validate input parameters
	if name == "" {
		return errors.New("column name cannot be empty")
	}
	for _, col := range t.Metadata.Columns {
		if col.Name == name {
			return fmt.Errorf("column '%s' already exists in the table", name)
		}
	}
	// Create the new column
	newColumn := Column{
		Name:       name,
		DataType:   dataType,
		Length:     length,
		Precision:  precision,
		Scale:      scale,
		PrimaryKey: primaryKey,
		ForeignKey: false,
		Unique:     unique,
		Nullable:   nullable,
		Default:    defaultValue,
		Check:      check,
		Comment:    comment,
	}
	// Add the new column to the table's metadata
	t.Metadata.Columns = append(t.Metadata.Columns, newColumn)
	if newColumn.PrimaryKey {
		t.addPrimaryKey(newColumn.Name)
	}
	if newColumn.PrimaryKey || newColumn.Unique {
		err := t.AddIndex(newColumn.Name)
		if err != nil {
			return err
		}
	}

	// Update the metadata file after adding the column
	if err := t.updateMetadataFile(); err != nil {
		return fmt.Errorf("failed to update metadata file: %v", err)
	}

	return nil
}

func (t *Table) addPrimaryKey(name string) {
	t.Metadata.PrimaryKeys = append(t.Metadata.PrimaryKeys, name)
}

func (t *Table) AddRow(columnValues map[string]interface{}) error {
	// Check if the target table is nil
	if t == nil {
		return errors.New("target table is nil")
	}
	createdAt := time.Now()
	newRow := Row{
		Data:      columnValues,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
	// Validate primary key and non-nullable columns
	for _, column := range t.Metadata.Columns {
		if column.PrimaryKey && columnValues[column.Name] == nil {
			return errors.New("primary key column not provided: " + column.Name)
		}
		if !column.Nullable && columnValues[column.Name] == nil {
			return errors.New("non-nullable column not provided: " + column.Name)
		}

		if column.ForeignKey {
			if !t.dataInForeignKeyExisted(&newRow, column.Name) {
				return fmt.Errorf("%v not found in foreign key column '%s'", newRow.Data[column.Name], column.Name)
			}
		}
	}

	// Add the new row to the table's Rows slice
	t.Metadata.Rows = append(t.Metadata.Rows, newRow)

	// Update the index with the new values
	t.IndexTable.AddRowToIndex(newRow)
	return nil
}

func (i *Index) AddRowToIndex(row Row) error {
	// Extract relevant values from the new row based on the indexed columns
	for _, column := range i.Columns {
		if value, ok := row.Data[column.Name]; ok {
			indexKey := make(map[string]interface{})
			indexKey[column.Name] = value
			i.Rows[fmt.Sprintf("%v", indexKey)] = &row
		}
	}
	return nil
}

func (t *Table) AddIndex(columnName string) error {
	// Create the index object]
	if t.IndexTable == nil {
		t.IndexTable = &Index{
			Name:       "default_index",
			Columns:    make(map[string]*Column),
			Rows:       make(map[string]*Row),
			Unique:     true,
			Comment:    "This is a default index",
			Include:    make([]string, 0),
			Predicate:  "",
			FillFactor: 0,
		}
	}
	// Populate the index columns
	found := false
	for _, column := range t.Metadata.Columns {
		if column.Name == columnName {
			t.IndexTable.Columns[columnName] = &column
			t.IndexTable.Include = append(t.IndexTable.Include, columnName)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("column '%v' not found in the table", columnName)
	}
	// Update the index file (.tti) and write the index data
	if err := t.updateIndexFile(); err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}

	return nil
}