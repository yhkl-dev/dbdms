package document

import (
	"dbdms/utils"
)

// Service document service instance
type Service interface {
	//GetDocuments() []*DatabaseDocument
	//GetResourceByID(id int) *Resource
	//DeleteResourceByID(id int) error
	GetDocumentPage(page int, pageSize int, document *DatabaseDocument) *utils.PageBean
	//SaveOrUpdateResource(resource *Resource) error
}

type documentService struct {
	repo Repo
}

var serviceIns = &documentService{}

// DocumentServiceInstance 获取 serviceIns 实例
func ServiceInstance(repo Repo) Service {
	serviceIns.repo = repo
	return serviceIns
}

func (ds *documentService) GetDocumentPage(page int, pageSize int, document *DatabaseDocument) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	addCons := make(map[string]interface{})
	if document != nil && document.ResourceID != 0 {
		addCons["resource_id = ?"] = document.ResourceID
	}
	if document != nil && document.DocumentDBName != "" {
		addCons["document_db_name LIKE ?"] = "%" + document.DocumentDBName + "%"
	}
	if document != nil && document.DocumentTableName != "" {
		addCons["document_table_name LIKE ?"] = document.DocumentTableName + "%"
	}

	pageBean := ds.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}
