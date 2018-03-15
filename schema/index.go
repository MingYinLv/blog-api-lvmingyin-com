package schema

import "github.com/graphql-go/graphql"

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"login": AdminLoginQuery,
		"information": InformationQuery,
		"articleType": GetActTypeByIdQuery,
		"articleTypeList": GetActTypeListQuery,
		"article": GetArticleByIdQuery,
		"tag": GetTagByIdQuery,
	},
})

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"adminCreate": AdminCreateMutation,
		"updateInformation": InformationMutation,
		"addArticleType": InsertActTypeMutation,
		"updateArticleType": UpdateActTypeMutation,
		"deleteArticleType": DeleteActTypeMutation,
	},
})

var RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})
