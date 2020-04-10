package commands

import (
	"strconv"
	"os/exec"

    "github.com/codegangsta/cli"

	"github.com/jonathantneal/asana/api"
	"github.com/jonathantneal/asana/config"
	"github.com/jonathantneal/asana/utils"
)

func Browse(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
