package user

import (
	"dbdms/db"
	"dbdms/midware/jwtauth"
	"dbdms/utils"
	"dbdms/utils/regex"
	"fmt"
	"log"
	"strings"
)

// User user model struct
type User struct {
	UserID       int    `gorm:"AUTO_INCREMENT;column:user_id;primary_key"`
	UserName     string `gorm:"type:varchar(32);column:user_name;unique_index;not null" json:"user_name" form:"user_name" binding:"required"`
	UserPhone    string `gorm:"type:varchar(32);column:user_phone;unique_index;not null" json:"user_phone" form:"user_phone" binding:"required"`
	UserPassword string `gorm:"type:varchar(64);column:user_password;not null" json:"user_password" form:"user_password" binding:"required"`
	UserEmail    string `gorm:"type:varchar(64);column:user_email;unique_index" json:"user_email" form:"user_email"`
}

// TableName define table name in database
func (user *User) TableName() string {
	return "sys_users"
}

func (user *User) String() string {
	return fmt.Sprintf("ID: %d, username: %s, email: %s", user.UserID, user.UserName, user.UserEmail)
}

func (user *User) validator() error {
	if ok, err := regex.MatchLetterNumMinAndMax(user.UserName, 4, 6, "user_name"); !ok {
		return err
	}
	//	if ok, err := regex.MatchMediumPassword(user.Password, 6, 13); !ok && user.ID == 0 {
	//		return err
	//	}
	if ok, err := regex.IsPhone(user.UserPhone); !ok {
		return err
	}
	if ok, err := regex.IsEmail(user.UserEmail); !ok {
		return err
	}
	return nil
}

func init() {
	db.SQL.AutoMigrate(&User{})

	err := utils.LoadTokenConfig("./config/token-config.yml")
	if err != nil {
		// utils.ErrorLogger.Errorln("Read Token Config Error: ", err)
		log.Fatal("Read Token Config Error: ", err)
	}
	if len(strings.TrimSpace(utils.GetTokenConfig().SignKey)) > 0 {
		jwtauth.SetSignKey(utils.GetTokenConfig().SignKey)
	}
}
