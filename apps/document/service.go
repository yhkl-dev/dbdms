package document

import (
	"dbdms/utils"
	"errors"
	"fmt"
)

// Service document service instance
type Service interface {
	//GetDocuments() []*DatabaseDocument
	GetDocumentByID(id int) *DatabaseDocument
	//DeleteResourceByID(id int) error
	GetDocumentPage(page int, pageSize int, document *DatabaseDocument) *utils.PageBean
	SaveDocument(document *DatabaseDocument) error
}

// VersionService document service instance
type VersionService interface {
	GetVersionPage(page int, pageSize int, version *DocumentVersion) *utils.PageBean
	SaveVersion(version *DocumentVersion) error
}

type documentService struct {
	repo Repo
}

type documentVersionService struct {
	repo VersionRepo
}

var serviceIns = &documentService{}
var versionServiceIns = &documentVersionService{}

// DocumentServiceInstance 获取 serviceIns 实例
func ServiceInstance(repo Repo) Service {
	serviceIns.repo = repo
	return serviceIns
}

// DocumentServiceInstance 获取 serviceIns 实例
func VersionServiceInstance(repo VersionRepo) VersionService {
	versionServiceIns.repo = repo
	return versionServiceIns
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
		addCons["document_table_name LIKE ?"] = "%" + document.DocumentTableName + "%"
	}
	if document != nil && document.DocumentVersion != "" {
		fmt.Println(">>>>", document.DocumentVersion)
		addCons["document_version = ?"] = document.DocumentVersion
	}

	pageBean := ds.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (ds *documentService) SaveDocument(document *DatabaseDocument) error {
	if document == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	return ds.repo.Insert(document)
}

func (ds *documentService) GetDocumentByID(id int) *DatabaseDocument {
	if id <= 0 {
		return nil
	}
	document := ds.repo.FindOne(id)
	if document != nil {
		return document.(*DatabaseDocument)
	}
	return nil
}

// GetVersionPage
func (vs *documentVersionService) GetVersionPage(page int, pageSize int, version *DocumentVersion) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	addCons := make(map[string]interface{})
	if version != nil && version.VersionName != "" {
		addCons["version_name = ?"] = version.VersionName
	}
	if version != nil && version.ResourceID != 0 {
		addCons["resource_id = ?"] = version.ResourceID
	}

	pageBean := vs.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (vs *documentVersionService) SaveVersion(version *DocumentVersion) error {
	if version == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	return vs.repo.Insert(version)
}
