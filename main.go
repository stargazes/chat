package main

import (
    "chat/impl"
    "github.com/gorilla/websocket"
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

func wsHandle(w http.ResponseWriter,r *http.Request)  {
    var(
        wsConn *websocket.Conn
        err error
        conn*impl.Connection
        data []byte
    )
    //完成ws协议的握手操作
    //upgrade:websocket
    if wsConn,err = upgrader.Upgrade(w,r,nil);err !=nil{
        return
    }
    if conn,err=impl.InitConnection(wsConn);err!=nil{
        log.Println("握手失败")
        goto ERR
    }
    //启动线程
    go func() {
        var (err error)
        for {
            if err =conn.WriteMessage([]byte("heartbeat"));err!=nil{
                return
            }
            time.Sleep(1*time.Second)
        }
    }()
    for {
        if data,err=conn.ReadMessage();err!=nil{
            goto ERR
        }
        if err = conn.WriteMessage(data);err!=nil {
            goto ERR
        }
    }
    ERR:
        conn.Close()

}

func main() {
    http.HandleFunc("/",wsHandle)
    //http.ListenAndServe("0.0.0.0:9990",nil)
    if err := http.ListenAndServe(":10001", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }else {
        log.Println("websocket创建成功")
    }

    //tools.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}) //自动生成表格
    //r := gin.Default()
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
    //r.GET("/", func(c *gin.Context) {
    //    c.JSON(http.StatusOK, gin.H{
    //        "message":  value,
    //        "userInfo": user,
    //    })
    //})
    //r.Run()

}
