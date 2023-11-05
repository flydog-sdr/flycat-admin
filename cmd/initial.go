package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/flydog-sdr/flycat-admin/common/utils"
	"github.com/flydog-sdr/flycat-admin/config"
)

func ProgramInit(args *config.Args, conf *config.Config, pref *config.Preference) error {
	figure.NewFigure(
		"FlyCat",
		"standard",
		true,
	).Print()
	fmt.Println()

	args.Read()
	if args.Version {
		PrintVersion()
		os.Exit(0)
	}

	err := conf.Read(args.Path)
	if err != nil {
		return err
	}

	if !utils.IsFileExist(conf.FlyCat.Preference) {
		err = pref.Init(conf.FlyCat.Preference)
		if err != nil {
			return err
		}

		return nil
	}

	err = pref.Read(conf.FlyCat.Preference)
	if err != nil {
		return err
	}

	return nil
}
