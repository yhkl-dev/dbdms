package resources

import (
	"dbdms/db"
	"dbdms/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListAllResources(context *gin.Context) {
	query := resourceQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	pageBean := resourceService.GetResourcePage(query.Page, query.PageSize, &Resource{ResourceName: query.ResourceName, ResourceHostIP: query.ResourceHostIP})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

func CreateResource(context *gin.Context) {
	resource := &Resource{}
	err := context.Bind(resource)
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceType := resourceTypeService.GetResourceTypeByID(resource.ResourceType.ResourceTypeID)
	if resourceType == nil {
		context.JSON(http.StatusBadRequest, &utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.ResourceTypeDoesNotExist),
		})
		return
	}
	resource.ResourceType = ResourceType{ResourceTypeID: resourceType.ResourceTypeID, ResourceTypeName: resourceType.ResourceTypeName, ResourceTypeDescription: resourceType.ResourceTypeDescription}
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}
	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	err = resourceService.SaveOrUpdateResource(resource)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusBadRequest, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}

func UpdateResource(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	resource := &Resource{}
	err := context.Bind(resource)

	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	resourceOrigin := resourceService.GetResourceByID(id)
	if resourceOrigin == nil {
		context.JSON(http.StatusBadRequest, &utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.ResourceDoesNotExist),
		})
		return
	}
	resource.ResourceID = resourceOrigin.ResourceID
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceType := resourceTypeService.GetResourceTypeByID(resource.ResourceType.ResourceTypeID)
	if resourceType == nil {
		context.JSON(http.StatusBadRequest, &utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.ResourceTypeDoesNotExist),
		})
		return
	}
	resource.ResourceType = ResourceType{ResourceTypeID: resourceType.ResourceTypeID, ResourceTypeName: resourceType.ResourceTypeName, ResourceTypeDescription: resourceType.ResourceTypeDescription}
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}
	err = resourceService.SaveOrUpdateResource(resource)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusBadRequest, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}

func DeleteResourceByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	err := resourceService.DeleteResourceByID(id)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.DeleteStatusErr),
			Content: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Message: utils.StatusText(utils.DeleteStatusOK),
	})
}

func ListAllResourceTypes(context *gin.Context) {
	query := resourceTypeQueryParams{}
	err := context.BindQuery(&query)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.ParamParseError),
			Content: err.Error(),
		})
	}
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	pageBean := resourceTypeService.GetResourceTypePage(query.Page, query.PageSize, &ResourceType{ResourceTypeName: query.ResourceTypeName})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

// CreateResourceType
func CreateResourceType(context *gin.Context) {
	resourceType := &ResourceType{}
	err := context.Bind(resourceType)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	err = resourceTypeService.SaveOrUpdateResourceType(resourceType)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusBadRequest, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}

// UpdateResourceType
func UpdateResourceType(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	resourceType := &ResourceType{}

	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceTypeQuery := resourceTypeService.GetResourceTypeByID(id)
	if resourceTypeQuery != nil {
		resourceType.ResourceTypeID = resourceTypeQuery.ResourceTypeID
	}
	err := resourceTypeService.SaveOrUpdateResourceType(resourceType)
	if err == nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.SaveStatusOK),
		})
		return
	}
	context.JSON(http.StatusBadRequest, &utils.JSONObject{
		Code:    "0",
		Message: err.Error(),
	})
}

func DeleteResourceTypeByID(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceType := resourceTypeService.GetResourceTypeByID(id)
	if resourceType == nil || resourceType.ResourceTypeID == 0 {
		context.JSON(http.StatusBadRequest, &utils.JSONObject{
			Code:    "1",
			Message: utils.StatusText(utils.ResourceDoesNotExist),
		})
		return
	}
	err := resourceTypeService.DeleteResourceTypeByID(id)
	if err != nil {
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.DeleteStatusErr),
			Content: err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Message: utils.StatusText(utils.DeleteStatusOK),
	})
}