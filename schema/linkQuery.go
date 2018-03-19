package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
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

func FindLinks(link Link) ([]Link, error) {
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
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		util.ErrorLog.Println(err)
		return []Link{}, errors.New(fmt.Sprintf("获取链接列表失败"))
	}
	defer stms.Close()

	rows, err := stms.Query(params...)

	if err != nil {
		util.ErrorLog.Println(err)
		return []Link{}, errors.New(fmt.Sprintf("获取链接列表失败"))
	}

	var result []Link
	for rows.Next() {
		var id, linkType int64
		var url, name, icon string
		err = rows.Scan(&id, &icon, &linkType, &url, &name)
		if err != nil {
			return []Link{}, errors.New(fmt.Sprintf("获取标签列表失败"))
		}
		result = append(result, Link{id, icon, linkType, url, name})
	}
	return result, nil

}

func FindLinkById(linkId int64) (Link, error) {
	stms, err := db.DB.Prepare("SELECT * FROM link where id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return Link{}, errors.New(fmt.Sprintf("获取链接信息失败"))
	}

	row := stms.QueryRow(linkId)
	stms.Close()
	var id, linkType int64
	var icon, url, name string
	err = row.Scan(&id, &icon, &linkType, &url, &name)
	if err != nil {
		return Link{}, errors.New(fmt.Sprintf("没有该链接"))
	}
	return Link{id, icon, linkType, url, name}, nil
}
