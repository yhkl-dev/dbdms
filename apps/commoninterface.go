package apps

import "dbdms/utils"

// RepositoryInterface common interface
type RepositoryInterface interface {
	Insert(m interface{}) error
	Update(m interface{}) error
	Delete(m interface{}) error
	FindOne(id int) interface{}
	FindSingle(condition string, params ...interface{}) interface{}
	FindMore(condition string, params ...interface{}) interface{}
	FindPage(page int, pageSize int, andCons map[string]interface{}, orCons map[string]interface{}) (pageBean *utils.PageBean)
}
