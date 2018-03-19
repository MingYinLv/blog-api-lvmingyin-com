package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
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
		return AddLink(&Link{Icon: icon, Type: int64(linkType), URL: url, Name: name})
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
		return UpdateLink(&Link{ID: int64(id), Icon: icon, Type: int64(linkType), URL: url, Name: name})
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

func AddLink(link *Link) (*Link, error) {
	stms, err := db.DB.Prepare("INSERT INTO link(icon,type,url,name) values(?,?,?,?)")

	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("链接创建失败")
	}
	defer stms.Close()

	result, err := stms.Exec(link.Icon, link.Type, link.URL, link.Name)
	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("链接创建失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("链接创建失败")
	}
	link.ID = id
	return link, nil

}

func DeleteLink(linkId int64) (int64, error) {
	stms, err := db.DB.Prepare("DELETE FROM link WHERE id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return 0, errors.New("链接删除失败")
	}
	defer stms.Close()

	result, err := stms.Exec(linkId)
	if err != nil {
		util.ErrorLog.Println(err)
		return 0, errors.New("链接删除失败")
	}

	return result.RowsAffected()
}

func UpdateLink(link *Link) (*Link, error) {
	_, err := FindLinkById(link.ID)
	if err != nil {
		return &Link{}, errors.New("修改的链接不存在")
	}

	stms, err := db.DB.Prepare("UPDATE link SET icon = ?, url = ?, type=?, name=? WHERE id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("修改的链接不存在")
	}
	defer stms.Close()
	result, err := stms.Exec(link.Icon, link.URL, link.Type, link.Name, link.ID)
	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("修改的链接不存在")
	}
	_, err = result.RowsAffected()
	if err != nil {
		util.ErrorLog.Println(err)
		return &Link{}, errors.New("修改的链接不存在")
	}
	return link, nil
}
