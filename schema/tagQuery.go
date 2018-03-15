package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

var GetTagByIdQuery = &graphql.Field{
	Type: TagType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := int64(params.Args["id"].(int))
		return FindTagById(id)
	},
}

func FindTagById(queryId int64) (Tag, error) {
	stms, err := db.DB.Prepare("SELECT * FROM tags where id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return Tag{}, errors.New(fmt.Sprintf("获取标签信息失败"))
	}

	row := stms.QueryRow(queryId)
	stms.Close()
	var id int64
	var tag_name string
	err = row.Scan(&id, &tag_name)
	if err != nil {
		return Tag{}, errors.New(fmt.Sprintf("没有该标签"))
	}
	return Tag{id, tag_name}, nil
}

func FindTagsByActId(actId int64) ([]Tag, error) {
	stms, err := db.DB.Prepare("SELECT tags.id,tags.tagName from actMappTag right join tags on tags.id = actMappTag.tag_id where actMappTag.act_id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return []Tag{}, errors.New(fmt.Sprintf("获取文章标签失败"))
	}
	defer stms.Close()

	rows, err := stms.Query(actId)

	if err != nil {
		util.ErrorLog.Println(err)
		return []Tag{}, errors.New(fmt.Sprintf("获取文章标签失败"))
	}

	var result []Tag
	for rows.Next() {
		var id int64
		var tag_name string
		err = rows.Scan(&id, &tag_name)
		if err != nil {
			return []Tag{}, errors.New(fmt.Sprintf("获取文章标签失败"))
		}
		result = append(result, Tag{id, tag_name})

	}
	return result, nil
}
