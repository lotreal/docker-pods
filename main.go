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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) {
		pods.Ps()
		// println(c.String("lang"))
		// name := "someone"
		// if len(c.Args()) > 0 {
		// 	name = c.Args()[0]
		// }
		// if c.String("lang") == "spanish" {
		// 	println("Hola", name)
		// } else {
		// 	println("Hello", name)
		// }
	}

	app.Commands = []cli.Command{
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "lang",
					Value: "english",
					Usage: "language for the greeting",
				},
			},
			Action: func(c *cli.Context) {
				fmt.Println(c.Bool("version"))
				fmt.Println(c.String("lang"))
				fmt.Println(c.Args().First())
			},
		},
		{
			Name:    "ls",
			Aliases: []string{"l"},
			Usage:   "list pods",
			Action: func(c *cli.Context) {
				// pods.List()
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "docker run a pod",
			Action: func(c *cli.Context) {
				pods.Run()
			},
		},
		{
			Name:    "status",
			Aliases: []string{"ps"},
			Usage:   "list pods",
			Action: func(c *cli.Context) {
				pods.Ps()
			},
		},
	}
	app.Run(os.Args)
}
