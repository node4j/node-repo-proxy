package main

import (
	"os"

	"github.com/labstack/echo"
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

	route.Init(e)
	util.Log.Fatal(e.Start(":" + port))
}
