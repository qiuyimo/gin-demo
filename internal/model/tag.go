package model

type Tags struct {
	*Model
	Name  string `gorm:"column:name" json:"name"`   //  标签名称
	State uint8  `gorm:"column:state" json:"state"` //  状态 0为禁用 1为启用
}

func (t Tags) TableName() string {
	return "tags"
}
