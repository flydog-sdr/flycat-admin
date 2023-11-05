package controller

import (
	"net/http"

	"github.com/flydog-sdr/flycat-admin/app"
	"github.com/flydog-sdr/flycat-admin/config"
	"github.com/flydog-sdr/flycat-admin/features/controller"
	"github.com/flydog-sdr/flycat-admin/server/response"
	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterModule(rg *gin.RouterGroup, option *app.ServerOptions) {
	rg.POST("/controller", func(c *gin.Context) {
		if option.SerialPort == nil {
			response.ErrorHandler(c, http.StatusMethodNotAllowed)
			return
		}

		var binding Binding
		if err := c.ShouldBind(&binding); err != nil {
			response.ErrorHandler(c, http.StatusBadRequest)
			return
		}

		err := controller.SetRegister(option.SerialPort, binding.Gain, binding.Antenna, binding.Filter)
		if err != nil {
			response.ErrorHandler(c, http.StatusBadRequest)
			return
		}

		option.Preference = &config.Preference{
			Gain:    binding.Gain,
			Filter:  binding.Filter,
			Antenna: binding.Antenna,
		}

		err = option.Preference.Update(option.Config.FlyCat.Preference)
		if err != nil {
			response.ErrorHandler(c, http.StatusInternalServerError)
			return
		}

		response.MessageHandler(c, "Controller data successfully updated", nil)
	})

	rg.GET("/controller", func(c *gin.Context) {
		res := make(map[any]int)
		res["gain"] = option.Preference.Gain
		res["antenna"] = option.Preference.Antenna
		res["filter"] = option.Preference.Filter

		response.MessageHandler(c, "Successfully acquired controller status", res)
	})
}
