package runtime

import (
	betasql "database/betasql"
	buffermanager "database/db/storageManager/bufferManager"
	"database/db/structure"
	"fmt"
	"net"
	"os"
)

type App struct {
	SqlCompliler   betasql.SqlCompliler
	LoadedDatabase []structure.Database
}

func CreateApp() *App {
	app := &App{}
	return app
}

func (app *App) OpenServer(port string, callback func(error)) {
	var main_domain string = ":" + port
	listener, err := net.Listen("tcp", main_domain)
	if err != nil {
		callback(err)
	}
	defer listener.Close()
	fmt.Println("Listen to port :" + port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			callback(err)
			continue
		}
		go app.handleClient(conn)
	}
}

func (app *App) LoadMetaDatabase(databaseFolder string) error {
	// Open the directory
	dir, err := os.Open(databaseFolder)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return err
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.ReadDir(0) // Read all entries
	if err != nil {
		fmt.Println("Error reading directory entries:", err)
		return err
	}

	// Iterate over the entries and print folder names
	for _, entry := range entries {
		database, err := buffermanager.LoadDatabaseMetadata(entry.Name())
		if err != nil {
			fmt.Println(err)
			return err
		}
		app.LoadedDatabase = append(app.LoadedDatabase, *database)
	}
	return nil
}

func (app *App) InitCompiler() {
	app.SqlCompliler = betasql.InitSqlCompiler()
}

func (app *App) LoadTableMetadata() {
	for i := range app.LoadedDatabase {
		for _, table := range app.LoadedDatabase[i].TableNames {
			_, err := buffermanager.LoadTableMetadata(&app.LoadedDatabase[i], table)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
