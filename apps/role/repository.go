package role

import (
	"dbdms/apps"
	"dbdms/utils"
	"gorm.io/gorm"
)

// Repo role interface inplemented from common interface
type Repo interface {
	apps.RepositoryInterface
	AddUserRole(roleId int, userID int) error
}

type roleRepo struct {
	db *gorm.DB
}

type userRoleMappingRepo struct {
	db *gorm.DB
}

var roleRepoInstance = &roleRepo{}
var userRoleRepoInstance = &userRoleMappingRepo{}

// RepoInterface instance for storage object
func RepoInterface(db *gorm.DB) Repo {
	roleRepoInstance.db = db
	userRoleRepoInstance.db = db
	return roleRepoInstance
}

func (urm *userRoleMappingRepo) Delete(roleID int) error {
	err := urm.db.Where("role_id = ?", roleID).Delete(UserRoleMapping{}).Error
	return err
}

func (rp *roleRepo) Insert(m interface{}) error {
	err := rp.db.Create(m).Error
	return err
}

func (rp *roleRepo) AddUserRole(roleId int, userID int) error {
	mapping := &UserRoleMapping{RoleID: roleId, UserID: userID}
	err := userRoleRepoInstance.db.Create(mapping).Error
	return err
}

func (rp *roleRepo) Update(m interface{}) error {
	err := rp.db.Save(m).Error
	return err
}

func (rp *roleRepo) Delete(m interface{}) error {
	err := rp.db.Delete(m).Error
	if err != nil {
		return err
	}
	// delete role from user role mapping table
	var userRoleMappingObject userRoleMappingRepo
	err = userRoleMappingObject.Delete(m.(Role).RoleID)
	return err
}

func (rp *roleRepo) FindMore(condition string, params ...interface{}) interface{} {
	roles := make([]*Role, 0)
	rp.db.Where(condition, params).Find(&roles)
	return roles
}

func (rp *roleRepo) FindOne(id int) interface{} {
	var role Role
	err := rp.db.Where("role_id = ?", id).First(&role).Error
	if role.RoleID == 0 || err != nil {
		return nil
	}
	return &role
}

func (rp *roleRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	//	rows := make([]User, 0)
	var rows []Role

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			rp.db = rp.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			rp.db = rp.db.Where(k, v)
		}
	}
	rp.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}

func (rp *roleRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var role Role
	rp.db.Where(condition, params).First(&role)
	if role.RoleName != "" {
		return &role
	}
	return nil
}
