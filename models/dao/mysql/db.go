package mysql

import (
	"github.com/ShaoZeMing/going/core"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitMysql(app *core.App) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	var err error
	dsn := app.Config.Mysql.Dsn
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[gstore_db] mysql fail, err=%s", err)
		panic(err)
	}

	gormDB, err := DB.DB()
	if err != nil {
		log.Printf("[gstore_db] mysql fail, err=%s", err)
		panic(err)
	}
	gormDB.SetMaxIdleConns(3)
	gormDB.SetMaxOpenConns(10)
	gormDB.SetConnMaxLifetime(time.Hour * time.Duration(1000))

	err = gormDB.Ping()
	if err != nil {
		log.Printf("[gstore_db] mysql fail, err:%s", err.Error())
		panic(err)
	}
	log.Printf("[gstore_db] mysql success")
}
