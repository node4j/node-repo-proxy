package route

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/srs/node-repo-proxy/node"
)

func nodeMetaData(c echo.Context) error {
	data, err := node.GetMetaData()
	return serveMetaData(c, data, err)
}

func buildNodeUrl(c echo.Context) string {
	file := c.Param("file")
	version := c.Param("version")

	last := strings.Join(strings.Split(file, "-")[2:], "-")
	arch := strings.Split(last, ".")[0]
	ext := strings.Join(strings.Split(last, ".")[1:], ".")
	url := "https://nodejs.org/dist/v" + version + "/"

	if (arch == "win-x64" || arch == "win-x86") && ext == "exe" {
		url += arch + "/node.exe"
	} else {
		url += "node-v" + version + "-" + arch + "." + ext
	}

	return url
}

/*
func nodeFile(c echo.Context) error {
	url := buildNodeUrl(c)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}
*/

func nodeFileRedirect(c echo.Context) error {
	url := buildNodeUrl(c)
	return serveRedirect(c, url)
}

/*
func nodeFileHead(c echo.Context) error {
	url := buildNodeUrl(c)

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}
*/
