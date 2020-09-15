package main

import (
	"fmt"
	"github.com/rz1226/blackboardkit"
	"github.com/rz1226/fasthttpserv"
	"github.com/valyala/fasthttp"
)

//web服务器， 对fasthttp的封装
const (
	RETURN_OK = `{"code":0,"msg":"ok"}`

	RETURN_NOT_VALID_JSON = `{"code":-1,"msg":"not valid json error"}`
	RETURN_ERR            = `{"code":-1,"msg":"err"}`
)

func Index(ctx *fasthttp.RequestCtx) string {

	return `{"hello world"}`

}

func StartServ(port string) {
	blackboardkit.StartMonitor("9090")

	serv := fasthttpserv.NewServ()
	serv.GET("/", Index)

	fmt.Println("ready to start")
	serv.StartCORS(port)
}
