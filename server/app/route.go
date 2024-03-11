package app

import (
	"strings"
)

type Route struct {
	Path    string
	Command string
	Handler func(req Request)
}

func (app *App) AddRoute(command string, path string, handler func(req Request)) {
	app.Routes = append(app.Routes, Route{Path: path, Command: command, Handler: handler})
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