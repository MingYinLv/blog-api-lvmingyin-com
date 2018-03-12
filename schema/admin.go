package schema

import "github.com/graphql-go/graphql"

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var AdminType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Admin",
	Fields: graphql.Fields{
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
	},
})
