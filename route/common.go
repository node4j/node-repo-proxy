package route

import (
	"io"
	"net/http"

	"github.com/labstack/echo"
)

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

func serveMetaData(c echo.Context, data []byte, err error) error {
	if err != nil {
		return err
	}

	if data == nil {
		return c.String(http.StatusNotFound, "Not found")
	}

	return c.XMLBlob(http.StatusOK, data)
}
