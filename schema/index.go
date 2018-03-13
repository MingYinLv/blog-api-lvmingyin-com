package schema

import "github.com/graphql-go/graphql"

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"login": AdminLoginQuery,
	},
})


var RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: query,
})
