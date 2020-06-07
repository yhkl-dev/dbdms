package user

type UserService interface {
	GetAll() []*User
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
	user := us.repo.FindOne("user_name = ?", username)
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
