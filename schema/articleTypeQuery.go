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
		id := int64(params.Args["id"].(int))
		return FindActTypeById(id)
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
