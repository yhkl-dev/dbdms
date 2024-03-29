package resources

import (
	"dbdms/db"
)

// Resource resource struct
type Resource struct {
	ResourceID               int    `gorm:"column:resource_id;primary_key" json:"resource_id" form:"resource_id"`
	ResourceName             string `gorm:"type:varchar(32);column:resource_name;unique;not null" json:"resource_name" form:"resource_name" binding:"required"`
	ResourceHostIP           string `gorm:"column:resource_host_ip" json:"resource_host_ip" form:"resource_host_ip" binding:"required"`
	ResourceUser             string `gorm:"type:varchar(32);column:resource_user" json:"resource_user" form:"resource_user" binding:"required"`
	ResourcePassword         string `gorm:"type:varchar(64);column:resource_password" json:"resource_password" form:"resource_password"`
	ResourcePublicKeyName    string `gorm:"type:varchar(500);column:resource_public_key_name" json:"resource_public_key_name" form:"resource_public_key_name"`
	ResourcePublicKeyContent string `gorm:"type:varchar(500);column:resource_public_key_content" json:"resource_public_key_content" form:"resource_public_key_content"`
	ResourcePort             int    `gorm:"column:resource_port" json:"resource_port" form:"resource_port" binding:"required"`
	ResourcePassSalt         string `gorm:"type:varchar(32);column:resource_pass_salt" json:"resource_pass_salt" form:"resource_pass_salt"`
	ResourceDatabaseName     string `gorm:"type:varchar(32);column:resource_database_name" json:"resource_database_name" form:"resource_database_name"`
	ResourceDescription      string `gorm:"type:varchar(32);column:resource_description" json:"resource_description" form:"resource_description" binding:"required"`
	ResourceTypeID           int
	ResourceType             ResourceType `gorm:"foreignKey:ResourceTypeID" json:"resource_type" form:"resource_port" binding:"required"`
}

// ResourceType resource type struct
type ResourceType struct {
	ResourceTypeID          int    `gorm:"column:resource_type_id;primary_key" json:"resource_type_id" form:"resource_type_id"`
	ResourceParentTypeID    int    `gorm:"column:resource_parent_id;default:0" json:"resource_parent_id" form:"resource_parent_id"`
	ResourceTypeName        string `gorm:"type:varchar(32);column:resource_type_name;unique;not null" json:"resource_type_name" form:"resource_type_name" `
	ResourceTypeDescription string `gorm:"type:varchar(32);column:resource_type_description;unique;not null" json:"resource_type_description" form:"resource_type_description"`
}

// TableName define table name
func (r *Resource) TableName() string {
	return "resources"
}

// TableName define table name
func (r *ResourceType) TableName() string {
	return "resources_type"
}

func init() {
	_ = db.SQL.AutoMigrate(&Resource{})
	_ = db.SQL.AutoMigrate(&ResourceType{})
}
