package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type GinApp interface {
	Registry(router gin.IRouter)
	Name() string
	Version() string
}

var (
	Apps = make(map[string]GinApp, 0)
)

func RegisterApp(app GinApp) {
	_, ok := Apps[app.Name()]
	if ok {
		panic(fmt.Sprintf("app %s has register", app.Name()))
	}
	Apps[app.Name()] = app
}
