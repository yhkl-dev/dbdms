package user

import (
	"github.com/graphql-go/graphql"
)

// UserType 定义查询对象的字段，支持嵌套
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "User Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"user_name": &graphql.Field{
			Type: graphql.String,
		},
		"user_phone": &graphql.Field{
			Type: graphql.String,
		},
		"user_email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryUser = graphql.Field{
	Name:        "QueryUser",
	Description: "Query User Info",
	Type:        graphql.NewList(UserType),
	// Args是定义在GraphQL查询中支持的查询字段，
	// 可自行随意定义，如加上limit,start这类
	Args: graphql.FieldConfigArgument{
		"user_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_phone": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"user_email": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	// Resolve是一个处理请求的函数，具体处理逻辑可在此进行
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		// Args里面定义的字段在p.Args里面，对应的取出来
		// 因为是interface{}的值，需要类型转换，可参考类型断言(type assertion):
		// https://zengweigang.gitbooks.io/core-go/content/eBook/11.3.html
		userName, _ := p.Args["user_name"].(string)
		userPhone, _ := p.Args["user_phone"].(string)
		userEmail, _ := p.Args["user_email"].(string)

		// 调用Hello这个model里面的Query方法查询数据
		return (&User{}).Query(userName, userPhone, userEmail), nil
	},
}

// 定义跟查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"user": &queryUser, // queryHello 参考schema/hello.go
	},
})

// Schema 定义Schema用于http handler处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, // 需要通过GraphQL更新数据，可以定义Mutation
})
