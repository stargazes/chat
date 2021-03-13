package tools

import "github.com/gomodule/redigo/redis"

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
