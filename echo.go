package main

import (
	"github.com/konjoot/try-kami/github"
	"github.com/labstack/echo"
)

func echoHandler(c *echo.Context) error {
	return nil
}

func loadEcho(routes []github.Route) *echo.Echo {
	var h interface{} = echoHandler

	e := echo.New()
	for _, r := range routes {
		switch r.Method {
		case "GET":
			e.Get(r.Path, h)
		case "POST":
			e.Post(r.Path, h)
		case "PUT":
			e.Put(r.Path, h)
		case "PATCH":
			e.Patch(r.Path, h)
		case "DELETE":
			e.Delete(r.Path, h)
		default:
			panic("Unknow HTTP method: " + r.Method)
		}
	}
	return e
}

func main() {
	router := loadEcho(github.API)
	router.Run(":8000")
}
