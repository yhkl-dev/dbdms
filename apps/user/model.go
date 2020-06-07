package user

import (
	helper "dbdms/helpers"
	"time"
)

type CrudTime struct {
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt *time.Time `sql:"index"`
}

type User struct {
	ID         int    `gorm:"AUTO_INCREMENT;primary_key"`
	UserName   string `gorm:"type:varchar(32);unique_index;not null" form:"username" binding:"required"`
	Password   string `gorm:"type:varchar(64);not null" json:"-" form:"password" binding:"required"`
	Phone      string `gorm:"type:varchar(11);unique" form:"phone" binding:"required"`
	Email      string `gorm:"type:varchar(64)" form:"email"`
	LoginCount int
	Status     int
	CrudTime
	LoginTime time.Time `gorm:"default:null"`
	Role      *Role     `gorm:"foreignkey:RoleId;save_associations:false"`
	RoleID    *string   `gorm:"type:int" form:"role_id"`
}

func init() {
	helper.SQL.AutoMigrate(&User{})
	helper.SQL.Model(&User{}).AddForeignKey("role_id", "role(id)", "no action", "no action")
}
