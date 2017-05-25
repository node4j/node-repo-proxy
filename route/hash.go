package route

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func isSha(c echo.Context) bool {
	return strings.HasSuffix(c.QueryString(), ".sha")
}

func isMd5(c echo.Context) bool {
	return strings.HasSuffix(c.QueryString(), ".md5")
}

func serveSha1(c echo.Context, data []byte) error {
	hasher := sha1.New()
	hasher.Write(data)

	return c.String(http.StatusOK, hex.EncodeToString(hasher.Sum(nil)))
}

func serveMd5(c echo.Context, data []byte) error {
	hasher := md5.New()
	hasher.Write(data)

	return c.String(http.StatusOK, hex.EncodeToString(hasher.Sum(nil)))
}
