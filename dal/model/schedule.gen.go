// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSchedule = "tblSchedule"

// Schedule mapped from table <tblSchedule>
type Schedule struct {
	ID        int64  `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`              // 自增ID
	UserID    int64  `gorm:"column:user_id;type:int unsigned;not null;index:userID,priority:1" json:"user_id"` // 用户ID
	Title     string `gorm:"column:title;type:varchar(255);not null" json:"title"`                             // 标题
	Content   string `gorm:"column:content;type:varchar(255);not null" json:"content"`                         // 内容
	BeginTime int64  `gorm:"column:begin_time;type:int;not null" json:"begin_time"`                            // 开始时间
	EndTime   int64  `gorm:"column:end_time;type:int;not null" json:"end_time"`                                // 结束时间
	CreatedAt int64  `gorm:"column:created_at;type:int;not null" json:"created_at"`                            // 创建时间
	UpdatedAt int64  `gorm:"1111" json:"updated_at"`                                                           // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int;not null" json:"deleted_at"`                            // 删除时间
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return TableNameSchedule
}