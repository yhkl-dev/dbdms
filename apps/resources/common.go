package resources

type resourceQueryParams struct {
	ResourceName   string `json:"resource_name" form:"resource_name"`
	ResourceHostIP string `json:"resource_host_ip" form:"resource_host_ip"`
	Page           int    `json:"page" form:"page"`
	PageSize       int    `json:"page_size" form:"page_size"`
}

type resourceTypeQueryParams struct {
	ResourceTypeName string `json:"resource_name" form:"resource_name"`
	Page             int    `json:"page" form:"page"`
	PageSize         int    `json:"page_size" form:"page_size"`
}
