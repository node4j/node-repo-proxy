package route

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/org/node/node/maven-metadata.xml", nodeMetaData)
	e.HEAD("/org/node/node/maven-metadata.xml", nodeMetaData)

	e.GET("/org/node/node/:version/:file", nodeFileRedirect)
	e.HEAD("/org/node/node/:version/:file", nodeFileRedirect)

	e.GET("/org/npmjs/:name/:version/:file", npmFileRedirect)
	e.HEAD("/org/npmjs/:name/:version/:file", npmFileRedirect)

	e.GET("/org/npmjs/:name/maven-metadata.xml", npmMetaData)
	e.HEAD("/org/npmjs/:name/maven-metadata.xml", npmMetaData)
}
