/*
	Copyright 2019 The Yingxi.company Authors. All rights reserved.
	Go
	https://jasperxu.github.io/gorm-zh/database.html#dbc
	Model
*/
package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
)

var DB *Database

// 定义数据库Struct
type Database struct {
	Master *gorm.DB
	Slave  *gorm.DB
}

// 定义自定义Log
type MyLogger struct {

}

// 初始化数据库连接
func (db *Database) Init(conf map[string]string) {
	DB = &Database{
		Master: GetMasterDB(conf["selection"], conf["master"]),
		Slave:  GetSlaveDB(conf["selection"], conf["slave"]),
	}
}

// 打开数据库
func openDb(selection, host string) *gorm.DB {
	db, err := gorm.Open(selection, host)
	if err != nil {
		log.Errorf(err, "[Web] %s-%s", selection, host)
	}
	// 详细日志
	logger := &MyLogger{}
	db.LogMode(true)
	db.SetLogger(logger)
	// SetMaxOpenConns用于设置最大打开的连接数
	// SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db
}

// 主库
func GetMasterDB(selection, master string) *gorm.DB {
	return openDb(selection, master)
}

// 从库
func GetSlaveDB(selection, slave string) *gorm.DB {
	return openDb(selection, slave)
}

// 关闭连接
func (db *Database) Close() {
	DB.Master.Close()
	DB.Slave.Close()
}

// 自定义日志
func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level = values[0]
		source = values[1]
		runtime = values[2]
	)

	if level == "sql" {
		sql := values[3].(string)
		log.Infof("[Web] %s | %s | %s %s ", level, source, runtime, sql)
	} else {
		log.Infof("[Web] %v ", values)
	}
}