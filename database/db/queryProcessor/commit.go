package queryProcessor

import (
	filemanager "database/db/storageManager/fileManager"
	"fmt"
)

func (q *QueryManager) Commit() error {
	// Update the index file
	if err := filemanager.UpdateIndexFile(q.Table); err != nil {
		return fmt.Errorf("failed to save index file: %v", err)
	}
	// Update the row file
	if err := filemanager.SaveRawData(q.Table); err != nil {
		return fmt.Errorf("failed to save raw data file: %v", err)
	}

	fmt.Printf("Rows commited to '%s' successfully\n", q.Table.Metadata.Name)
	return nil
}
