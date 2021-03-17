package main

import (
	"chat/impl"
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

var g errgroup.Group//主要为了开启协程，记录协程中日志的错误信息

//http服务入口
func httpRoute() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery()) //中间件
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})
	return e
}

//websocket入口
func websocketRoute() http.Handler {
    e:=gin.New()
    e.Use(gin.Recovery())
    e.GET("/",wsHandle)
    return e
}

func main() {

	//tools.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}) //自动生成表格
    httpServer :=&http.Server{

        Addr: ":8080",//指定端口
        Handler: httpRoute(),//定义处理函数
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,

    }

    websocketServer :=&http.Server{

        Addr: ":10001",
        Handler: websocketRoute(),
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,

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

	//r := gin.Default()
    //
	//r.GET("/ws", wsHandle)
	//conn := tools.RedisPool.Get()
	//
	//if conn == nil {
	//    fmt.Println("获取连接失败")
	//}
	//
	////redis操作例子
	//_, err := conn.Do("set", "username", string("lyh"))
	//
	//if err != nil {
	//    fmt.Println("设置值失败")
	//}
	//
	//value, err := redis.String(conn.Do("get", "username"))
	//fmt.Println(value)
	//
	//if err != nil {
	//    fmt.Println("获取username失败")
	//}
	//
	////mysql操作例子
	//var user models.User
	////user.Username="lyh"
	////user.Password="123456"
	////tools.Eloquent.Create(&user)
	////fmt.Println(user.ID)
	//
	//tools.Eloquent.First(&user, 1)
	//
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message":  "123",
	//		"userInfo": "user",
	//	})
	//})
	//r.Run()

}
