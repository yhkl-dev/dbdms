package role

import (
	repository "dbdms/apps/repository"
	helper "dbdms/helpers"
	"errors"

	"github.com/jinzhu/gorm"
)

// RoleRepository role repository
type RoleRepository interface {
	repository.Repository
}

type roleRepository struct {
	db *gorm.DB
}

var roleRepoIns = &roleRepository{}

// RoleRepositoryIntance instance for storage object
func RoleRepositoryIntance(db *gorm.DB) RoleRepository {
	roleRepoIns.db = db
	return roleRepoIns
}

func (r *roleRepository) Insert(role interface{}) error {
	err := r.db.Create(role).Error
	return err
}

func (r *roleRepository) Update(role interface{}) error {
	//	r.db.Model(&role).Association("Permissions").Clear()
	err := r.db.Model(role.(*Role)).Association("Permissions").Replace(role.(*Role).Permissions).Error
	if err != nil {
		return err
	}
	err = r.db.Save(role).Error
	return err
}

func (r *roleRepository) Delete(role interface{}) error {
	var count int
	r.db.Raw("select count(role_id) from user_role_mapping where role_id = ?", role.(*Role).ID).Row().Scan(&count)
	if count > 0 {
		return errors.New("Role has been used by user")
	}

	err := r.db.Save(role).Error
	return err
}

// find role by name
func (r *roleRepository) FindSingle(condition string, params ...interface{}) interface{} {
	var role Role
	r.db.Where(condition, params).First(&role)
	if role.RoleName != "" {
		return &role
	}
	return nil
}

// find role by id
func (r *roleRepository) FindOne(id int) interface{} {
	var role Role
	err := r.db.Where("id = ? and is_deleted = 0", id).First(&role).Error
	if err != nil {
		return nil
	}
	if role.ID == 0 {
		return nil
	}
	return &role
}

// 条件查询返回多值
func (r *roleRepository) FindMore(condition string, params ...interface{}) interface{} {
	roles := make([]*Role, 0)
	r.db.Where(condition, params).Find(&roles)
	return roles
}

func (r *roleRepository) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *helper.PageBean) {
	total := 0
	rows := make([]Role, 0)
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
	r.db.Preload("Permissions").Limit(pageSize).Offset((page - 1) * pageSize).Order("update_at desc").Find(&rows).Count(&total)
	return &helper.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
