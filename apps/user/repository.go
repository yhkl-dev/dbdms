package user

import (
	repository "dbdms/apps/repository"

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
	r.db.Where(condition, params).First(&user)
	if user.UserName != "" {
		return &user
	}
	return nil
}

// find user by id
func (r *userRepository) FindOne(id int) interface{} {
	var user User
	r.db.Where("id = ?", id).First(&user)
	if user.UserName == "" {
		return nil
	}
	return &user
}

// 条件查询返回多值
func (r *userRepository) FindMore(condition string, params ...interface{}) interface{} {
	users := make([]*User, 0)
	r.db.Where(condition, params).Find(&users)
	return users
}
