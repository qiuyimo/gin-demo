package model

type ArticleTags struct {
	ID        int64 `gorm:"column:id" json:"id"`
	ArticleId int64 `gorm:"column:article_id" json:"article_id"` //  文章id
	TagId     int64 `gorm:"column:tag_id" json:"tag_id"`         //  标签id
}

func (a ArticleTags) TableName() string {
	return "article_tags"
}
