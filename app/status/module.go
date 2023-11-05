package status

import (
	"github.com/flydog-sdr/flycat-admin/app"
	"github.com/flydog-sdr/flycat-admin/server/response"
	"github.com/gin-gonic/gin"
)

func (s *Status) RegisterModule(rg *gin.RouterGroup, options *app.ServerOptions) {
	rg.GET("/status", func(c *gin.Context) {
		response.MessageHandler(c, "Successfully acquired system status", System{
			Controller: GetController(options.Preference),
			Memory:     GetMemory(),
			Uptime:     GetUptime(),
			CPU:        GetCPU(),
			OS:         GetOS(),
		})
	})
}
