package buffermanager

import (
	db_constant "database/db/constant"
	"database/db/structure"
	"encoding/gob"
	"fmt"
	"os"
)

func LoadDatabaseMetadata(name string) (*structure.Database, error) {
	metaFilePath := db_constant.DATABASE_PATH + fmt.Sprintf("%s/%s.meta", name, name)
	metaFile, err := os.Open(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open metadata file: %v", err)
	}
	defer metaFile.Close()

	decoder := gob.NewDecoder(metaFile)
	var db structure.Database
	if err := decoder.Decode(&db); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %v", err)
	}
	return &db, nil
}
