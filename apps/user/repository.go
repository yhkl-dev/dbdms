package user

import (
	"dbdms/apps"
	"dbdms/utils"

	"gorm.io/gorm"
)

// Repo user interface inplemented from common interface
type Repo interface {
	apps.RepositoryInterface
}

type userRepo struct {
	db *gorm.DB
}

var userRepoInstance = &userRepo{}

// RepoInterface instance for storage object
func RepoInterface(db *gorm.DB) Repo {
	userRepoInstance.db = db
	return userRepoInstance
}

func (ur *userRepo) Insert(m interface{}) error {
	err := ur.db.Create(m).Error
	return err
}

func (ur *userRepo) Update(m interface{}) error {
	err := ur.db.Save(m).Error
	return err
}
func (ur *userRepo) Delete(m interface{}) error {
	err := ur.db.Delete(m).Error
	return err
}

func (ur *userRepo) FindOne(id int) interface{} {
	var user User
	ur.db.Where("user_id = ?", id).First(&user)
	if user.UserID == 0 {
		return nil
	}
	return &user
}

func (ur *userRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var user User
	ur.db.Where(condition, params).First(&user)
	if user.UserName != "" {
		return &user
	}
	return nil
}
func (ur *userRepo) FindMore(condition string, params ...interface{}) interface{} {
	users := make([]*User, 0)
	ur.db.Where(condition, params).Find(&users)
	return users
}

func (ur *userRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	//	rows := make([]User, 0)
	var rows []User

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			ur.db = ur.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			ur.db = ur.db.Where(k, v)
		}
	}
	// ur.db.Limit(pageSize).Offset((page - 1) * pageSize).Order("login_time desc").Find(&rows).Count(&total)
	ur.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}

}
