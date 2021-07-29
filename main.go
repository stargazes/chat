package main

import (
	"bufio"
	"chat/controller"
	"chat/middleware"
	"chat/models"
	"chat/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

var g errgroup.Group //主要为了开启协程，记录协程中日志的错误信息

//http服务入口
func httpRoute() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery()) //中间件
	e.Use(middleware.Logger()) //中间件
	e.GET("/", func(c *gin.Context) {
		value := controller.TestRedis()
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
				"value": value,
			},
		)
	})
	//e.POST("/init", InitModel) //
	e.POST("/user/register", controller.Register) //用户注册
	e.POST("/user/login", controller.Login)       //用户登录

	//需要登录才能访问的路由
	loginRoute := e.Group("/room")
	loginRoute.Use(middleware.JWTAuth())
	{
		loginRoute.POST("/create", controller.CreateRoom) //创建聊天室
		loginRoute.POST("/add", controller.AddRoom)       //加入聊天室

	}

	return e
}


/**
生成表格
 */
func InitModel(c *gin.Context)  {
	err :=tools.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{},&models.Room{})

	if err != nil {

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "数据库配置有误",
			},
		)
		return

	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": "数据库初始化成功",
		},
	)

	return
}

func main() {

	httpServer := &http.Server{

		Addr:         ":8081",     //指定端口
		Handler:      httpRoute(), //定义处理函数
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}


	//go func() {
	//     httpServer.ListenAndServe()
	//}()
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}


}

func ListenAndServe(address string){
	//绑定监听地址
	listener,err := net.Listen("tcp",address)

	if err != nil {
		log.Fatal(fmt.Sprint("listen err :%v",err))
	}
	defer listener.Close()

	log.Println(fmt.Sprintf("bind:%s,start listening...",address))

	for  {
		conn,err:=listener.Accept()
		if err!=nil {
			log.Fatal(fmt.Sprintf("accept err:%v",err))
		}
		go Handle(conn)
	}

}

func Handle(conn net.Conn)  {

	reader := bufio.NewReader(conn)
	for  {
		msg,err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
			}else {
				log.Println(err)
			}
			return
		}
		b:=[]byte(msg)
		conn.Write(b)
	}
}
