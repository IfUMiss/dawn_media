package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"dawn_media/conf"
)

var db *gorm.DB

/**
初始化数据库
包括生成外键和自动生成数据库
 */
func Init() {
	if db != nil {
		return
	}
	uri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", conf.C().Username, conf.C().Password, conf.C().Database)
	t, err := gorm.Open("mysql", uri)
	if err != nil {
		panic(err)
	}
	t.LogMode(true)
	t.AutoMigrate(
		&User{},
		&UserRecord{},
		&Media{},
		&Category{},
		&Star{},
		&Comment{},
		&MediaAttribute{},
	)
	t.Table("media_categories").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	t.Table("media_categories").AddForeignKey("media_id", "media(id)", "CASCADE", "CASCADE")
	t.Table("user_records").AddForeignKey("media_id", "media(id)", "CASCADE", "CASCADE")
	t.Table("user_records").AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	t.Table("stars").AddForeignKey("media_id", "media(id)", "CASCADE", "CASCADE")
	t.Table("stars").AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	t.Table("comments").AddForeignKey("media_id", "media(id)", "CASCADE", "CASCADE")
	t.Table("comments").AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	t.Table("media_attributes").AddForeignKey("media_id", "media(id)", "CASCADE", "CASCADE")
	db = t
}

func DB() *gorm.DB {
	if db == nil {
		panic("gorm db is not exist")
	}
	return db
}
