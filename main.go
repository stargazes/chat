package main

import (
    "chat/models"
    "chat/tools"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gomodule/redigo/redis"
    "net/http"
)

func main()  {

    //tools.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}) //自动生成表格
    r:=gin.Default()
    conn:=tools.RedisPool.Get()

    if conn == nil {
      fmt.Println("获取连接失败")
    }

    //redis操作例子
    _,err:=conn.Do("set","username",string("lyh"))

    if err!= nil{
      fmt.Println("设置值失败")
    }

    value,err:=redis.String(conn.Do("get","username"))
    fmt.Println(value)

    if err!=nil{
      fmt.Println("获取username失败")
    }

    //mysql操作例子
    var user models.User
    //user.Username="lyh"
    //user.Password="123456"
    //tools.Eloquent.Create(&user)
    //fmt.Println(user.ID)

    tools.Eloquent.First(&user,1)


    r.GET("/", func(c *gin.Context) {
     c.JSON(http.StatusOK,gin.H{
         "message":value,
         "userInfo":user,
     })
    })
    r.Run()

}
