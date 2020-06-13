package user

import (
	repository "dbdms/apps/repository"
	"dbdms/apps/role"
	helper "dbdms/helpers"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	repository.Repository
}

type userRepository struct {
	db *gorm.DB
}

var userRepoIns = &userRepository{}

// UserRepositoryInterface instance for storage object
func UserRepositoryInterface(db *gorm.DB) UserRepository {
	userRepoIns.db = db
	return userRepoIns
}

func (r *userRepository) Insert(user interface{}) error {
	err := r.db.Create(user).Error
	return err
}

func (r *userRepository) Update(user interface{}) error {
	err := r.db.Save(user).Error
	return err
}

func (r *userRepository) Delete(user interface{}) error {
	err := r.db.Delete(user).Error
	return err
}

// find user by name
func (r *userRepository) FindSingle(condition string, params ...interface{}) interface{} {
	var user User
	r.db.Preload("Roles").Where(condition, params).First(&user)
	if user.UserName != "" {
		return &user
	}
	return nil
}

// find user by id
func (r *userRepository) FindOne(id int) interface{} {
	var user User
	var roles []role.Role
	r.db.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return nil
	}
	r.db.Model(&user).Association("Roles").Find(&roles)
	user.Roles = roles
	return &user
}

// 条件查询返回多值
func (r *userRepository) FindMore(condition string, params ...interface{}) interface{} {
	users := make([]*User, 0)
	r.db.Preload("Roles").Where(condition, params).Find(&users)
	return users
}

func (r *userRepository) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *helper.PageBean) {
	total := 0
	//	rows := make([]User, 0)
	var rows []User

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			r.db = r.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			r.db = r.db.Where(k, v)
		}
	}
	r.db.Preload("Roles").Limit(pageSize).Offset((page - 1) * pageSize).Order("login_time desc").Find(&rows).Count(&total)
	return &helper.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
