package structure

type JoinType int

const (
	InnerJoin JoinType = iota
	LeftJoin
	RightJoin
	FullJoin
)

type On struct {
	Self     string // Name of the column to check condition against
	Operator string // Operator for comparison (e.g., "=", ">", "<", etc.)
	Another  string
}