package route

import (
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/srs/node-repo-proxy/node"
)

func nodeMetaData(c echo.Context) error {
	return c.XMLBlob(http.StatusOK, node.MetaData)
}

func nodeMetaDataMd5(c echo.Context) error {
	return serveMd5(c, node.MetaData)
}

func nodeMetaDataSha1(c echo.Context) error {
	return serveSha1(c, node.MetaData)
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

func serveResp(c echo.Context, resp *http.Response) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return c.String(resp.StatusCode, resp.Status)
	}

	for k, v := range resp.Header {
		c.Response().Header().Set(k, v[0])
	}

	c.Response().WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}

func nodeFile(c echo.Context) error {
	url := buildNodeUrl(c)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}

func nodeFileHead(c echo.Context) error {
	url := buildNodeUrl(c)

	resp, err := http.Head(url)
	if err != nil {
		return err
	}

	return serveResp(c, resp)
}
