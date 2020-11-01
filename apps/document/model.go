package document

import (
	"dbdms/apps/resources"
	"dbdms/db"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// DatabaseDocument database  struct
type DatabaseDocument struct {
	DocumentID        int    `gorm:"column:document_id;primary_key" json:"document_id" form:"document_id"`
	DocumentUUID      string `gorm:"type:varchar(64);column:document_uuid" json:"document_uuid" form:"document_uuid"`
	DocumentVersion   string `gorm:"type:varchar(64);column:document_version" json:"document_version" form:"document_version"`
	DocumentFileName  string `gorm:"type:varchar(64);column:document_file_name" json:"document_file_name" form:"document_file_name"`
	DocumentDBName    string `gorm:"type:varchar(64);column:document_db_name" json:"document_db_name" form:"document_db_name"`
	DocumentTableName string `gorm:"type:varchar(64);column:document_table_name" json:"document_table_name" form:"document_table_name"`
	DocumentContent   string `gorm:"type:text;column:document_content" json:"document_content" form:"document_content"`
	Created           int64  `gorm:"autoCreateTime;column:create_at"`
	ResourceID        int
	Resource          resources.Resource `gorm:"foreignKey:ResourceID" json:"resource_id" form:"resource_id" binding:"required"`
}

type DocumentVersion struct {
	VersionID        int    `gorm:"column:version_id;primary_key" json:"version_id" form:"version_id"`
	VersionName      string `gorm:"type:varchar(64);column:version_name;" json:"version_name" form:"version_name"`
	ResourceID        int
	Resource          resources.Resource `gorm:"foreignKey:ResourceID" json:"resource_id" form:"resource_id" binding:"required"`
}

// TableName define table name
func (r *DatabaseDocument) TableName() string {
	return "database_document"
}

// TableName define table name
func (r *DocumentVersion) TableName() string {
	return "database_document_version"
}

func (r *DatabaseDocument) BeforeCreate(tx *gorm.DB) (err error) {
	ul := uuid.NewV4()
	r.DocumentUUID = ul.String()
	return
}

func init() {
	_ = db.SQL.AutoMigrate(&DatabaseDocument{})
	_ = db.SQL.AutoMigrate(&DocumentVersion{})
}
