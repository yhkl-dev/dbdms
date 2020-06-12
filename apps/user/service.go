package user

import (
	helper "dbdms/helpers"
	"errors"
)

type UserService interface {
	GetAll() []*User
	GetUserByName(username string) *User
	GetUserByPhone(phone string) *User
	GetByID(id int) *User
	SaveOrUpdate(user *User) error
}

var userServiceIns = &userService{}

// 获取 userService 实例
func UserServiceInstance(repo UserRepository) UserService {
	userServiceIns.repo = repo
	return userServiceIns
}

type userService struct {
	repo UserRepository
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
	user := us.repo.FindSingle("phone = ?", phone)
	if user != nil {
		return user.(*User)
	}
	return nil
}

func (us *userService) GetByID(id int) *User {
	if id <= 0 {
		return nil
	}
	user := us.repo.FindOne(id).(*User)
	return user
}

func (us *userService) SaveOrUpdate(user *User) error {
	if user == nil {
		return errors.New(helper.StatusText(helper.SaveObjIsNil))
	}
	userByName := us.GetUserByName(user.UserName)
	userByPhone := us.GetUserByPhone(user.Phone)
	if user.ID == 0 {
		if userByName != nil && userByName.ID != 0 {
			return errors.New(helper.StatusText(helper.ExistSameNameError))
		}
		if userByPhone != nil && userByPhone.ID != 0 {
			return errors.New(helper.StatusText(helper.ExistSamePhoneError))
		}
		user.Password = helper.SHA256(user.Password)
		return us.repo.Insert(user)
	}
	persist := us.GetByID(user.ID)
	if persist == nil || persist.ID == 0 {
		return errors.New(helper.StatusText(helper.UpdateObjIsNil))
	}
	if userByName != nil && userByName.ID != user.ID {
		return errors.New(helper.StatusText(helper.ExistSameNameError))
	}
	if userByPhone != nil && userByPhone.ID != user.ID {
		return errors.New(helper.StatusText(helper.ExistSamePhoneError))
	}
	user.Password = persist.Password
	return us.repo.Update(user)
}

func (us *userService) DeleteByID(id int) error {
	user := us.repo.FindOne(id).(*User)
	if user == nil || user.ID == 0 {
		return errors.New(helper.StatusText(helper.DeleteObjIsNil))
	}
	err := us.repo.Delete(user)
	return err
}

func (us *userService) GetPage(page int, pageSize int, user *User) *helper.PageBean {
	addCons := make(map[string]interface{})
	if user != nil && user.UserName != "" {
		addCons["user_name LIKE ?"] = user.UserName + "?"
	}
	if user != nil && user.Phone != "" {
		addCons["phone LIKE ?"] = user.Phone + "?"
	}
	if user != nil && user.Email != "" {
		addCons["email LIKE ?"] = user.Email + "?"
	}
	pageBean := us.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}
