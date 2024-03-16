package table

import (
	"errors"
	"fmt"
	"time"
)

func (t *Table) AddIdColumn() error {
	err := t.AddColumn("id", "int", 0, 0, 0, true, false, true, false, nil, "", "")
	if err != nil {
		return err
	}
	return nil
}

func (t *Table) AddColumn(name string, dataType string, length int, precision int, scale int, primaryKey bool, foreignKey bool, unique bool, nullable bool, defaultValue interface{}, check string, comment string) error {
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
		Name:         name,
		DataType:     dataType,
		Length:       length,
		Precision:    precision,
		Scale:        scale,
		PrimaryKey:   primaryKey,
		ForeignKey:   foreignKey,
		Unique:       unique,
		Nullable:     nullable,
		Default:      defaultValue,
		Check:        check,
		Comment:      comment,
	}
	// Add the new column to the table's metadata
	t.Metadata.Columns = append(t.Metadata.Columns, newColumn)
	if (newColumn.PrimaryKey) {
		t.addPrimaryKey(newColumn.Name)
	}
	if (newColumn.PrimaryKey || newColumn.Unique) {
		err := t.AddIndex(newColumn.Name, []string{newColumn.Name}, true, HashIndex, "Index for the " + newColumn.Name, "default", nil, "", 70)
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
	createdAt := time.Now();
	for _, column := range t.Metadata.Columns {
		if column.PrimaryKey && columnValues[column.Name] == nil {
			return errors.New("primary key column not provided: " + column.Name)
		}
		if !column.Nullable && columnValues[column.Name] == nil {
			return errors.New("non-nullable column not provided: " + column.Name)
		}
	}

	newRow := Row{
		Data:      columnValues,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	// Add the new row to the table's Rows slice
	t.Metadata.Rows = append(t.Metadata.Rows, newRow)

	// Update the index with the new values
	for _, index := range t.IndexTable {
		index.AddRowToIndex(newRow)
	}
	return nil
}

func (i *Index) AddRowToIndex(row Row) error {
	// Extract relevant values from the new row based on the indexed columns
	indexKey := make(map[int]interface{})
	for columnIndex, column := range i.Columns {
		if value, ok := row.Data[column.Name]; ok {
			indexKey[columnIndex] = value
		}
	}
	// Add the row to the index using the extracted key
	if len(indexKey) > 0 {
		i.Rows[fmt.Sprintf("%v", indexKey)] = &row
		return nil
	}

	return errors.New("no indexable columns found in the row")
}

func (t *Table) AddIndex(name string, columns []string, unique bool, indexType IndexType, comment string, tablespace string, include []string, predicate string, fillFactor int) error {
	// Create the index object
	index := Index{
		Name:       name,
		Columns:    make(map[int]*Column),
		Rows:       make(map[string]*Row),
		Unique:     unique,
		Using:      indexType,
		Comment:    comment,
		Tablespace: tablespace,
		Include:    include,
		Predicate:  predicate,
		FillFactor: fillFactor,
	}

	// Populate the index columns
	for idx, colName := range columns {
		found := false
		for _, col := range t.Metadata.Columns {
			if col.Name == colName {
				index.Columns[idx] = &col
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("column '%s' not found in the table", colName)
		}
	}
	// Add the index to the table's index table
	t.IndexTable[index.Name] = &index
	// Create the index file (.tti) and write the index data
	if err := createIndexFile(t, t.Metadata.MetadataPath); err != nil {
		return fmt.Errorf("failed to create index file: %v", err)
	}

	return nil
}

func (t *Table) AddForeignKey(fk ForeignKey) {
	t.Metadata.ForeignKeys = append(t.Metadata.ForeignKeys, fk)
}
