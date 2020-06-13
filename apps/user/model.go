package user

import (
	helper "dbdms/helpers"
	regex "dbdms/helpers/regex"
	"time"
)

// User the user struct
type User struct {
	ID        int        `gorm:"AUTO_INCREMENT;primary_key"`
	UserName  string     `gorm:"type:varchar(32);unique_index;not null" json:"username" form:"username" binding:"required"`
	Password  string     `gorm:"type:varchar(64);not null" json:"password" form:"password" binding:"required"`
	Phone     string     `gorm:"type:varchar(11);unique" form:"phone" binding:"required"`
	Email     string     `gorm:"type:varchar(64)" form:"email"`
	IsDeleted int        `gorm:"type:int;default:0" json:"is_deleted"` // 0: no 1: yes
	Status    int        `gorm:"column:status"`
	CreateAt  time.Time  `gorm:"column:create_at;default:current_timestamp"`
	UpdateAt  time.Time  `gorm:"column:update_at;default:current_timestamp ON update current_timestamp"`
	DeleteAt  *time.Time `gorm:"column:delete_at"`
	LoginTime *time.Time `gorm:"column:login_time"`
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
