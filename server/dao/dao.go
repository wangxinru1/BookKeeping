package dao

import (
	"golang-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Register(user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:yuanshuai520@tcp(127.0.0.1:3306)/golang_blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接数据库失败，err: ", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
}

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}
