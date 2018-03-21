package schema

import (
	"github.com/graphql-go/graphql"
)

type Link struct {
	ID   int64  `json:"id"`
	Icon string `json:"icon"`
	Type int64  `json:"type"` // 0 友情链接 1 其他个人空间
	URL  string `json:"url"`
	Name string `json:"name"`
}


var LinkType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Link",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"type": &graphql.Field{
			Type: graphql.Int,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
