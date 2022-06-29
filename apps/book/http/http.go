package http

import (
	"github.com/gin-gonic/gin"
	"webframe/pkg/app"
)

var h = &handler{}

type handler struct {
	service string
	log     string
}

func (h *handler) Version() string {
	//TODO implement me
	return "v1"
}

func (h *handler) Config() error {
	h.log = ""
	h.service = "xxx"
	return nil
}

func (h *handler) Name() string {
	return "books"
}

func (h *handler) Registry(r gin.IRouter) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"book": "1"})
	})
}

func init() {
	app.RegisterApp(h)
}
