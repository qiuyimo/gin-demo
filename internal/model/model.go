package model

import "time"

type Model struct {
	ID         int64     `gorm:"column:id" json:"id"`
	CreatedOn  time.Time `gorm:"column:created_on" json:"created_on"`   //  创建时间
	CreateBy   string    `gorm:"column:create_by" json:"create_by"`     //  创建人
	ModifiedOn time.Time `gorm:"column:modified_on" json:"modified_on"` //  修改时间
	ModifiedBy string    `gorm:"column:modified_by" json:"modified_by"` //  修改人
	DeletedOn  time.Time `gorm:"column:deleted_on" json:"deleted_on"`   //  删除时间
	IsDel      int64     `gorm:"column:is_del" json:"is_del"`           //  是否删除 0为未什出 1为已删除
}
