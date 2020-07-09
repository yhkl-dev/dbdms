package dbdm

import "dbdms/apps/resource"

type Document struct {
	ID       int    `gorm:"AUTO_INCREMENT;primary_key"`
	Version  string `grom:"type:varchar(32);not null" json:"version"`
	Context  string `gorm:"type:text; not null"`
	BelongDB resource.DBResource
}
