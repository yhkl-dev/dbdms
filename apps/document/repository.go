package document

import (
	"dbdms/apps"
	"dbdms/utils"
	"gorm.io/gorm"
)

// Repo document interface implement from common interface
type Repo interface {
	apps.RepositoryInterface
}

// VersionRepo document interface implement from common interface
type VersionRepo interface {
	Insert(m interface{}) error
	FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean)
}

type documentRepo struct {
	db *gorm.DB
}

type versionRepo struct {
	db *gorm.DB
}

var documentRepoInstance = &documentRepo{}
var versionRepoInstance = &versionRepo{}

// Interface instance for storage object
func Interface(db *gorm.DB) Repo {
	documentRepoInstance.db = db
	return documentRepoInstance
}

// VersionInterface instance for storage object
func VersionRepoInterface(db *gorm.DB) VersionRepo {
	versionRepoInstance.db = db
	return versionRepoInstance
}

func (vr *versionRepo) Insert(m interface{}) error {
	err := vr.db.Create(m).Error
	return err
}

func (vr *versionRepo) Update(m interface{}) error {
	err := vr.db.Save(m).Error
	return err
}

func (vr *versionRepo) Delete(m interface{}) error {
	return vr.db.Delete(m).Error
}

func (dr *documentRepo) Insert(m interface{}) error {
	err := dr.db.Create(m).Error
	return err
}

func (dr *documentRepo) Update(m interface{}) error {
	err := dr.db.Save(m).Error
	return err
}

func (dr *documentRepo) Delete(m interface{}) error {
	return dr.db.Delete(m).Error
}

func (dr *documentRepo) FindMore(condition string, params ...interface{}) interface{} {
	roles := make([]*DatabaseDocument, 0)
	dr.db.Where(condition, params).Find(&roles)
	return roles
}

func (dr *documentRepo) FindOne(id int) interface{} {
	var document DatabaseDocument
	err := dr.db.Where("role_id = ?", id).First(&document).Error
	if document.DocumentID == 0 || err != nil {
		return nil
	}
	return &document
}

func (dr *documentRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	var rows []DatabaseDocument

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			dr.db = dr.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			dr.db = dr.db.Where(k, v)
		}
	}
	dr.db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}

func (dr *documentRepo) FindSingle(condition string, params ...interface{}) interface{} {
	var document DatabaseDocument
	dr.db.Where(condition, params).First(&document)
	if document.DocumentUUID != "" {
		return &document
	}
	return nil
}

func (dr *versionRepo) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean) {
	var total int64
	var rows []DocumentVersion

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			dr.db = dr.db.Where(k, v)
		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			dr.db = dr.db.Where(k, v)
		}
	}
	dr.db.Limit(pageSize).Offset((page - 1) * pageSize).Preload("Resource").Find(&rows).Count(&total)
	return &utils.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
