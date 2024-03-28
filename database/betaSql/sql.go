package betasql

import (
	"database/compiler"
	"database/db"
)

type Sql struct {
	Compiler       compiler.Compiler
	LoadedDatabase []db.Database
}

func NewSql() *Sql {
	sql := Sql{
		Compiler: *compiler.InitCompiler(SQL_GRAMMER),
	}
	return &sql
}

func (sql *Sql) LoadDatabase(databae string) (*db.Database, error) {
	database, err := db.GetDataBase(databae)
	if err != nil {
		return nil, err
	}
	sql.LoadedDatabase = append(sql.LoadedDatabase, *database)
	return database, nil
}

func (sql *Sql) GetDataBase(name string) (*db.Database, error) {
	for _, database := range sql.LoadedDatabase {
		if database.Name == name {
			return &database, nil
		}
	}
	database, err := sql.LoadDatabase(name)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func (sql *Sql) ReadSql(input string) {

}
