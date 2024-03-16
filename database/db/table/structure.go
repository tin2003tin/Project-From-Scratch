package table

import "time"

// TableRegistry stores tables by name for easy lookup
type TableRegistry struct {
	Tables map[string]*Table
}

// Table represents a database table
type Table struct {
	Metadata    TableMetadata    // Table metadata
	IndexTable  map[string]*Index // Indexes on the table
}

// TableMetadata represents metadata about a table
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

// Column represents a column in a table
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

// Row represents a row of data in a table
type Row struct {
	Data       map[string]interface{} // Map to store column name-value pairs
	CreatedAt  time.Time               // Timestamp of row creation
	UpdatedAt  time.Time               // Timestamp of last update
}

// ForeignKey represents a foreign key constraint
type ForeignKey struct {
	Name         string     // Constraint name
	Columns      []string   // Columns in the current table
	RefTable     string     // Referenced table
	RefColumns   []string   // Referenced columns in the referenced table
	OnUpdate     string     // Action to perform on update (e.g., CASCADE, SET NULL)
	OnDelete     string     // Action to perform on delete (e.g., CASCADE, SET NULL)
}

// Index represents an index on a table
type Index struct {
	Name       string             // Index name
	Columns    map[int]*Column
	Rows       map[string]*Row    // Columns included in the index with additional properties
	Unique     bool               // Indicates if the index enforces uniqueness
	Using      IndexType          // Index method
	Comment    string             // Index comment or description
	Tablespace string             // Tablespace where the index is stored
	Include    []string           // Included columns (for covering indexes)
	Predicate  string             // Index predicate expression (for partial indexes)
	FillFactor int                // Fill factor (for certain index types)
}

// IndexType represents the type of index
type IndexType int

const (
	BTreeIndex IndexType = iota
	HashIndex
	GINIndex
)