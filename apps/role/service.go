package role

import (
	"dbdms/utils"
	"errors"
)

// Service role service interface
type Service interface {
	ListAllRoles() []*Role
	GetByID(id int) *Role
	GetPage(page int, pageSize int, role *Role) *utils.PageBean
	DeleteRoleByID(id int) error
	SaveOrUpdate(role *Role) error
	AddRoleToUser(roleID int, userID int) error
}

type roleService struct {
	repo Repo
}

type roleUserMappingService struct {
	repo Repo
}

var roleServiceIns = &roleService{}
var roleUserMappingServiceIns = &roleUserMappingService{}

// ServiceInstance 获取 roleService 实例
func ServiceInstance(repo Repo) Service {
	roleServiceIns.repo = repo
	roleUserMappingServiceIns.repo = repo
	return roleServiceIns
}

func (rs *roleService) ListAllRoles() []*Role {
	roles := rs.repo.FindMore("1=1").([]*Role)
	return roles
}

func (rs *roleService) AddRoleToUser(roleID int, userID int) error {
	err := rs.repo.AddUserRole(roleID, userID)
	return err
}

func (rs *roleService) GetByID(id int) *Role {
	if id <= 0 {
		return nil
	}
	role := rs.repo.FindOne(id)
	if role != nil {
		return role.(*Role)
	}
	return nil
}

func (rs *roleService) DeleteRoleByID(id int) error {
	role := rs.repo.FindOne(id)
	if role == nil {
		return errors.New(utils.StatusText(utils.DeleteObjIsNil))
	}
	return rs.repo.Delete(role.(*Role))
}

func (rs *roleService) GetPage(page int, pageSize int, role *Role) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	addCons := make(map[string]interface{})
	if role != nil && role.RoleName != "" {
		addCons["role_name LIKE ?"] = "%" + role.RoleName + "%"
	}

	pageBean := rs.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (rs *roleService) SaveOrUpdate(role *Role) error {
	if role == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	if role.RoleID == 0 {
		return rs.repo.Insert(role)
	}
	return rs.repo.Update(role)
}
