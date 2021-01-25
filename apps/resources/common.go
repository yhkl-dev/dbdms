package resources

type resourceQueryParams struct {
	ResourceName           string `json:"resource_name" form:"resource_name"`
	ResourceHostIP         string `json:"resource_host_ip" form:"resource_host_ip"`
	ResourceTypeName       string `json:"resource_type_name" form:"resource_type_name"`
	ResourceParentTypeName string `json:"resource_parent_type_name" form:"resource_parent_type_name"`
	Page                   int    `json:"page" form:"page"`
	PageSize               int    `json:"page_size" form:"page_size"`
}

type resourceTypeQueryParams struct {
	ResourceTypeName string `json:"resource_name" form:"resource_name"`
	Page             int    `json:"page" form:"page"`
	PageSize         int    `json:"page_size" form:"page_size"`
}

type testConnection struct {
	IP           string `json:"ip" form:"ip"`
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"passswor"`
	Port         int    `json:"port" form:"port"`
	Schema       string `json:"schema" form:"schema"`
	ResourceType int    `json:"resource_type" form:"resource_type"`
}
