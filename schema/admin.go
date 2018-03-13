package schema

import "github.com/graphql-go/graphql"

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
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
		"salt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AdminLoginQuery = &graphql.Field{
	Type: AdminType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		username, _ := params.Args["username"].(string)
		password, _ := params.Args["password"].(string)
		if username == "lvmingyin" && password == "lvmingyin123" {
			return Admin{Username: username}, nil
		}
		return Admin{}, nil
	},
}
