package helper

// LoginParams login paramters
type LoginParams struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// JSONObject the interface data struct
type JSONObject struct {
	Code    string
	Content interface{}
	Message string
}

// PageBean global paginator
type PageBean struct {
	Page     int
	PageSize int
	Total    int
	Rows     interface{}
}
