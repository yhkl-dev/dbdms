package document

import (
	"dbdms/apps/resources"
	"dbdms/db"
	"dbdms/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ListAllResourcesDocuments
func ListAllResourcesDocuments(context *gin.Context) {
	query := documentQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	resourceIns := resources.ResourceServiceInstance(resources.RepoInterface(db.SQL))
	documentService := ServiceInstance(Interface(db.SQL))
	if query.ResourceID != 0 {
		resource := resourceIns.GetResourceByID(query.ResourceID)
		if resource == nil {
			context.JSON(http.StatusOK, utils.JSONObject{
				Code:    "0",
				Message: utils.StatusText(utils.ParamParseError),
				Content: "xxxxx",
			})
			return
		}
		pageBean := documentService.GetDocumentPage(
			query.Page,
			query.PageSize,
			&DatabaseDocument{
				DocumentDBName: query.DocumentDBName,
				DocumentTableName: query.DocumentTableName,
				DocumentVersion:   query.DocumentVersion,
				ResourceID:        resource.ResourceID,
			})
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Content: pageBean,
		})
		return
	}
	pageBean := documentService.GetDocumentPage(query.Page,
		query.PageSize,
		&DatabaseDocument{
			DocumentDBName: query.DocumentDBName,
			DocumentTableName: query.DocumentTableName,
			DocumentVersion:   query.DocumentVersion,
			ResourceID:        query.ResourceID,
		})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

// GenerateDocument
func GenerateDocument(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	resourceService := resources.ResourceServiceInstance(resources.RepoInterface(db.SQL))
	DSN := resourceService.GenerateDSN(id)
	//document := &DatabaseDocument{}
	documentService := ServiceInstance(Interface(db.SQL))
	versionService := VersionServiceInstance(VersionRepoInterface(db.SQL))
	go Doc(DSN, id, documentService, versionService)
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Message: utils.StatusText(utils.SaveStatusOK),
		Content: "",
	})
	return
}

func ListDocumentVersions(context *gin.Context) {
	query := versionQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
		return
	}

	versionService := VersionServiceInstance(VersionRepoInterface(db.SQL))
	pageBean := versionService.GetVersionPage(
		query.Page,
		query.PageSize, &DocumentVersion{
			ResourceID: query.ResourceID,
			VersionName: query.Version,
		})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}
