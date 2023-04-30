package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br"},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "busca IPS de endereços na internet",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: searchServers,
		},
	}
	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")

	server, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range server {
		fmt.Println(server.Host)
	}

}
