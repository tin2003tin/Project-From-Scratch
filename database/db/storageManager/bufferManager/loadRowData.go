package buffermanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

func LoadRawData(t *structure.Table) error {
	rawDataFilePath := t.Metadata.MetadataPath + "/" + t.Metadata.Name + ".ttr"
	if _, err := os.Stat(rawDataFilePath); err != nil {
		return err
	}
	// Construct the path for the .ttr file
	ttrFilePath := filepath.Join(t.Metadata.MetadataPath, t.Metadata.Name+".ttr")

	// Open the .ttr file for reading
	ttrFile, err := os.Open(ttrFilePath)
	if err != nil {
		return fmt.Errorf("failed to open .ttr file: %v", err)
	}
	defer ttrFile.Close()

	// Initialize a decoder for reading binary data
	decoder := gob.NewDecoder(ttrFile)

	// Decode the rows from the .ttr file
	var rows []structure.Row
	if err := decoder.Decode(&rows); err != nil {
		return fmt.Errorf("failed to decode rows from .ttr file: %v", err)
	}

	// Update the table's rows with the decoded rows
	t.Rows = rows

	fmt.Printf("Rows read from '%s' successfully\n", ttrFilePath)
	return nil
}
