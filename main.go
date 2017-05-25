package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/robfig/cron"
	"github.com/srs/node-repo-proxy/node"
	"github.com/srs/node-repo-proxy/route"
	"github.com/srs/node-repo-proxy/util"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	e := echo.New()
	util.SetLogger(e.Logger)

	err := node.LoadMetaData()
	if err != nil {
		panic(err)
	}

	c := cron.New()
	c.AddFunc("@midnight", node.UpdateMetaData)
	c.Start()

	route.Init(e)
	util.Log.Fatal(e.Start(":" + port))
}
