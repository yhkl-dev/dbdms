package resources

import (
	"dbdms/db"
	"dbdms/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListAllResources list all resources
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
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	if query.ResourceTypeName != "" {
		resourceTypeIns := resourceTypeService.GetResourceTypeByName(query.ResourceTypeName)
		if resourceTypeIns == nil {
			context.JSON(http.StatusOK, utils.JSONObject{
				Code:    "0",
				Message: utils.StatusText(utils.ParamParseError),
				Content: "",
			})
			return
		}
		pageBean := resourceService.GetResourcePage(query.Page,
			query.PageSize,
			&Resource{ResourceName: query.ResourceName, ResourceHostIP: query.ResourceHostIP, ResourceTypeID: resourceTypeIns.ResourceTypeID})
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Content: pageBean,
		})
		return
	}
	if query.ResourceParentTypeName != "" {
		resourceTypeIns := resourceTypeService.GetResourceTypeByName(query.ResourceParentTypeName)
		if resourceTypeIns == nil {
			context.JSON(http.StatusOK, utils.JSONObject{
				Code:    "0",
				Message: utils.StatusText(utils.ParamParseError),
				Content: "",
			})
			return
		}
		pageBean := resourceService.GetResourcePage(query.Page,
			query.PageSize,
			&Resource{ResourceName: query.ResourceName, ResourceHostIP: query.ResourceHostIP, ResourceTypeID: resourceTypeIns.ResourceParentTypeID})
		context.JSON(http.StatusOK, utils.JSONObject{
			Code:    "1",
			Content: pageBean,
		})
		return
	}
	pageBean := resourceService.GetResourcePage(query.Page,
		query.PageSize,
		&Resource{ResourceName: query.ResourceName, ResourceHostIP: query.ResourceHostIP})
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "1",
		Content: pageBean,
	})
	return
}

// CreateResource create resource
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

// UpdateResource update resource
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

// DeleteResourceByID delete resource by id
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

// ListAllResourceTypes list all resource types
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

// CreateResourceType create resoruce type
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

// UpdateResourceType udpate resource type
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

// DeleteResourceTypeByID delete resource type by id
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

// TestDBConnection test database connection function
func TestDBConnection(context *gin.Context) {
	resource := &Resource{}
	err := context.Bind(resource)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, &utils.JSONObject{
			Code:    "0",
			Message: utils.StatusText(utils.BindModelError),
			Content: err.Error(),
		})
		return
	}

	resourceService := ResourceServiceInstance(RepoInterface(db.SQL))
	resourceTypeService := ResourceTypeServiceInstance(TypeRepoInterface(db.SQL))
	resourceType := resourceTypeService.GetResourceTypeByID(resource.ResourceType.ResourceTypeID)
	res, err := resourceService.TestConnection(resource, resourceType.ResourceTypeName)
	context.JSON(http.StatusOK, utils.JSONObject{
		Code:    "0",
		Message: strconv.FormatBool(res),
		Content: err.Error(),
	})
}
