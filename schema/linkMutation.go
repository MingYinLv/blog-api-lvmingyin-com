package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
)

var AddLinkMutation = &graphql.Field{
	Type: LinkType,
	Args: graphql.FieldConfigArgument{
		"icon": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"url": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		icon, _ := p.Args["icon"].(string)
		linkType, _ := p.Args["type"].(int)
		url, _ := p.Args["url"].(string)
		name, _ := p.Args["name"].(string)
		return AddLink(Link{Icon: icon, Type: int64(linkType), URL: url, Name: name})
	},
}

var UpdateLinkMutation = &graphql.Field{
	Type: LinkType,
	Args: graphql.FieldConfigArgument{
		"icon": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"type": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"url": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		icon, _ := p.Args["icon"].(string)
		linkType, _ := p.Args["type"].(int)
		id, _ := p.Args["id"].(int)
		url, _ := p.Args["url"].(string)
		name, _ := p.Args["name"].(string)
		return UpdateLink(Link{ID: int64(id), Icon: icon, Type: int64(linkType), URL: url, Name: name})
	},
}

var DeleteLinkMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if id, idOK := params.Args["id"].(int); idOK {
			return DeleteLink(int64(id))
		} else {
			return 0, errors.New("请输入要删除的的链接id")
		}
	},
}

func AddLink(link Link) (interface{}, error) {
	return linkDao.Insert("INSERT INTO link(icon,type,url,name) values(?,?,?,?)", link, link.Icon, link.Type, link.URL, link.Name)

}

func DeleteLink(linkId int64) (int64, error) {
	return linkDao.Delete("DELETE FROM link WHERE id = ?", linkId)
}

func UpdateLink(link Link) (interface{}, error) {
	return linkDao.Update("UPDATE link SET icon = ?, url = ?, type=?, name=? WHERE id = ?", link, link.Icon, link.URL, link.Type, link.Name, link.ID)
}
