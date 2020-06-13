package user

import (
	helper "dbdms/helpers"
	regex "dbdms/helpers/regex"
	"time"
)

// CrudTime the struct for User
type CrudTime struct {
	CreateAt time.Time  `gorm:"column:create_at:default:null"`
	UpdateAt time.Time  `gorm:"column:update_at;default:'1970-01-01'"`
	DeleteAt *time.Time `sql:"index" gorm:"column:delete_at"`
}

// User the user struct
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;primary_key"`
	UserName string `gorm:"type:varchar(32);unique_index;not null" json:"username" form:"username" binding:"required"`
	Password string `gorm:"type:varchar(64);not null" json:"password" form:"password" binding:"required"`
	Phone    string `gorm:"type:varchar(11);unique" form:"phone" binding:"required"`
	Email    string `gorm:"type:varchar(64)" form:"email"`
	// 0: no 1: yes
	IsDeleted int `gorm:"type:int;default:0" json:"is_deleted"`
	Status    int
	CrudTime
	LoginTime time.Time `gorm:"default:null"`
	//	Role      *Role     `gorm:"foreignkey:RoleId;save_associations:false"`
	//	RoleID    *string   `gorm:"type:int" form:"role_id"`
}

// Validator user column validator
func (user *User) Validator() error {
	if ok, err := regex.MatchLetterNumMinAndMax(user.UserName, 4, 6, "username"); !ok {
		return err
	}
	//	if ok, err := regex.MatchMediumPassword(user.Password, 6, 13); !ok && user.ID == 0 {
	//		return err
	//	}
	if ok, err := regex.IsPhone(user.Phone); !ok {
		return err
	}
	if ok, err := regex.IsEmail(user.Email); !ok {
		return err
	}
	return nil
}

func init() {
	helper.SQL.AutoMigrate(&User{})
	helper.SQL.Model(&User{}).AddForeignKey("role_id", "role(id)", "no action", "no action")
}
