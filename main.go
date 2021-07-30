package main

import (
	"chat/middleware"
	"chat/route"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group //主要为了开启协程，记录协程中日志的错误信息

//http服务入口
func httpRoute() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery()) //中间件
	e.Use(middleware.Logger()) //中间件
	route.RoutesForAll(e)
	return e
}

func main() {

	httpServer := &http.Server{

		Addr:         ":8081",     //指定端口
		Handler:      httpRoute(), //定义处理函数
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}


}
