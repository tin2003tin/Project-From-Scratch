package queryProcessor

import (
	"database/db/lib"
	filemanager "database/db/storageManager/fileManager"
	"database/db/structure"
	"errors"
	"fmt"
)

func (q *QueryManager) AddIdColumn() error {
	err := q.AddColumn("id", "int", 0, true, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (q *QueryManager) AddColumn(name string, dataType string, length int, primaryKey bool, unique bool, nullable bool, defaultValue interface{}) error {
	// Validate input parameters
	if name == "" {
		return errors.New("column name cannot be empty")
	}
	for _, col := range q.Table.Columns {
		if col.Name == name {
			return fmt.Errorf("column '%s' already exists in the table", name)
		}
	}
	// Create the new column
	newColumn := structure.Column{
		Index:      len(q.Table.Columns),
		Name:       name,
		DataType:   dataType,
		Length:     length,
		Precision:  0,
		Scale:      0,
		PrimaryKey: primaryKey,
		ForeignKey: false,
		Unique:     unique,
		Nullable:   nullable,
		Default:    defaultValue,
		Check:      "",
		Comment:    "",
	}
	// Check type of defualt
	if newColumn.Default != nil {
		newDefualt, err := lib.ConvertValue(newColumn.Default, newColumn.DataType, newColumn.Length)
		if err != nil {
			return fmt.Errorf("cannot add column, %v %v", newColumn.Name, err)
		}
		newColumn.Default = newDefualt
	}

	// Add the new column to the table
	q.Table.Columns = append(q.Table.Columns, newColumn)
	if newColumn.PrimaryKey {
		q.addPrimaryKey(newColumn.Name)
	}

	// Create index if column is primary key or unique
	if newColumn.PrimaryKey || newColumn.Unique {
		err := q.AddIndex(newColumn.Name)
		if err != nil {
			return err
		}
	}

	// Update the metadata file after adding the column
	if err := filemanager.UpdateTableFile(q.Table); err != nil {
		return fmt.Errorf("failed to update metadata file: %v", err)
	}

	return nil
}

func (q *QueryManager) addPrimaryKey(name string) {
	q.Table.Metadata.PrimaryKeys = append(q.Table.Metadata.PrimaryKeys, name)
}

func (q *QueryManager) hasColumn(columnName string) bool {
	for _, col := range q.Table.Columns {
		if col.Name == columnName {
			return true
		}
	}
	return false
}

func (q *QueryManager) PrintColumn() {
	for i, col := range q.Table.Columns {
		fmt.Println(i, col.Name, col.DataType)
	}
}
