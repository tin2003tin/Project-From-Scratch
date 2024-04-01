package structure

type Database struct {
	Name         string         // Database name
	Tables       []*Table       // Slice of tables
	Registry     *TableRegistry // TableRegistry for managing tables by name
	TableNames   []string
	Username     string
	Password     string
	MetadataPath string // Path to the metadata file
}
