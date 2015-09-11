package main

import (
	"github.com/guregu/kami"
	"github.com/konjoot/try-kami/github"
	"golang.org/x/net/context"
	"net/http"
)

func kamiHandle(ctx context.Context, w http.ResponseWriter, r *http.Request) {}

func loadKami(routes []github.Route) {
	h := kamiHandle

	for _, route := range routes {
		kami.Handle(route.Method, route.Path, h)
	}
}

func main() {
	loadKami(github.API)
	kami.Serve()
}
