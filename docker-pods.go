package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/lotreal/docker-pods/src/pods"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-pods"
	app.Usage = "manage your dockers"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "docker run a pod",
			Action: func(c *cli.Context) {
				// pods.Run(c.Args().First())

				fmt.Printf("%#v\n", c.Args().First())
				fmt.Printf("%#v\n", c.Args().Present())
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "list pods",
			Action: func(c *cli.Context) {
				pods.Status()
			},
		},
	}
	app.Run(os.Args)
}
