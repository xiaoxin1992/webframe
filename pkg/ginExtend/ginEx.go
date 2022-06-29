package ginExtend

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"time"
	"webframe/pkg/app"
	"webframe/pkg/logger"
)

type GinExtend struct {
	Route   *gin.Engine
	Service http.Server
}

func (g *GinExtend) registerRoute() {
	for _, app := range app.Apps {
		app.Registry(g.Route.Group(path.Join(app.Name(), app.Version())))
	}
}
func (g *GinExtend) Start() (err error) {
	g.registerRoute()
	logger.Logger.S("gin").Infof("http listen: %s", g.Service.Addr)
	err = g.Service.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			logger.Logger.S("gin").Info("service is stopped")
		}
		return fmt.Errorf("start gin error %s", err.Error())
	}
	return nil
}

func (g *GinExtend) Stop() error {
	logger.Logger.S("gin").Info("shutdown gin http service")
	ctx, concanl := context.WithTimeout(context.Background(), 30*time.Second)
	defer concanl()
	if err := g.Service.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
