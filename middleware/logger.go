package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

//记录日志
func Logger() gin.HandlerFunc  {

	logFilePath:="./log"
	logFileName:="app.log"

	//文件路径
	fileName:=path.Join(logFilePath,logFileName)
	// 写入文件
	src,err:=os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
		fmt.Println("err",err)
	}

	logger:=logrus.New()

	// 设置输出
	logger.Out=src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

		writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		startTime:=time.Now()

		c.Next()

		endTime:=time.Now()

		// 执行时间
		allTime:=endTime.Sub(startTime)

		// 请求方式
		reqMethod:=c.Request.Method

		// 请求路由
		reqUrl:=c.Request.RequestURI

		// 状态码
		statusCode:=c.Writer.Status()

		// 请求ip
		clientIP:=c.ClientIP()

		// 换一下日期格式
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
		})

		// 换成json格式
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:"2006-01-02 15:04:05",
		})

		// 日志变成json
		logger.WithFields(logrus.Fields{
			"status_code":statusCode,
			"all_time":allTime,
			"client_ip":clientIP,
			"req_method":reqMethod,
			"req_url":reqUrl,
		}).Info()

		// 日志是字符串
		// logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//  statusCode,
		//  allTime,
		//  clientIP,
		//  reqMethod,
		//  reqUrl,
		// )
	}
}