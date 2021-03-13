package tools

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Eloquent *gorm.DB

func init(){

    dsn:="root:123456@tcp(127.0.0.1:3306)/chat?" +
        "charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms"

    Eloquent,_ = gorm.Open(mysql.Open(dsn),&gorm.Config{})

}
