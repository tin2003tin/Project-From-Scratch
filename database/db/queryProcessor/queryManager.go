package queryProcessor

import "database/db/structure"

type QueryManager struct {
	Table          *structure.Table
	CurrentColumns []structure.Column
	CurrentRows    [][]*structure.Row
	CurrentIndexes   []*structure.Index
}

func NewQueryManager(t *structure.Table) *QueryManager {
	q := QueryManager{
		Table:          t,
		CurrentColumns: t.Columns,
		CurrentRows:    nil,
	}
	return &q
}

func (q *QueryManager) ResetCurrent() {
	q.CurrentColumns = q.Table.Columns
	q.CurrentRows = nil
}
