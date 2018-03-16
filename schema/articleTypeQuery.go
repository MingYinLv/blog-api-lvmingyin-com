package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

var GetActTypeByIdQuery = &graphql.Field{
	Type: ActTType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, idOK := params.Args["id"].(int)
		if !idOK {
			return nil, errors.New("请输入类型id")
		}
		return FindActTypeById(int64(id))
	},
}

var GetActTypeListQuery = &graphql.Field{
	Type: graphql.NewList(ActTType),
	Args: graphql.FieldConfigArgument{
		"ids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"type_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"show_menu": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		typeName, nameOK := params.Args["type_name"]
		showMenu, menuOK := params.Args["show_menu"]
		ids, idsOK := params.Args["ids"].([]interface{})
		sql := "SELECT * FROM articleType WHERE 1 = 1"
		var param []interface{}
		if idsOK {
			if len(ids) > 0 {
				sql = fmt.Sprintf("%s AND id in (", sql)
				for i := range ids {
					sql = fmt.Sprintf("%s ?", sql)
					if i < len(ids)-1 {
						sql = fmt.Sprintf("%s,", sql)
					}
					param = append(param, ids[i])
				}
				sql = fmt.Sprintf("%s)", sql)
			}
		}
		if nameOK {
			sql = fmt.Sprintf("%s AND typeName = ?", sql)
			param = append(param, typeName)
		}
		if menuOK {
			sql = fmt.Sprintf("%s AND showMenu = ?", sql)
			param = append(param, showMenu)
		}
		return FindActTypeList(sql, param...)
	},
}

func FindActTypeById(queryId int64) (ArticleType, error) {
	stms, err := db.DB.Prepare("SELECT * FROM articleType where id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return ArticleType{}, errors.New(fmt.Sprintf("获取分类信息失败"))
	}

	row := stms.QueryRow(queryId)
	stms.Close()
	var id, showMenu int64
	var typeName, logo string
	err = row.Scan(&id, &typeName, &showMenu, &logo)
	if err != nil {
		return ArticleType{}, errors.New(fmt.Sprintf("没有该分类"))
	}
	return ArticleType{id, typeName, showMenu, logo}, nil
}

func FindActTypeByName(queryName string) (ArticleType, error) {
	stms, err := db.DB.Prepare("SELECT * FROM articleType where typeName = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return ArticleType{}, errors.New(fmt.Sprintf("获取分类信息失败"))
	}

	row := stms.QueryRow(queryName)
	stms.Close()
	var id, showMenu int64
	var typeName, logo string
	err = row.Scan(&id, &typeName, &showMenu, &logo)
	if err != nil {
		return ArticleType{}, errors.New(fmt.Sprintf("没有 %s 分类", queryName))
	}
	return ArticleType{id, typeName, showMenu, logo}, nil
}

func FindActTypeList(sql string, args ...interface{}) ([]ArticleType, error) {
	stms, err := db.DB.Prepare(sql)
	var result []ArticleType
	if err != nil {
		util.ErrorLog.Println(err)
		return []ArticleType{}, errors.New(fmt.Sprintf("获取分类列表失败"))
	}
	defer stms.Close()
	rows, err := stms.Query(args...)
	if err != nil {
		util.ErrorLog.Println(err)
		return []ArticleType{}, errors.New(fmt.Sprintf("获取分类列表失败"))
	}

	for rows.Next() {
		var id, showMenu int64
		var typeName, logo string
		err = rows.Scan(&id, &typeName, &showMenu, &logo)
		if err != nil {
			util.ErrorLog.Println(err)
			return result, errors.New(fmt.Sprintf("获取分类列表失败"))
		}
		result = append(result, ArticleType{id, typeName, showMenu, logo})
	}
	return result, nil
}
