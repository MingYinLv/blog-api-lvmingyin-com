package schema

import "github.com/graphql-go/graphql"

type ArticleType struct {
	ID       int64  `json:"id"`
	TypeName string `json:"type_name"`
	ShowMenu int64  `json:"show_menu"` // 0 不显示 1显示
	Logo     string `json:"logo"`
}

var ActTType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ArticleType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"type_name": &graphql.Field{
			Type: graphql.String,
		},
		"show_menu": &graphql.Field{
			Type: graphql.Int,
		},
		"logo": &graphql.Field{
			Type: graphql.String,
		},
	},
})
