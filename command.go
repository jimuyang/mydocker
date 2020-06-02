package main

import (
	"fmt"
	"mydocker/container"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var RunCommand = cli.Command{
	Name: "run",
	Usage: `create a container with namespace and cgroups limit
			mydocker run -ti [command]`,

	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},

	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}
		tty := ctx.Bool("ti")
		Run(tty, cmdArray)
		return nil
	},
}

var InitCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}
