package protocol

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"webframe/conf"
	"webframe/pkg/ginExtend"
	"webframe/pkg/ginExtend/middleware"
	"webframe/pkg/logger"
)

func NewGin() *ginExtend.GinExtend {
	httpConfig := conf.GetConfig().Http
	if httpConfig.Level != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.New()
	route.SetTrustedProxies(nil)
	route.Use(middleware.ZapWithConfig(logger.Logger), middleware.RecoveryWithZap(logger.Logger, true))
	return &ginExtend.GinExtend{
		Route: route,
		Service: http.Server{
			Addr:              fmt.Sprintf("%s:%d", httpConfig.Host, httpConfig.Port),
			Handler:           route,
			ReadTimeout:       60 * time.Second,
			ReadHeaderTimeout: 60 * time.Second,
			WriteTimeout:      60 * time.Second,
			IdleTimeout:       60 * time.Second,
			MaxHeaderBytes:    1 << 25},
	}
}
