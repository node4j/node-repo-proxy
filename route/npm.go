package route

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/node4j/node-repo-proxy/npm"
)

func buildNpmUrl(c echo.Context) string {
	name := c.Param("name")
	version := c.Param("version")
	len := len(name + "-" + version)

	rest := c.Param("file")[len:]

	ext := strings.Join(strings.Split(rest, ".")[1:], ".")
	if ext == "tar.gz" {
		ext = "tgz"
	}

	url := "https://registry.npmjs.org/" + name + "/-/" + name + "-" + version + "." + ext

	return url
}

/*
func npmFile(c echo.Context) error {
	url := buildNpmUrl(c)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}
*/

func npmFileRedirect(c echo.Context) error {
	url := buildNpmUrl(c)
	return serveRedirect(c, url)
}

/*
func npmFileHead(c echo.Context) error {
	url := buildNpmUrl(c)

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}
*/

func npmMetaData(c echo.Context) error {
	name := c.Param("name")
	data, err := npm.GetMetaData(name)

	return serveMetaData(c, data, err)
}
