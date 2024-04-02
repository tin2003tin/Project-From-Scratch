package betasql

import (
	"database/db/queryProcessor"
	buffermanager "database/db/storageManager/bufferManager"
	"database/db/structure"
	"sync"
)

func (sdb *SelectDatabase) loadMultiFile(tableName []string) ([]*structure.Table, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var tables []*structure.Table
	var threadErr error

	for _, tableName := range tableName {
		if loadedtable, ok := sdb.DataBase.Registry.Tables[tableName]; !ok {
			wg.Add(1)
			go func(tName string) {
				defer wg.Done()
				table, err := sdb.loadFile(tName)
				if err != nil {
					mu.Lock()
					threadErr = err
					mu.Unlock()
					return
				}
				mu.Lock()
				tables = append(tables, table)
				mu.Unlock()
			}(tableName)
		} else {
			tables = append(tables, loadedtable)
		}
	}

	wg.Wait()

	if threadErr != nil {
		return nil, threadErr
	}
	return tables, nil
}

func (sdb *SelectDatabase) loadFile(tableName string) (*structure.Table, error) {
	var table *structure.Table
	var err error
	table, err = buffermanager.LoadTableMetadata(sdb.DataBase, tableName)
	if err != nil {
		return nil, err
	}

	err = buffermanager.LoadIndex(table)
	if err != nil {
		return nil, err
	}

	err = buffermanager.LoadRawData(table)
	if err != nil {
		return nil, err
	}

	err = buffermanager.BuildIndex(table)
	if err != nil {
		return nil, err
	}
	queryProcessor.NewQueryManager(table)

	return table, nil
}
