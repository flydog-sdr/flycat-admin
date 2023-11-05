package server

import (
	"github.com/flydog-sdr/flycat-admin/app"
	"github.com/gin-gonic/gin"
)

type ApiServices interface {
	RegisterModule(rg *gin.RouterGroup, options *app.ServerOptions)
}
