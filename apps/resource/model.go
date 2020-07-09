package resource

import (
	helper "dbdms/helpers"
	"os/user"
	"time"
)

// DBResource struct
type DBResource struct {
	ID          int        `gorm:"AUTO_INCREMENT;primary_key"`
	ConnectName string     `gorm:"type:varchar(32);unique_index;not null" json:"connect_name" form:"connect_name" binding:"required"`
	DBType      int        `gorm:"type:int;comment:'数据库类型: 0: MySQL 1: Oracle: 2: PostgreSQL';default:0" json:"db_type" form:"db_type"`
	IP          string     `gorm:"type:varchar(20);not null" json:"ip" form:"ip"`
	DBUser      string     `gorm:"type:varchar(32);not null" json:"db_user" form:"db_user"`
	DBPassword  string     `gorm:"type:varchar(62);not null" json:"db_password" form:"db_password"`
	Port        int        `gorm:"type:int;not null" json:"port" form:"port"`
	Description string     `gorm:"type:varchar(128);" json:"description" form:"port"`
	IsShare     int        `gorm:"type:int; default:0; comment:'0: no 1:yes'" json:"is_share" form:"is_share"`
	User        *user.User `gorm:"ForeignKey:ID" json:"user" form:"user"`
	CreateAt    time.Time  `gorm:"column:create_at;default:current_timestamp"`
	UpdateAt    time.Time  `gorm:"column:update_at;default:current_timestamp ON update current_timestamp"`
	IsDeleted   int        `gorm:"type:int;default:0" json:"is_deleted"` // 0: no 1: yes
	DeleteAt    *time.Time `gorm:"column:delete_at"`
}

func init() {
	helper.SQL.AutoMigrate(&Resource{})
}
