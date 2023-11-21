package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a aplcação pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "aplicação de Linha de comando"
	app.Usage = "Buscar IPs e Nomes de servidores na internet"

	flags :=  []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereço na intenet",
			Flags: flags,
			Action: buscarIps ,
		},{
			Name: "servidores",
			Usage: "Busca o nome de servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	
	return app
}

func buscarIps(c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")
	servidores, erro := net.LookupNS(host) //Nome do servidor
	if  erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	}
}