package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/jonathantneal/asana/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "asana"
	app.Version = "0.2.1"
	app.Usage = "asana cui client ( https://github.com/jonathantneal/asana )"

	app.Commands = defs()
	app.Run(os.Args)
}

func defs() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Asana configuration. Your settings will be saved in ~/.asana.yml",
			Action: func(c *cli.Context) error {
				commands.Config(c)
				return nil
			},
		},
		{
			Name:    "workspaces",
			Aliases: []string{"w"},
			Usage:   "get workspaces",
			Action: func(c *cli.Context) error {
				commands.Workspaces(c)
				return nil
			},
		},
		{
			Name:    "tasks",
			Aliases: []string{"ts"},
			Usage:   "get tasks",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "no-cache, n", Usage: "without cache"},
				&cli.BoolFlag{Name: "refresh, r", Usage: "update cache"},
			},
			Action: func(c *cli.Context) error {
				commands.Tasks(c)
				return nil
			},
		},
		{
			Name:    "task",
			Aliases: []string{"t"},
			Usage:   "get a task",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "verbose, v", Usage: "verbose output"},
			},
			Action: func(c *cli.Context) error {
				commands.Task(c)
				return nil
			},
		},
		{
			Name:    "comment",
			Aliases: []string{"cm"},
			Usage:   "Post comment",
			Action: func(c *cli.Context) error {
				commands.Comment(c)
				return nil
			},
		},
		{
			Name:  "done",
			Usage: "Complete task",
			Action: func(c *cli.Context) error {
				commands.Done(c)
				return nil
			},
		},
		{
			Name:  "due",
			Usage: "set due date",
			Action: func(c *cli.Context) error {
				commands.DueOn(c)
				return nil
			},
		},
		{
			Name:    "browse",
			Aliases: []string{"b"},
			Usage:   "open a task in the web browser",
			Action: func(c *cli.Context) error {
				commands.Browse(c)
				return nil
			},
		},
	}
}
