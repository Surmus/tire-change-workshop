package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/surmus/tire-change-workshop/pkg/manchester"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"time"
)

const (
	version        = "v1.0.0"
	listenPortFlag = "port"
	defaultPort    = 9004
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:  listenPortFlag + ", p",
		Value: fmt.Sprintf("%d", defaultPort),
		Usage: "Port for server to listen incoming connections",
	},
}

// @title Manchester tire workshop API
// @version 1.0
// @description Tire workshop service IOT integration.
// @BasePath /api/v1
func main() {
	app := cli.NewApp()
	app.Version = version
	app.Usage = "Manchester tire workshop API server"
	app.Flags = flags
	app.Action = initServer

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func initServer(c *cli.Context) error {
	listenToPort := c.Uint(listenPortFlag)

	if listenToPort == 0 {
		return fmt.Errorf("invalid server listen port supplied: %s", c.String(listenPortFlag))
	}

	return setupServer(listenToPort)
}

func setupServer(port uint) error {
	apiRouter := manchester.Init()
	// The url pointing to API definition
	swaggerURL := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
	workshopServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      apiRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Infof("application initialized, listening to port %d", port)
	return workshopServer.ListenAndServe()
}
