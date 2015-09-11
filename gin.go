package main

import (
	"github.com/gin-gonic/gin"
	"github.com/konjoot/try-kami/github"
	"io"
)

func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Params.ByName("name"))
}

func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGin(routes []github.Route) *gin.Engine {
	h := ginHandle

	router := gin.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, h)
	}
	return router
}

func main() {
	router := loadGin(github.API)
	router.Run(":8000")
}
