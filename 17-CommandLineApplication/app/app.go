package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Search ips and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Internet address search IPS",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: getIps,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
