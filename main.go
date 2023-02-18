package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/migration", func(c *gin.Context) {
		type Common struct {
			ID         uint      `gorm:"primarykey"`
			CreatedOn  time.Time `gorm:"comment:创建时间"`
			CreateBy   string    `gorm:"size:100;default:;comment:创建人"`
			ModifiedOn time.Time `gorm:"comment:修改时间"`
			ModifiedBy string    `gorm:"size:100;default:;comment:修改人"`
			DeletedOn  time.Time `gorm:"comment:删除时间"`
			IsDel      bool      `gorm:"default:0;comment:是否删除 0为未什出 1为已删除"`
		}

		type Tag struct {
			Common
			Name  string `gorm:"default:;comment:标签名称"`
			State uint8  `gorm:"default:1;comment:状态 0为禁用 1为启用"`
		}

		type Article struct {
			Common
			Title         string `gorm:"size:100;default:;comment:文章标题"`
			Desc          string `gorm:"size:255;default:;comment:文章简述"`
			CoverImageUrl string `gorm:"size:255;default:;comment:封面图片地址"`
			Content       string `gorm:"type:text;comment:文章内容"`
			State         uint8  `gorm:"default:1;comment:状态 0为禁用 1为启用"`
		}

		type ArticleTag struct {
			ID        uint `gorm:"primarykey"`
			ArticleID uint `gorm:"not null;comment:文章ID"`
			TagID     uint `gorm:"not null;default:0;comment:标签ID"`
		}

		// 连接数据库
		dsn := "root:root@tcp(127.0.0.1:3306)/gin-demo?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			c.JSON(500, gin.H{"message": "connect db error: " + err.Error()})
			panic("failed to connect database")
		}

		// 数据迁移，给已经迁移过的结构体 Product 添加成员，再次迁移，表 products 会新增字段，但删除结构体的成员，表不会删除字段
		if e := db.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{}); e != nil {
			c.JSON(500, gin.H{"message": "migrate error: " + err.Error()})
			panic(e.Error())
		}

		c.JSON(200, gin.H{"message": "migrate completed"})
	})

	r.Run()
}
