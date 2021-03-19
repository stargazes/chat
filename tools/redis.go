package tools

import (
    "github.com/gomodule/redigo/redis"
)

type Connection redis.Conn

var (
    RedisPool *redis.Pool
    redisHost = "localhost:6379"
    password  = ""
)

func ConnForRedis()  {

    RedisPool = &redis.Pool{
        MaxIdle:   3, //最大空闲连接数
        MaxActive: 8, //最大激活连接数
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", redisHost, redis.DialPassword(password))
            if err != nil {
                return nil, err
            }
            return c, nil
        },
    }

}

func init() {
    ConnForRedis()

}

//设置键值
func SetString(key string,str string,expire int) bool {
    _, err := RedisPool.Get().Do("set",key,string(str))

    if err !=nil {
        return false
    }
    return true

}
//获取键值
func GetString(key string) interface{}  {

    str,err :=redis.String(RedisPool.Get().Do("get",key))

    if err != nil {
        return ""
    }
    return str
}

/**
加入群组
 */
func AddChatRoom(roomHash string ,userId int64) bool {
    _,err := RedisPool.Get().Do("sadd",roomHash,userId)
    if err !=nil {
        return false
    }
    return true
}

/**
退出群组
 */
func EXitChatRoom(roomHash string ,userId int64) bool{

    _,err := RedisPool.Get().Do("srem",roomHash,userId)
    if err !=nil {
        return false
    }
    return true
}
