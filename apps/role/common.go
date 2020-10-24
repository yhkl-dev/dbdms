package role

type roleQueryParams struct {
	RoleName string `json:"role_name" form:"role_name"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
}
