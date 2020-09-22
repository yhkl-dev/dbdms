package project

import (
	repository "dbdms/apps/repository"
	helper "dbdms/helpers"

	"github.com/jinzhu/gorm"
)

type ProjectRepository interface {
	repository.Repository
}

type projectRepository struct {
	db *gorm.DB
}

var projectRepoIns = &projectRepository{}

// ProjectRepositoryInterface instance for storage object
func ProjectRepositoryInterface(db *gorm.DB) ProjectRepository {
	projectRepoIns.db = db
	return projectRepoIns
}

func (r *projectRepository) Insert(project interface{}) error {
	err := r.db.Create(project).Error
	return err
}

func (r *projectRepository) Update(project interface{}) error {
	//r.db.Model(project.(*Permission)).Association("Roles").Replace(project.(*Permission).Roles)
	err := r.db.Save(project).Error
	return err
}

func (r *projectRepository) Delete(project interface{}) error {
	err := r.db.Delete(project).Error
	return err
}

// find project by name
func (r *projectRepository) FindSingle(condition string, params ...interface{}) interface{} {
	var project Project
	//r.db.Preload("Roles").Where(condition, params).First(&project)
	r.db.Where(condition, params).First(&project)
	if project.ID == 0 {
		return nil
	}
	return &project
}

// find project by id
func (r *projectRepository) FindOne(id int) interface{} {
	var project Project
	r.db.Where("id = ?", id).First(&project)
	if project.ID == 0 {
		return nil
	}
	//	r.db.Model(&project).Association("Roles").Find(&roles)
	//	project.Roles = roles
	return &project
}

// 条件查询返回多值
func (r *projectRepository) FindMore(condition string, params ...interface{}) interface{} {
	projects := make([]*Project, 0)
	//r.db.Preload("Roles").Where(condition, params).Find(&projects)
	r.db.Where(condition, params).Find(&projects)
	return projects

}

func (r *projectRepository) FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *helper.PageBean) {
	total := 0
	//	rows := make([]Permission, 0)
	var rows []Project

	if andCons != nil && len(andCons) > 0 {
		for k, v := range andCons {
			r.db = r.db.Where(k, v)

		}
	}
	if orCons != nil && len(orCons) > 0 {
		for k, v := range orCons {
			r.db = r.db.Where(k, v)
		}
	}
	//r.db.Preload("Roles").Limit(pageSize).Offset((page - 1) * pageSize).Order("login_time desc").Find(&rows).Count(&total)
	r.db.Limit(pageSize).Offset((page - 1) * pageSize).Order("create_at desc").Find(&rows).Count(&total)
	return &helper.PageBean{Page: page, PageSize: pageSize, Total: total, Rows: rows}
}
