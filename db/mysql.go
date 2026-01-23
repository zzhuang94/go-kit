package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"xorm.io/xorm"
)

type MysqlCfg struct {
	DSN     string `json:"dsn"`
	MaxIdle int    `json:"max_idle"`
	MaxOpen int    `json:"max_open"`
	Log     bool   `json:"log"`
}

func (c *MysqlCfg) ConnGorm() (*gorm.DB, error) {
	logMode := logger.Silent
	if c.Log {
		logMode = logger.Info
	}
	opts := &gorm.Config{Logger: logger.Default.LogMode(logMode)}
	db, err := gorm.Open(mysql.Open(c.DSN), opts)
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdle)
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	return db, nil
}

func (c *MysqlCfg) ConnXorm() (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", c.DSN)
	if err != nil {
		return nil, err
	}
	db.ShowSQL(c.Log)
	db.SetMaxIdleConns(c.MaxIdle)
	db.SetMaxOpenConns(c.MaxOpen)
	return db, nil
}
