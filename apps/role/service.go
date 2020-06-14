package role

import (
	helper "dbdms/helpers"
	"errors"
	"fmt"
	"time"
)

// RoleService role service interface
type RoleService interface {
	SaveOrUpdate(role *Role) error
	GetByID(id int) *Role
	GetByRoleName(rolename string) *Role
	GetPage(page int, pageSize int, role *Role) *helper.PageBean
	DeleteByID(id int) error
}

var roleServiceInstance = &roleService{}

// RoleServiceInstance role service instance
func RoleServiceInstance(repo RoleRepository) RoleService {
	roleServiceInstance.repo = repo
	return roleServiceInstance
}

type roleService struct {
	repo RoleRepository
}

func (rs *roleService) SaveOrUpdate(role *Role) error {
	if role == nil {
		return errors.New(helper.StatusText(helper.SaveObjIsNil))
	}
	roleByName := rs.GetByRoleName(role.RoleName)
	if role.ID == 0 {
		if roleByName != nil && roleByName.ID != 0 {
			return errors.New(helper.StatusText(helper.ExistSameNameError))
		}
		return rs.repo.Insert(role)
	}
	persist := rs.GetByID(role.ID)
	if persist == nil || persist.ID == 0 {
		return errors.New(helper.StatusText(helper.UpdateObjIsNil))
	}
	if roleByName != nil && roleByName.ID != role.ID {
		return errors.New(helper.StatusText(helper.ExistSameNameError))
	}
	role.CreateAt = persist.CreateAt
	role.UpdateAt = time.Now()

	return rs.repo.Update(role)
}

func (rs *roleService) GetByID(id int) *Role {
	if id <= 0 {
		return nil
	}
	role := rs.repo.FindOne(id)
	if role != nil {
		fmt.Println(1111111112)
	}
	return nil
}

func (rs *roleService) GetByRoleName(rolename string) *Role {
	role := rs.repo.FindSingle("role_name = ?", rolename)
	if role != nil {
		return role.(*Role)
	}
	return nil
}

func (rs *roleService) DeleteByID(id int) error {
	roleS := rs.repo.FindOne(id)
	if roleS == nil {
		return errors.New(helper.StatusText(helper.DeleteStatusErr))
	}
	role := roleS.(*Role)
	if role.ID == 0 {
		return errors.New(helper.StatusText(helper.DeleteStatusErr))
	}
	role.IsDeleted = 1
	deleteTime := time.Now()
	role.DeleteAt = &deleteTime
	err := rs.repo.Delete(role)
	if err != nil {
		return err
	}
	return nil
}

func (rs *roleService) GetPage(page int, pageSize int, role *Role) *helper.PageBean {
	if page == 0 {
		page = 1

	}
	if pageSize == 0 {
		pageSize = 10

	}
	addCons := make(map[string]interface{})
	addCons["is_deleted = ?"] = "0"
	if role != nil && role.RoleName != "" {
		addCons["role_name LIKE ?"] = "%" + role.RoleName + "%"

	}
	pageBean := rs.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}
