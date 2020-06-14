package permission

import (
	"dbdms/apps/repository"

	"github.com/jinzhu/gorm"
)

type PermissionRepository interface {
	repository.Repository
}

type permissionRepository struct {
	db *gorm.DB
}

var permissionRepoIns = &permissionRepository{}

func PermissionRepositoryInstance(db *gorm.DB) PermissionRepository {
	permissionRepoIns.db = db
	return permissionRepoIns
}
