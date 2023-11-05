package app

import (
	"io"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/flydog-sdr/flycat-admin/config"
)

type ServerOptions struct {
	Gzip           int
	Cors           bool
	Version        string
	ApiPrefix      string
	WebPrefix      string
	Config         *config.Config
	SerialPort     io.ReadWriteCloser
	Preference     *config.Preference
	AuthMiddleware *jwt.GinJWTMiddleware
}
