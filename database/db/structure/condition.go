package structure

type Condition struct {
	ColumnName string      // Name of the column to check condition against
	Operator   string      // Operator for comparison (e.g., "=", ">", "<", etc.)
	Value      interface{} // Value to compare against
}
