package main

import (
	"chat/controller"
	"chat/impl"
	"chat/middleware"
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandle(c *gin.Context) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	//完成ws协议的握手操作
	//upgrade:websocket
	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		log.Println("握手失败")
		goto ERR
	}
	//启动线程
	//心跳包维持链接
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()

}

var g errgroup.Group //主要为了开启协程，记录协程中日志的错误信息

//http服务入口
func httpRoute() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery()) //中间件
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

//websocket入口
func websocketRoute() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", wsHandle)
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

		Addr:         ":8080",     //指定端口
		Handler:      httpRoute(), //定义处理函数
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	websocketServer := &http.Server{

		Addr:         ":10001",
		Handler:      websocketRoute(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//go func() {
	//     httpServer.ListenAndServe()
	//}()
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		return websocketServer.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
