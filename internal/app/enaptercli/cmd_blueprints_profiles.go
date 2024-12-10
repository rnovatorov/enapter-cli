package enaptercli

import (
	"github.com/urfave/cli/v2"
)

type cmdBlueprintsProfiles struct {
	cmdBase
}

func buildCmdBlueprintsProfiles() *cli.Command {
	return &cli.Command{
		Name:  "profiles",
		Usage: "Manage blueprints profiles",
		Subcommands: []*cli.Command{
			buildCmdBlueprintsProfilesDownload(),
			buildCmdBlueprintsProfilesUpload(),
		},
	}
}
