/**
AUTHOR:  zeng_xiao_yu
GITHUB:  https://github.com/zengxiaolou
EMAIL:   zengevent@gmail.com
TIME:    2020/10/26-14:53
**/

package main

import (
	"flag"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"paid_ask/config"
	"paid_ask/model"
	"time"
)

var configFile = flag.String("config", "./setting.yaml", "配置文件路径")

func init()  {
	flag.Parse()
	// 初始化配置
	conf := config.Init(*configFile)

	// gorm 配置
	gormConf := &gorm.Config{}

	// 初始化日志
	if file, err := os.OpenFile(conf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(file)
		if conf.ShowSql {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful: true,
				LogLevel: logger.Info,
			})
		}
	}else {
		logrus.Error(err)
	}

	// 连接数据库
	if err := simple.OpenDB(conf.MySqlUrl, gormConf, 10, 20, model.Models...); err != nil {
		logrus.Error(err)
	}
}

func main()  {
	app.StartOn()
}