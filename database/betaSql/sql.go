package betasql

import (
	"database/compiler"
	"database/db/structure"
	"fmt"
)

type SqlCompliler struct {
	Compiler compiler.Compiler
}

func InitSqlCompiler() SqlCompliler {
	com := compiler.InitCompiler(SQL_GRAMMER)
	sql := SqlCompliler{Compiler: *com}
	return sql
}

type SelectDatabase struct {
	DataBase *structure.Database
	Output   []byte
}

func (sq *SqlCompliler) Prase(database *structure.Database, sqltext string) ([]byte, error) {
	lexer := sq.Compiler.NewLexer()
	tokens, err := lexer.Convert(sqltext)
	fmt.Println(tokens)
	if err != nil {
		return nil, err
	}
	selectDb := &SelectDatabase{DataBase: database}
	handlers := compiler.SetOfFunc{Handlers: []func([][]string) ([][]string, error){
		selectDb.loadColumn,
		selectDb.loadSql,
		selectDb.addColumn,
		selectDb.loadTable,
		selectDb.loadOp,
		selectDb.loadType,
		selectDb.loadCondition,
		selectDb.condition,
		selectDb.loadJoin,
		selectDb.loadValue,
		selectDb.loadInsert,
		selectDb.addValue,
		selectDb.loadColumns,
		selectDb.deleteRow,
		selectDb.loadEqual,
		selectDb.updateRow,
		selectDb.loadSet,
		selectDb.loadSets,
		selectDb.addSet,
	}}
	parser := sq.Compiler.NewParser(handlers)
	err = parser.Parse(tokens)
	if err != nil {
		return nil, err
	}
	return selectDb.Output, nil
}
