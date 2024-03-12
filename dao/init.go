package dao

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  // string字段默认长度
		DisableDatetimePrecision:  true, //禁止datatime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true, //禁止重命名索引， 即禁止把索引删了重建，5.7
		DontSupportRenameColumn:   true, //禁止重命名列，8之前不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		//日志
		Logger: ormLogger,
		//命名策略，单数化不加s
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  //设置空闲连接池
	sqlDB.SetMaxOpenConns(100)                 //最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //限制连接的生命周期，连接超时后，会在需要时惰性回收复用
	_db = db

	//主从配置
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                    //负载均衡
		}))
	migration()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
