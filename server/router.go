package server

import (
	"github.com/flydog-sdr/flycat-admin/app"
	"github.com/flydog-sdr/flycat-admin/app/controller"
	"github.com/flydog-sdr/flycat-admin/app/status"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(rg *gin.RouterGroup, options *app.ServerOptions) {
	services := []ApiServices{
		&controller.Controller{},
		&status.Status{},
	}

	for _, s := range services {
		s.RegisterModule(rg, options)
	}
}
