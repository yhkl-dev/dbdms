package resources

import (
	"dbdms/utils"
	"encoding/base64"
	"errors"
	"fmt"
)

// Service resource service instance
type ResourceService interface {
	GetResources() []*Resource
	GetResourceByID(id int) *Resource
	GetResourceByName(name string) *Resource
	DeleteResourceByID(id int) error
	GetResourcePage(page int, pageSize int, resource *Resource) *utils.PageBean
	SaveOrUpdateResource(resource *Resource) error
	GenerateDSN(id int) string
}

// ResourceTypeService resource type service instance
type ResourceTypeService interface {
	GetResourceTypes() []*ResourceType
	GetResourceTypeByID(id int) *ResourceType
	GetResourceTypePage(page int, pageSize int, resourceType *ResourceType) *utils.PageBean
	DeleteResourceTypeByID(id int) error
	SaveOrUpdateResourceType(resource *ResourceType) error
	GetResourceTypeByName(resourceTypeName string) *ResourceType
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

func (us *resourceService) GetResourceByName(name string) *Resource {
	if name == "" {
		return nil
	}
	resource := us.repo.FindByName(name)
	if resource != nil {
		return resource.(*Resource)
	}
	return nil
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
		salt := utils.GetRandomString(16)
		resource.ResourcePassSalt = string(salt[:])
		encryptBytes, err :=  utils.AesEncrypt([]byte(resource.ResourcePassword), salt)

		if err != nil {
			return err
		}
		resource.ResourcePassword = base64.StdEncoding.EncodeToString(encryptBytes)

		bytesPass, err := base64.StdEncoding.DecodeString(resource.ResourcePassword)
		if err != nil {
			fmt.Println(err)
		}

		tpass, err := utils.AesDecrypt(bytesPass, []byte(resource.ResourcePassSalt))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("解密后:%s\n", tpass)

		return us.repo.Insert(resource)
	}
	persist := us.GetResourceByID(resource.ResourceID)
	if persist == nil || persist.ResourceID == 0 {
		return errors.New(utils.StatusText(utils.UpdateObjIsNil))
	}
	if persist.ResourcePassword != resource.ResourcePassword {
		encryptBytes, err :=  utils.AesEncrypt([]byte(resource.ResourcePassword), []byte(resource.ResourcePassSalt))
		if err != nil {
			return err
		}
		resource.ResourcePassword = base64.StdEncoding.EncodeToString(encryptBytes)

		bytesPass, err := base64.StdEncoding.DecodeString(resource.ResourcePassword)
		if err != nil {
			fmt.Println(err)
		}

		tpass, err := utils.AesDecrypt(bytesPass, []byte(resource.ResourcePassSalt))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("更新解密后:%s\n", tpass)
	}
	return us.repo.Update(resource)
}

func (us *resourceService) DeleteResourceByID(id int) error {
	resource := us.repo.FindOne(id).(*Resource)
	if resource == nil || resource.ResourceID == 0 {
		return errors.New(utils.StatusText(utils.DeleteObjIsNil))
	}
	return us.repo.Delete(resource)
}

func(us *resourceService) GenerateDSN(id int) string {
	r := us.GetResourceByID(id)
	if r == nil {
		return ""
	}
	bytesPass, err := base64.StdEncoding.DecodeString(r.ResourcePassword)
	if err != nil {
		fmt.Println(err)
	}

	tpass, err := utils.AesDecrypt(bytesPass, []byte(r.ResourcePassSalt))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("解密后:%s\n", tpass)
	fmt.Println("ResourceTypeName", r.ResourceType.ResourceTypeName)
	if r.ResourceType.ResourceTypeName == "postgres" {
		// postgres://dbuser:dbpass@hostname:5432/dbname
		return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", r.ResourceUser, tpass, r.ResourceHostIP, r.ResourcePort, r.ResourceDatabaseName)
	}
	return ""
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
	if resource != nil && resource.ResourceTypeID != 0 {
		addCons["resource_type_id = ?"] = resource.ResourceTypeID
	}
	pageBean := us.repo.FindPage(page, pageSize, addCons, nil)
	return pageBean
}

func (us *resourceTypeService) GetResourceTypes() []*ResourceType {
	resourceTypes := us.repo.FindMore("1 = 1").([]*ResourceType)
	return resourceTypes
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

func (us *resourceTypeService) GetResourceTypeByName(resourceTypeName string) *ResourceType {
	if resourceTypeName == "" {
		return nil
	}
	resourceType := us.repo.FindByName(resourceTypeName)
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