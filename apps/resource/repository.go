package resource

import (
	repository "dbdms/apps/repository"
	helper "dbdms/helpers"

	"github.com/jinzhu/gorm"
)

type ResourceRepository interface {
	repository.Repository
}

type resourceRepository struct {
	db *gorm.DB
}

var resourceRepoIns = &resourceRepository{}

// ResourceRepositoryInterface instance for storage object
func ResourceRepositoryInterface(db *gorm.DB) ResourceRepository {
	resourceRepoIns.db = db
	return resourceRepoIns
}

func (r *resourceRepository) Insert(resource interface{}) error {
	err := r.db.Create(resource).Error
	return err
}

func (r *resourceRepository) Update(resource interface{}) error {
	err := r.db.Save(resource).Error
	return err
}

func (r *resourceRepository) Delete(resource interface{}) error {
	err := r.db.Delete(resource).Error
	return err
}

// find user by name
func (r *resourceRepository) FindSingle(condition string, params ...interface{}) interface{} {
	return nil
}

// find user by id
func (r *resourceRepository) FindOne(id int) interface{} {
	return nil
}

// 条件查询返回多值
func (r *resourceRepository) FindMore(condition string, params ...interface{}) interface{} {
	return nil
}

func (r *resourceRepository) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *helper.PageBean) {
	total := 0
	var rows []DBResource

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
	return &helper.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
