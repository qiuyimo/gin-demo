package model

import (
	"fmt"
	"github.com/qiuyuhome/gin-demo/global"
	"github.com/qiuyuhome/gin-demo/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Model struct {
	ID         int64     `gorm:"column:id" json:"id"`
	CreatedOn  time.Time `gorm:"column:created_on" json:"created_on"`   //  创建时间
	CreateBy   string    `gorm:"column:create_by" json:"create_by"`     //  创建人
	ModifiedOn time.Time `gorm:"column:modified_on" json:"modified_on"` //  修改时间
	ModifiedBy string    `gorm:"column:modified_by" json:"modified_by"` //  修改人
	DeletedOn  time.Time `gorm:"column:deleted_on" json:"deleted_on"`   //  删除时间
	IsDel      int64     `gorm:"column:is_del" json:"is_del"`           //  是否删除 0为未什出 1为已删除
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dsn := fmt.Sprintf(s, databaseSetting.UserName, databaseSetting.Password, databaseSetting.Host,
		databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	conf := &gorm.Config{}
	if global.ServerSetting.RunMode == "debug" {
		conf = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}
	db, err := gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		return nil, err
	}
	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
