package main

import (
	"log"

	"github.com/flydog-sdr/flycat-admin/app"
	"github.com/flydog-sdr/flycat-admin/common/serial"
	"github.com/flydog-sdr/flycat-admin/config"
	"github.com/flydog-sdr/flycat-admin/features/controller"
	"github.com/flydog-sdr/flycat-admin/server"
)

func main() {
	var (
		args config.Args
		conf config.Config
		pref config.Preference
	)

	err := ProgramInit(&args, &conf, &pref)
	if err != nil {
		log.Fatalln(err)
	}

	port, err := serial.PairDevice()
	if err != nil {
		log.Fatalln(err)
	}
	defer port.Close()

	err = controller.SetRegister(port, pref.Gain, pref.Antenna, pref.Filter)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("FlyCat Admin daemon initialized")
	server.ServerDaemon(
		conf.Server.Host,
		conf.Server.Port,
		&app.ServerOptions{
			Gzip:       9,
			WebPrefix:  "/",
			ApiPrefix:  "/api",
			Cors:       true,
			SerialPort: port,
			Config:     &conf,
			Preference: &pref,
			Version:    "v1",
		})
}
