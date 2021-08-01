package main

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Echo(ctx *fasthttp.RequestCtx) {
	_, err := ctx.WriteString(string(ctx.QueryArgs().Peek("msg")))
	if err != nil {
		return
	}
}

func Handler() func(request *fasthttp.RequestCtx) {
	api := router.New()

	api.GET("/echo", Echo)

	return api.Handler
}

func main() {
	if err := fasthttp.ListenAndServe("localhost:80", Handler()); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
