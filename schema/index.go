package schema

import "github.com/graphql-go/graphql"

var query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"login":           AdminLoginQuery,
		"information":     InformationQuery,
		"articleType":     GetActTypeByIdQuery,
		"articleTypeList": GetActTypeListQuery,
		"article":         GetArticleByIdQuery,
		"tag":             GetTagByIdQuery,
		"tags":            GetTagsQuery,
		"articles":        GetArticlesQuery,
		"articlesByType":  GetArticleByTypeIdQuery,
		"links":           GetLinksQuery,
		"link":            GetLinkQuery,
		"special":         GetSpecialByIdQuery,
		"specials":        GetSpecialsQuery,
	},
})

var mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"adminCreate":       AdminCreateMutation,
		"updateInformation": InformationMutation,
		"addArticleType":    AddActTypeMutation,
		"updateArticleType": UpdateActTypeMutation,
		"deleteArticleType": DeleteActTypeMutation,
		"addTag":            AddTagMutation,
		"deleteTag":         DeleteTagMutation,
		"updateTag":         UpdateTagMutation,
		"addArticle":        AddArticleMutation,
		"updateArticle":     UpdateArticleMutation,
		"deleteArticle":     DeleteArticleMutation,
		"addLink":           AddLinkMutation,
		"updateLink":        UpdateLinkMutation,
		"deleteLink":        DeleteLinkMutation,
		"addSpecial":        AddSpecialMutation,
		"deleteSpecial":     DeleteSpecialMutation,
		"updateSpecial":     UpdateSpecialMutation,
		"deleteSpecialActs": DeleteSpecialActsMutation,
	},
})

var RootSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})
