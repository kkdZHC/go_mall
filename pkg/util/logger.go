package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutPutFile()
	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	logger.Out = src                   //设置输出
	logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	//规范日志路径
	logFilePath := ""
	dir, err := os.Getwd() //获取工作目录
	if err == nil {
		logFilePath = dir + "/logs/"
	}
	//若不存在就创建
	_, err = os.Stat(logFilePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logFilePath, 0777)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//规范日志名
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	//若不存在就创建
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		err = os.MkdirAll(fileName, 0777)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入日志文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return src, nil
}
