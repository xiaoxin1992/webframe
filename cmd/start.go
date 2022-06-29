package cmd

// 启动服务

import (
	"os"
	"webframe/pkg/logger"
	"webframe/protocol"
)

type services struct {
	http protocol.Service
}

func (s *services) Start() error {
	logger.Logger.S("cli").Infof("start http app")
	return s.http.Start()
}

func (s *services) waitSign(sign chan os.Signal) {
	for {
		select {
		case sg := <-sign:
			switch v := sg.(type) {
			default:
				logger.Logger.S("http").Infof("receive signal %v, shutdown", v.String())
				if err := s.http.Stop(); err != nil {
					logger.Logger.S("http").Error("http graceful shutdown err: %s force exit", err)
				} else {
					logger.Logger.S("http").Info("http service stop complete")
				}
				return
			}
		}
	}
}

func newService() *services {
	http := protocol.NewGin()
	return &services{
		http: http,
	}
}
