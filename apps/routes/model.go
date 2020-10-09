package routes

import (
	"dbdms/db"
	"fmt"
)

// Routes route struct
type Routes struct {
	RouteID     int    `gorm:"AUTO_INCREMENT;column:route_id;primary_key" json:"route_id" form:"route_id"`
	RouteName   string `gorm:"type:varchar(32);column:route_name" json:"route_name" form:"route_name"`
	RoutePath   string `gorm:"type:varchar(32);column:route_path" json:"route_path" form:"route_path"`
	RouteMethod string `gorm:"type:varchar(32);column:route_method" json:"route_method" form:"route_method"`
	RoutePID    int    `gorm:"column:route_pid" json:"route_pid" form:"route_pid"`
	RouteIcon   string `gorm:"type:varchar(32);column:route_icon" json:"route_icon" form:"route_icon"`
	IsShow      int    `gorm:"column:is_show" json:"is_show" form:"is_show"`
}

// RoleRouteMapping  role & route mapping struct
type RoleRouteMapping struct {
	MappingID int `gorm:"column:mapping_id; primary_key"`
	RoleID    int `gorm:"column:role_id"`
	RouteID   int `gorm:"column:route_id"`
}

// TableName define table name in database
func (r *Routes) TableName() string {
	return "sys_routes"
}

// TableName define table name in database
func (r *RoleRouteMapping) TableName() string {
	return "role_route_mapping"
}

func (r *Routes) String() string {
	return fmt.Sprintf("<RouteID: %d, RouteName: %s, RoutePath: %s, RouteMethod: %s>", r.RouteID, r.RouteName, r.RoutePath, r.RouteMethod)
}

func (r *Routes) validator() error {
	// if ok, err := regex.MatchLetterNumMinAndMax(user.UserName, 4, 6, "user_name"); !ok {
	// 	return err
	// }
	// //	if ok, err := regex.MatchMediumPassword(user.Password, 6, 13); !ok && user.ID == 0 {
	// //		return err
	// //	}
	// if ok, err := regex.IsPhone(user.UserPhone); !ok {
	// 	return err
	// }
	// if ok, err := regex.IsEmail(user.UserEmail); !ok {
	// 	return err
	// }
	return nil
}

func init() {
	db.SQL.AutoMigrate(&Routes{})
	db.SQL.AutoMigrate(&RoleRouteMapping{})
}
