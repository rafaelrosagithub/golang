package app

import "github.com/urfave/cli"

// Generate will return the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search ips and server names on the internet"

	return app
}
