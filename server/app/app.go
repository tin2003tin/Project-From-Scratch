package app

import (
	"fmt"
	"net"
)

type App struct {
	Routes []Route
}

func CreateApp() *App {
	app := &App{}
	return app
}

func (app *App) OpenServer(port string, callback func(error)) {
	var main_domain string = "localhost:"+port;
	listener, err := net.Listen("tcp", main_domain);
	if err != nil {
		callback(err)
	}
	defer listener.Close()
	fmt.Println("Listen to port :"+port);
	for {
		conn,err := listener.Accept();
		if err != nil {
			callback(err)
			continue
		}
		go app.handleClient(conn)
	}
}

