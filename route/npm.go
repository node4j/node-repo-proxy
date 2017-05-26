package route

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/srs/node-repo-proxy/npm"
)

func buildNpmUrl(c echo.Context) string {
	name := c.Param("name")
	version := c.Param("version")
	len := len(name + "-" + version)

	rest := c.Param("file")[len:]

	ext := strings.Join(strings.Split(rest, ".")[1:], ".")
	url := "https://registry.npmjs.org/" + name + "/-/" + name + "-" + version + "." + ext

	return url
}

func npmFile(c echo.Context) error {
	url := buildNpmUrl(c)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}

func npmFileHead(c echo.Context) error {
	url := buildNpmUrl(c)

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}

func npmMetaData(c echo.Context) error {
	name := c.Param("name")
	data, err := npm.GetMetaData(name)

	if err != nil {
		return err
	}

	if data == nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.XMLBlob(http.StatusOK, data)
}
