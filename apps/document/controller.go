package document

import (
	"dbdms/apps/resources"
	"dbdms/db"
	"dbdms/tbls/command"
	"dbdms/utils"
	"fmt"
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
	if query.ResourceName != "" {
		resource := resourceIns.GetResourceByName(query.ResourceName)
		if resource == nil {
			context.JSON(http.StatusOK, utils.JSONObject{
				Code:    "0",
				Message: utils.StatusText(utils.ParamParseError),
				Content: "",
			})
			return
		}
		pageBean := documentService.GetDocumentPage(query.Page,
			query.PageSize,
			&DatabaseDocument{DocumentDBName: query.DocumentDBName, DocumentTableName: query.DocumentTableName, ResourceID: resource.ResourceID})
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Content: pageBean,
		})
		return
	}
	pageBean := documentService.GetDocumentPage(query.Page,
		query.PageSize,
		&DatabaseDocument{DocumentDBName: query.DocumentDBName, DocumentTableName: query.DocumentTableName})
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
	fmt.Println("DSN",  DSN)
	command.Doc(DSN)
	fmt.Println("DSNE")


	return
}
