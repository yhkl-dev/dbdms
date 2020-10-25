package resources

import "dbdms/db"

// Resources role struct
type Resource struct {
	ResourceID          int    `gorm:"column:resource_id;primary_key" json:"resource_id" form:"resource_id"`
	ResourceName        string `gorm:"type:varchar(32);column:resource_name;unique;not null" json:"resource_name" form:"resource_name" binding:"required"`
	ResourceHostIP      string `gorm:"column:resource_host_ip" json:"resource_host_ip" form:"resource_host_ip" binding:"required"`
	ResourceUser        string `gorm:"type:varchar(32);column:resource_user" json:"resource_user" form:"resource_user" binding:"required"`
	ResourcePassword    string `gorm:"type:varchar(64);column:resource_password" json:"resource_password" form:"resource_password" binding:"required"`
	ResourcePort        int    `gorm:"column:resource_port" json:"resource_port" form:"resource_port" binding:"required"`
	ResourceDescription string `gorm:"type:varchar(32);column:resource_description" json:"resource_description" form:"resource_description" binding:"required"`
	ResourceTypeID      int
	ResourceType        ResourceType `gorm:"foreignKey:ResourceTypeID" json:"resource_type" form:"resource_port" binding:"required"`
}

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
