package schema

import "github.com/graphql-go/graphql"

type Special struct {
	ID          int64  `json:"id"`
	SpecialName string `json:"special_name"`
	Logo        string `json:"logo"`
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
		"logo": &graphql.Field{
			Type: graphql.String,
		},
	},
})
