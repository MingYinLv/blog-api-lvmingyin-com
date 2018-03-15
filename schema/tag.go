package schema

import "github.com/graphql-go/graphql"

type Tag struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
}

var TagType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Tag",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"tag_name": &graphql.Field{
			Type: graphql.String,
		},
		"articles": &graphql.Field{
			Type: graphql.NewList(ActType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return FindArticleListByTagId(p.Source.(Tag).ID)
			},
		},
	},
})
