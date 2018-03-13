package schema

import "github.com/graphql-go/graphql"

type Information struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Subtitle    string `json:"subtitle"`
	RealName    string `json:"realname"`
	Logo        string `json:"logo"`
	Email       string `json:"email"`
	QQ          string `json:"qq"`
	Telephone   string `json:"telephone"`
	Copyright   string `json:"copyright"`
	ICP         string `json:"icp"`
}

var InformationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Information",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"subtitle": &graphql.Field{
			Type: graphql.String,
		},
		"realname": &graphql.Field{
			Type: graphql.String,
		},
		"logo": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"qq": &graphql.Field{
			Type: graphql.String,
		},
		"telephone": &graphql.Field{
			Type: graphql.String,
		},
		"copyright": &graphql.Field{
			Type: graphql.String,
		},
		"icp": &graphql.Field{
			Type: graphql.String,
		},
	},
})
