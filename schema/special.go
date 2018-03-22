package schema

import "github.com/graphql-go/graphql"

type Special struct {
	ID          int64  `json:"id"`
	SpecialName string `json:"special_name"`
	Cover       string `json:"cover"`
}

var SpecialType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Special",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"special_name": &graphql.Field{
			Type: graphql.String,
		},
		"cover": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var SpecialListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "SpecialList",
	Fields: graphql.Fields{
		"page": &graphql.Field{
			Type: graphql.Int,
		},
		"size": &graphql.Field{
			Type: graphql.Int,
		},
		"total": &graphql.Field{
			Type: graphql.Int,
		},
		"data": &graphql.Field{
			Type: graphql.NewList(SpecialType),
		},
	},
})

func init() {
	SpecialType.AddFieldConfig("articles", &graphql.Field{
		Type: graphql.NewList(ActType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return FindArticlesBySpecialId(params.Source.(Special).ID)
		},
	})
}
