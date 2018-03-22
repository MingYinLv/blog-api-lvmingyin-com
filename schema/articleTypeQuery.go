package schema

import (
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
		typeName, nameOK := params.Args["type_name"].(string)
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
			sql = fmt.Sprintf("%s AND typeName like ?", sql)
			param = append(param, "%"+typeName+"%")
		}
		if menuOK {
			sql = fmt.Sprintf("%s AND showMenu = ?", sql)
			param = append(param, showMenu)
		}
		return FindActTypeList(sql, param...)
	},
}

func FindActTypeById(queryId int64) (interface{}, error) {
	return articleTypeDao.QueryRow("SELECT * FROM articleType where id = ?", queryId)
}

func FindActTypeByName(queryName string) (interface{}, error) {
	return articleTypeDao.QueryRow("SELECT * FROM articleType where typeName = ?", queryName)
}

func FindActTypeList(sql string, args ...interface{}) (interface{}, error) {
	return articleTypeDao.Query(sql, args...)
}
