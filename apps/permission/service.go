package permission

import helper "dbdms/helpers"

type PermissionService interface {
	GetPage(page int, pageSize int, permission *Permission) *helper.PageBean
	GetByID(id int) *Permission
}

type permissionService struct {
	repo PermissionRepository
}

var permissionServiceIns = &permissionService{}

func PermissionServiceInstance(repo PermissionRepository) PermissionService {
	permissionServiceIns.repo = repo
	return permissionServiceIns
}

func (p *permissionService) GetByID(id int) *Permission {
	if id <= 0 {
		return nil
	}
	perm := p.repo.FindOne(id)
	if perm != nil {
		return perm.(*Permission)
	}
	return nil
}

func (p *permissionService) GetPage(page int, pageSize int, permission *Permission) *helper.PageBean {
	if page == 0 {
		page = 1

	}
	if pageSize == 0 {
		pageSize = 10

	}
	addCons := make(map[string]interface{})
	addCons["is_deleted = ?"] = "0"
	if permission != nil && permission.CodeName != "" {
		addCons["code_name LIKE ?"] = "%" + permission.CodeName + "%"

	}
	pageBean := p.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}