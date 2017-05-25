package route

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/org/node/node/maven-metadata.xml", nodeMetaData)
	e.GET("/org/node/node/maven-metadata.xml.md5", nodeMetaDataMd5)
	e.GET("/org/node/node/maven-metadata.xml.sha", nodeMetaDataSha1)
	e.GET("/org/node/node/:version/:file", nodeFile)
}
