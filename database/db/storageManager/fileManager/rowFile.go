package filemanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func SaveRawData(t *structure.Table) error {
	ttrFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".ttr")

	// Create or open the .ttr file
	ttrFile, err := os.Create(ttrFilePath)
	if err != nil {
		return fmt.Errorf("failed to create .ttr file: %v", err)
	}
	defer ttrFile.Close()

	// Initialize an encoder for writing binary data
	encoder := gob.NewEncoder(ttrFile)

	// Encode and write the rows to the .ttr file
	if err := encoder.Encode(t.Rows); err != nil {
		return fmt.Errorf("failed to encode rows: %v", err)
	}
	return nil
}
