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

type documentRepo struct {
	db *gorm.DB
}

var documentRepoInstance = &documentRepo{}

// Interface instance for storage object
func Interface(db *gorm.DB) Repo {
	documentRepoInstance.db = db
	return documentRepoInstance
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
