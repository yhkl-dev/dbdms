package resources

import (
	"dbdms/utils"
	"errors"
)

// Service resource service instance
type ResourceService interface {
	GetResources() []*Resource
	GetResourceByID(id int) *Resource
	DeleteResourceByID(id int) error
	GetResourcePage(page int, pageSize int, resource *Resource) *utils.PageBean
	SaveOrUpdateResource(resource *Resource) error
}

// ResourceTypeService resource type service instance
type ResourceTypeService interface {
	GetResourceTypes() []*ResourceType
	GetResourceTypeByID(id int) *ResourceType
	GetResourceTypePage(page int, pageSize int, resourceType *ResourceType) *utils.PageBean
	DeleteResourceTypeByID(id int) error
	SaveOrUpdateResourceType(resource *ResourceType) error
}

type resourceService struct {
	repo Repo
}

type resourceTypeService struct {
	repo Repo
}

var resourceServiceIns = &resourceService{}
var resourceTypeServiceIns = &resourceTypeService{}

// ResourceServiceInstance 获取 ResourceServiceInstance 实例
func ResourceServiceInstance(repo Repo) ResourceService {
	resourceServiceIns.repo = repo
	return resourceServiceIns
}

// ResourceTypeServiceInstance 获取 ResourceTypeServiceInstance 实例
func ResourceTypeServiceInstance(repo Repo) ResourceTypeService {
	resourceTypeServiceIns.repo = repo
	return resourceTypeServiceIns
}

func (us *resourceService) GetResources() []*Resource {
	resources := us.repo.FindMore("1=1").([]*Resource)
	return resources
}

func (us *resourceService) GetResourceByID(id int) *Resource {
	if id <= 0 {
		return nil
	}
	resource := us.repo.FindOne(id)
	if resource != nil {
		return resource.(*Resource)
	}
	return nil
}

func (us *resourceService) SaveOrUpdateResource(resource *Resource) error {
	if resource == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	if resource.ResourceID == 0 {
		return us.repo.Insert(resource)
	}
	persist := us.GetResourceByID(resource.ResourceID)
	if persist == nil || persist.ResourceID == 0 {
		return errors.New(utils.StatusText(utils.UpdateObjIsNil))
	}
	return us.repo.Update(resource)
}

func (us *resourceService) DeleteResourceByID(id int) error {
	resource := us.repo.FindOne(id).(*Resource)
	if resource == nil || resource.ResourceID == 0 {
		return errors.New(utils.StatusText(utils.DeleteObjIsNil))
	}
	return us.repo.Update(resource)
}

func (us *resourceService) GetResourcePage(page int, pageSize int, resource *Resource) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	addCons := make(map[string]interface{})
	if resource != nil && resource.ResourceName != "" {
		addCons["resource_name LIKE ?"] = "%" + resource.ResourceName + "%"
	}
	if resource != nil && resource.ResourceHostIP != "" {
		addCons["resource_host_ip LIKE ?"] = resource.ResourceHostIP + "%"
	}
	pageBean := us.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (us *resourceTypeService) GetResourceTypes() []*ResourceType {
	resoruceTypes := us.repo.FindMore("1=1").([]*ResourceType)
	return resoruceTypes
}

func (us *resourceTypeService) GetResourceTypeByID(id int) *ResourceType {
	if id <= 0 {
		return nil
	}
	resourceType := us.repo.FindOne(id)
	if resourceType != nil {
		return resourceType.(*ResourceType)
	}
	return nil
}

func (us *resourceTypeService) SaveOrUpdate(resourceType *ResourceType) error {
	if resourceType == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	if resourceType.ResourceTypeID == 0 {
		return us.repo.Insert(resourceType)
	}
	persist := us.GetResourceTypeByID(resourceType.ResourceTypeID)
	if persist == nil || persist.ResourceTypeID == 0 {
		return errors.New(utils.StatusText(utils.UpdateObjIsNil))
	}
	return us.repo.Update(resourceType)
}

func (us *resourceTypeService) DeleteResourceTypeByID(id int) error {
	resourceType := us.repo.FindOne(id).(*ResourceType)
	if resourceType == nil || resourceType.ResourceTypeID == 0 {
		return errors.New(utils.StatusText(utils.DeleteObjIsNil))
	}
	return us.repo.Update(resourceType)
}

func (us *resourceTypeService) GetResourceTypePage(page int, pageSize int, resourceType *ResourceType) *utils.PageBean {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	addCons := make(map[string]interface{})
	if resourceType != nil && resourceType.ResourceTypeName != "" {
		addCons["resource_type_name LIKE ?"] = "%" + resourceType.ResourceTypeName + "%"
	}
	pageBean := us.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (us *resourceTypeService) SaveOrUpdateResourceType(resourceType *ResourceType) error {
	if resourceType == nil {
		return errors.New(utils.StatusText(utils.SaveObjIsNil))
	}
	if resourceType.ResourceTypeID == 0 {
		return us.repo.Insert(resourceType)
	}
	persist := us.GetResourceTypeByID(resourceType.ResourceTypeID)
	if persist == nil || persist.ResourceTypeID == 0 {
		return errors.New(utils.StatusText(utils.UpdateObjIsNil))
	}
	return us.repo.Update(resourceType)
}