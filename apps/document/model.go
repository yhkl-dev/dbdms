package document

import "dbdms/apps/resources"

// DatabaseDocument database  struct
type DatabaseDocument struct {
	DocumentID        int    `gorm:"column:document_id;primary_key" json:"document_id" form:"document_id"`
	DocumentUUID      string `gorm:"type:varchar(32);column:document_uuid" json:"document_uuid" form:"document_uuid"`
	DocumentDBName    string `gorm:"type:varchar(32);column:document_db_name" json:"document_db_name" form:"document_db_name"`
	DocumentTableName string `gorm:"type:varchar(32);column:document_table_name" json:"document_table_name" form:"document_table_name"`
	Created           int64  `gorm:"autoCreateTime"`
	ResourceID        int
	Resource          resources.Resource `gorm:"foreignKey:ResourceID" json:"resource_id" form:"resource_id" binding:"required"`
}

// TableName define table name
func (r *DatabaseDocument) TableName() string {
	return "database_document"
}