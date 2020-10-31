package document

type documentQueryParams struct {
	DocumentDBName    string `json:"document_db_name" form:"resource_name"`
	DocumentTableName string `json:"document_table_name" form:"resource_host_ip"`
	ResourceName      string `json:"resource_name" form:"resource_name"`
	Page              int    `json:"page" form:"page"`
	PageSize          int    `json:"page_size" form:"page_size"`
}
