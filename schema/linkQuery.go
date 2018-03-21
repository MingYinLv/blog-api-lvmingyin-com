package schema

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

var GetLinks = &graphql.Field{
	Type: graphql.NewList(LinkType),
	Args: graphql.FieldConfigArgument{
		"type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"keyword": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		linkType, _ := params.Args["type"].(int)
		keyword, _ := params.Args["keyword"].(string)
		return FindLinks(Link{Type: int64(linkType), Name: keyword, URL: keyword})

	},
}

var GetLink = &graphql.Field{
	Type: LinkType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, _ := params.Args["id"].(int)
		return FindLinkById(int64(id))

	},
}

func FindLinks(link Link) (interface{}, error) {
	sql := "SELECT * FROM link WHERE 1 "
	var params []interface{}
	if strings.TrimSpace(link.Name) != "" && strings.TrimSpace(link.URL) != "" {
		sql = fmt.Sprintf("%s AND (name like ?", sql)
		params = append(params, "%"+link.Name+"%")

		sql = fmt.Sprintf("%s OR url like ?)", sql)
		params = append(params, "%"+link.URL+"%")
	}
	if link.Type != 0 {
		sql = fmt.Sprintf("%s AND type = ?", sql)
		params = append(params, link.Type)
	}
	return linkDao.Query(sql, params...)

}

func FindLinkById(linkId int64) (interface{}, error) {
	return linkDao.QueryRow("SELECT * FROM link where id = ?", linkId)
}
