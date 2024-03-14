package app

import (
	"fmt"
	tinPro "server/app/protocol"
	"strings"
)

type Route struct {
	Path    string
	Command string
	Handler func(req Request, res Response) error
}

func (app *App) AddRoute(command string, path string, handler func(req Request,res Response) error) {
	app.Routes = append(app.Routes, Route{Path: path, Command: command, Handler: handler})
}

//Find handler Route and excute that handler
func (app *App)handleCommand(protocol *tinPro.TinReqProtocol) (func(Request, Response) error, map[string]string, error) {
	index, params := app.findHandlerRoute(protocol.GetHeader().Command, protocol.GetHeader().Path)
	if index != -1 {
		return app.Routes[index].Handler, params, nil
	}
	return nil, nil, fmt.Errorf("handler not found for path: %s with command: %s", protocol.GetHeader().Path, protocol.GetHeader().Command)
}


func (app *App) findHandlerRoute(command string, path string) (int, map[string]string) {
	params := make(map[string]string)
	for i, route := range app.Routes {
		if route.Command == command {
			routeParts := strings.Split(route.Path, "/")
			pathParts := strings.Split(path, "/")
			if len(routeParts) != len(pathParts) {
				continue
			}
			match := true
			for j, part := range routeParts {
				if part != pathParts[j] {
					if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
						params[strings.Trim(part, "{}")] = pathParts[j]
					} else {
						match = false
						break
					}
				}
			}
			if match {
				return i, params
			}
		}
	}
	return -1, nil
}