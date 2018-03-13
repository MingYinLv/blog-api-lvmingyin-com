package schema

import "github.com/graphql-go/graphql"

type Tags struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
}

var TagsType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tags",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"tag_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})
