package user

import (
	"dbdms/apps/role"
	helper "dbdms/helpers"
	regex "dbdms/helpers/regex"
	"time"
)

// User the user struct
type User struct {
	ID        int         `gorm:"AUTO_INCREMENT;primary_key"`
	UserName  string      `gorm:"type:varchar(32);unique_index;not null" json:"username" form:"username" binding:"required"`
	Password  string      `gorm:"type:varchar(64);not null" json:"password" form:"password" binding:"required"`
	Phone     string      `gorm:"type:varchar(11);unique" form:"phone" binding:"required"`
	Email     string      `gorm:"type:varchar(64)" form:"email"`
	IsDeleted int         `gorm:"type:int;default:0" json:"is_deleted"` // 0: no 1: yes
	Status    int         `gorm:"column:status"`
	CreateAt  time.Time   `gorm:"column:create_at;default:current_timestamp"`
	UpdateAt  time.Time   `gorm:"column:update_at;default:current_timestamp ON update current_timestamp"`
	DeleteAt  *time.Time  `gorm:"column:delete_at"`
	LoginTime *time.Time  `gorm:"column:login_time"`
	RoleList  []int       `gorm:"-" json:"roles" form:"roles"`
	Roles     []role.Role `gorm:"many2many:user_role_mapping"  `
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

func (user *User) RegisterPermission() map[string]map[string]string {
	var permList = make(map[string]map[string]string)
	permList["can_view_users"] = map[string]string{"ModelName": "User", "PermissionName": "GET", "CodeName": "/api/v1/user"}
	//	permList["can_add_user"] = "/api/v1/user"
	//	permList["can_update_user"] = "/api/v1/user/(.d+)"
	//	permList["can_delete_user"] = "/api/v1/user/(.d+)"
	return permList
}

func init() {
	helper.SQL.AutoMigrate(&User{})
}
