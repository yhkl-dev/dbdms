package resources

import (
	"dbdms/apps"
	"dbdms/utils"
	"gorm.io/gorm"
)

// Repo resource interface implement  from common interface
type Repo interface {
	apps.RepositoryInterface
	FindByName(name string) interface{}
}

type resourceRepo struct {
	db *gorm.DB
}

var resourceRepoInstance = &resourceRepo{}

// RepoInterface instance for storage object
func RepoInterface(db *gorm.DB) Repo {
	resourceRepoInstance.db = db
	return resourceRepoInstance
}

func (rp *resourceRepo) Insert(m interface{}) error {
	err := rp.db.Create(m).Error
	return err
}

func (rp *resourceRepo) Update(m interface{}) error {
	err := rp.db.Save(m).Error
	return err
}

func (rp *resourceRepo) Delete(m interface{}) error {
	err := rp.db.Delete(m).Error
	return err
}

func (rp *resourceRepo) FindMore(condition string, params ...interface{}) interface{} {
	resources := make([]*Resource, 0)
	rp.db.Where(condition, params).Find(&resources)
	return resources
}

func (rp *resourceRepo) FindOne(id int) interface{} {
	var resource Resource
	err := rp.db.Where("resource_id = ?", id).First(&resource).Error
	if resource.ResourceID == 0 || err != nil {
		return nil
	}
	return &resource
}

func (rp *resourceRepo) FindByName(resourceTypeName string) interface{} {
	var resourceType ResourceType
	err := rp.db.Where("resource_name = ?", resourceTypeName).First(&resourceType).Error
	if resourceType.ResourceTypeID == 0 || err != nil {
		return nil
	}
	return &resourceType
}

func (rp *resourceRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	var rows []Resource

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
	rp.db.Preload("ResourceType").Limit(pageSize).Offset((page - 1) * pageSize).Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}

func (rp *resourceRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var resource Resource
	rp.db.Where(condition, params).First(&resource)
	if resource.ResourceID != 0 {
		return &resource
	}
	return nil
}

// Repo resource interface implement  from common interface
type ReTypeRepo interface {
	apps.RepositoryInterface
	FindByName(resourceTypeName string) interface{}
}

type resourceTypeRepo struct {
	db *gorm.DB
}

var resourceTypeRepoInstance = &resourceTypeRepo{}

// RepoInterface instance for storage object
func TypeRepoInterface(db *gorm.DB) Repo {
	resourceTypeRepoInstance.db = db
	return resourceTypeRepoInstance
}

func (rp *resourceTypeRepo) Insert(m interface{}) error {
	err := rp.db.Create(m).Error
	return err
}

func (rp *resourceTypeRepo) Update(m interface{}) error {
	err := rp.db.Save(m).Error
	return err
}

func (rp *resourceTypeRepo) Delete(m interface{}) error {
	err := rp.db.Delete(m).Error
	return err
}

func (rp *resourceTypeRepo) FindMore(condition string, params ...interface{}) interface{} {
	resourceType := make([]*ResourceType, 0)
	rp.db.Where(condition, params).Preload("resource_type").Find(&resourceType)
	return resourceType
}

func (rp *resourceTypeRepo) FindOne(id int) interface{} {
	var resourceType ResourceType
	err := rp.db.Where("resource_type_id = ?", id).First(&resourceType).Error
	if resourceType.ResourceTypeID == 0 || err != nil {
		return nil
	}
	return &resourceType
}

func (rp *resourceTypeRepo) FindByName(resourceTypeName string) interface{} {
	var resourceType ResourceType
	err := rp.db.Where("resource_type_name = ?", resourceTypeName).First(&resourceType).Error
	if resourceType.ResourceTypeID == 0 || err != nil {
		return nil
	}
	return &resourceType
}

func (rp *resourceTypeRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	var rows []ResourceType

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

func (rp *resourceTypeRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var resourceType ResourceType
	rp.db.Where(condition, params).First(&resourceType)
	if resourceType.ResourceTypeID != 0 {
		return &resourceType
	}
	return nil
}
