package model

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)

	if err != nil {
		panic("Mysql数据库连接错误！")
	}
	db.LogMode(true)
	if gin.Mode() == "release" { //发行版不用输出日志
		db.LogMode(false)
	}

	db.SingularTable(true)       // 表明不加s
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //设置最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration() //数据库连接是进行迁移
}
