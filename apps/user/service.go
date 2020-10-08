package user

import (
	"dbdms/utils"
	"errors"
)

// Service user service instance
type Service interface {
	GetAll() []*User
	GetUserByName(username string) *User
	GetUserByPhone(phone string) *User
	GetByID(id int) *User
	GetPage(page int, pageSize int, user *User) *utils.PageBean
	DeleteByID(id int) error
	SaveOrUpdate(user *User) error
}

var userServiceIns = &userService{}

// ServiceInstance 获取 userService 实例
func ServiceInstance(repo Repo) Service {
	userServiceIns.repo = repo
	return userServiceIns
}

type userService struct {
	repo Repo
}

func (us *userService) GetAll() []*User {
	users := us.repo.FindMore("1=1").([]*User)
	return users
}

func (us *userService) GetUserByName(username string) *User {
	user := us.repo.FindSingle("user_name = ?", username)
	if user != nil {
		return user.(*User)
	}
	return nil
}

func (us *userService) GetUserByPhone(phone string) *User {
	user := us.repo.FindSingle("user_phone = ?", phone)
	if user != nil {
		return user.(*User)
	}
	return nil
}

func (us *userService) GetByID(id int) *User {
	if id <= 0 {
		return nil
	}
	user := us.repo.FindOne(id)
	if user != nil {
		return user.(*User)
	}
	return nil
}

func (us *userService) SaveOrUpdate(user *User) error {
	if user == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	userByName := us.GetUserByName(user.UserName)
	userByPhone := us.GetUserByPhone(user.UserPhone)
	if user.UserID == 0 {
		if userByName != nil && userByName.UserID != 0 {
			return errors.New(utils.StatusText(utils.ExistSameNameError))
		}
		if userByPhone != nil && userByPhone.UserID != 0 {
			return errors.New(utils.StatusText(utils.ExistSamePhoneError))
		}

		user.UserPassword = utils.SHA256(user.UserPassword)
		return us.repo.Insert(user)
	}
	persist := us.GetByID(user.UserID)
	if persist == nil || persist.UserID == 0 {
		return errors.New(utils.StatusText(utils.UpdateObjIsNil))
	}
	if userByName != nil && userByName.UserID != user.UserID {
		return errors.New(utils.StatusText(utils.ExistSameNameError))
	}
	if userByPhone != nil && userByPhone.UserID != user.UserID {
		return errors.New(utils.StatusText(utils.ExistSamePhoneError))
	}
	if persist.UserPassword != user.UserPassword {
		user.UserPassword = utils.SHA256(user.UserPassword)
	} else {
		user.UserPassword = persist.UserPassword
	}
	return us.repo.Update(user)
}

func (us *userService) DeleteByID(id int) error {
	user := us.repo.FindOne(id).(*User)
	if user == nil || user.UserID == 0 {
		return errors.New(utils.StatusText(utils.DeleteObjIsNil))
	}
	return us.repo.Update(user)
}

func (us *userService) GetPage(page int, pageSize int, user *User) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	addCons := make(map[string]interface{})
	// addCons["is_deleted = ?"] = "0"
	if user != nil && user.UserName != "" {
		addCons["user_name LIKE ?"] = "%" + user.UserName + "%"
	}
	if user != nil && user.UserPhone != "" {
		addCons["user_phone LIKE ?"] = user.UserPhone + "%"
	}
	if user != nil && user.UserEmail != "" {
		addCons["user_email LIKE ?"] = user.UserEmail + "%"
	}
	pageBean := us.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}
