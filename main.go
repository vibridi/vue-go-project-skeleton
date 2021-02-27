package vgserver

import (
	"log"
	"net/http"
	"os"
	"time"
	"vgserver/router"

	"github.com/tylerb/graceful"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "app"
	app.Version = "v0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			EnvVar: "FRONTEND_ADDR",
			Name:   "addr",
		},
	}

	app.Action = func(c *cli.Context) error {

		r := router.New()

		srv := &graceful.Server{
			Timeout: 10 * time.Second,
			Server: &http.Server{
				Addr:    c.String("addr"),
				Handler: r,
			},
		}
		log.Println("starting server at", srv.Addr)

		return srv.ListenAndServe()
	}

	app.Run(os.Args)
}
