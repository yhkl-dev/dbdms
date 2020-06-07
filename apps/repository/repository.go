package repository

type Repository interface {
	Insert(m interface{}) error
	Update(m interface{}) error
	Delete(m interface{}) error
	FindOne(id int) interface{}
	FindSingle(condition string, params ...interface{}) interface{}
	FindMore(condition string, params ...interface{}) interface{}
	FindPage(condition int, pageSize int, andCons map[string]interface{}, orCOns map[string]interface{}) (pageBean *helper.PageBean)
}
