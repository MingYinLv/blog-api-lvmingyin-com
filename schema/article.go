package schema

import "github.com/graphql-go/graphql"

type Article struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Cover    string `json:"cover"`
	TypeId   int64  `json:"type_id"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

var ActType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"cover": &graphql.Field{
			Type: graphql.String,
		},
		"type_id": &graphql.Field{
			Type: graphql.Int,
		},
		"create_at": &graphql.Field{
			Type: graphql.Int,
		},
		"update_at": &graphql.Field{
			Type: graphql.Int,
		},
		"type": &graphql.Field{
			Type: ActTType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return FindActTypeById(p.Source.(Article).TypeId)
			},
		},
	},
})

func init() {
	ActType.AddFieldConfig("tags", &graphql.Field{
		Type: graphql.NewList(TagType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return FindTagsByActId(p.Source.(Article).ID)
		},
	})
}
