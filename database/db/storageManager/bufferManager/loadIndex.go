package buffermanager

import (
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
)

func LoadIndex(t *structure.Table) error {
	indexFilePath := t.Metadata.MetadataPath + "/" + t.Metadata.Name + ".tti"
	if _, err := os.Stat(indexFilePath); err != nil {
		return err
	}
	indexFile, err := os.Open(indexFilePath)
	if err != nil {
		return fmt.Errorf("failed to open index file '%s': %v", indexFilePath, err)
	}
	defer indexFile.Close()

	// Initialize a decoder for reading binary data
	indexDecoder := gob.NewDecoder(indexFile)

	var indexTable *structure.Index
	if err := indexDecoder.Decode(&indexTable); err != nil {
		return fmt.Errorf("failed to decode index data from '%s': %v", indexFilePath, err)
	}

	if err != nil {
		return err
	}
	t.IndexTable = indexTable
	return nil
}
