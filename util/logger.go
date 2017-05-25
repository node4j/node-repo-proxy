package util

import "github.com/labstack/echo"

var Log echo.Logger

func SetLogger(log echo.Logger) {
	Log = log
}
