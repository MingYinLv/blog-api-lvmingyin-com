package schema

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

var GetSpecialByIdQuery = &graphql.Field{
	Type: SpecialType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, idOK := params.Args["id"].(int)
		if !idOK {
			return nil, errors.New("id不能为空")
		}
		return FindSpecialById(int64(id))
	},
}

var GetSpecialsQuery = &graphql.Field{
	Type: SpecialListType,
	Args: graphql.FieldConfigArgument{
		"special_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		name, _ := params.Args["special_name"].(string)
		page, pOK := params.Args["page"].(int)
		size, sOK := params.Args["size"].(int)
		if !pOK {
			page = 1
		}
		if !sOK {
			size = 10
		}
		return FindSpecials(Special{SpecialName: name}, int64(page), int64(size))
	},
}

func FindSpecialById(speId int64) (interface{}, error) {
	return specialDao.QueryRow("SELECT * FROM special WHERE id = ?", speId)
}

func FindSpecials(special Special, page, size int64) (interface{}, error) {
	sql := "SELECT * FROM special"
	countSql := "SELECT COUNT(*) FROM special"
	var params []interface{}
	if strings.TrimSpace(special.SpecialName) != "" {
		sql = fmt.Sprintf("%s WHERE specialName like ?", sql)
		countSql = fmt.Sprintf("%s WHERE specialName like ?", countSql)
		params = append(params, "%"+special.SpecialName+"%")
	}

	sql = fmt.Sprintf("%s limit %d,%d", sql, (page-1)*size, size)
	result, err := specialDao.Query(sql, params...)

	if err != nil {
		return DBErrorLog("查询失败", err)
	}

	row, err := QueryRow(countSql, params...)
	var count int64
	err = row.Scan(&count)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}

	return ListResult{page, size, count, result}, nil
}
