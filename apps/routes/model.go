package routes

import (
	"dbdms/db"
	"fmt"
)

// Routes route struct
type Routes struct {
	RouteID     int
	RouteName   string
	RoutePath   string
	RouteMethod string
	RoutePID    int
	RouteIcon   string
	isShow      bool
}

type RoleRouteMapping struct {
	MappingID int
	RoleID    int
	RouteID   int
}

// TableName define table name in database
func (r *Routes) TableName() string {
	return "sys_routes"
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
