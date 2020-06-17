package permission

import (
	"dbdms/apps/repository"
	helper "dbdms/helpers"

	"github.com/jinzhu/gorm"
)

type PermissionRepository interface {
	repository.Repository
}

type permissionRepository struct {
	db *gorm.DB
}

var permissionRepoIns = &permissionRepository{}

func PermissionRepositoryInterface(db *gorm.DB) PermissionRepository {
	permissionRepoIns.db = db
	return permissionRepoIns
}
func (p *permissionRepository) Insert(permission interface{}) error {
	err := p.db.Create(permission).Error
	return err

}

func (p *permissionRepository) Update(permission interface{}) error {
	err := p.db.Save(permission).Error
	return err

}

func (p *permissionRepository) Delete(permission interface{}) error {
	return nil
}

func (p *permissionRepository) FindSingle(condition string, params ...interface{}) interface{} {
	return nil
}

func (p *permissionRepository) FindOne(id int) interface{} {
	var permission Permission
	p.db.Where("id = ?", id).First(&permission)
	if permission.ID == 0 {
		return nil
	}
	return &permission

}

func (p *permissionRepository) FindMore(condition string, params ...interface{}) interface{} {
	return nil
}

func (p *permissionRepository) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *helper.PageBean) {
	total := 0
	var rows []Permission

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			p.db = p.db.Where(k, v)

		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			p.db = p.db.Where(k, v)

		}
	}
	p.db.Limit(pageSize).Offset((page - 1) * pageSize).Order("id").Find(&rows).Count(&total)
	return &helper.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
