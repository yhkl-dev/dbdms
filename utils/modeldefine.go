package utils

// LoginParams login paramters
type LoginParams struct {
	UserName     string `json:"user_name" form:"user_name"`
	UserPassword string `json:"user_password" form:"user_password"`
}

// RegisterParams user register parameters
type RegisterParams struct {
	UserName     string `json:"user_name" form:"user_name"`
	UserPhone    string `json:"user_phone" form:"user_phone"`
	UserPassword string `json:"user_password" form:"user_password"`
	UserEmail    string `json:"user_email" form:"user_email"`
}

// UserQueryParams user query parameters
type UserQueryParams struct {
	UserName  string `json:"user_name" form:"user_name"`
	UserPhone string `json:"user_phone" form:"user_phone"`
	UserEmail string `json:"user_email" form:"user_email"`
	Page      int    `json:"page" form:"page"`
	PageSize  int    `json:"page_size" form:"page_size"`
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
	Total    int64
	Rows     interface{}
}
