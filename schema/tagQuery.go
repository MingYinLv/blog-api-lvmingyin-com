package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

var GetTagsQuery = &graphql.Field{
	Type: TagListType,
	Args: graphql.FieldConfigArgument{
		"ids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		ids, _ := params.Args["ids"].([]interface{})
		page, pOK := params.Args["page"].(int)
		size, sOK := params.Args["size"].(int)
		if !pOK {
			page = 1
		}
		if !sOK {
			size = 10
		}
		return FindTags(ids, page, size)

	},
}

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

func FindTagById(queryId int64) (interface{}, error) {

	return tagDao.QueryRow("SELECT * FROM tags where id = ?", queryId)
}

func FindTagByName(tagName string) (interface{}, error) {
	return tagDao.QueryRow("SELECT * FROM tags where tag_name = ?", tagName)
}

func FindTagsByActId(actId int64) (interface{}, error) {
	return tagDao.Query("SELECT tags.id,tags.tag_name from actMappTag right join tags on tags.id = actMappTag.tag_id where actMappTag.act_id = ?", actId)
}

func FindTags(ids []interface{}, page, size int) (ListResult, error) {
	sql := util.GenInKeys("tags", "id", ids, page, size)
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		util.ErrorLog.Println(err)
		return ListResult{}, errors.New(fmt.Sprintf("获取标签列表失败"))
	}
	defer stms.Close()

	rows, err := stms.Query(ids...)

	if err != nil {
		util.ErrorLog.Println(err)
		return ListResult{}, errors.New(fmt.Sprintf("获取标签列表失败"))
	}

	var result []interface{}
	for rows.Next() {
		var id int64
		var tag_name string
		err = rows.Scan(&id, &tag_name)
		if err != nil {
			return ListResult{}, errors.New(fmt.Sprintf("获取标签列表失败"))
		}
		result = append(result, Tag{id, tag_name})

	}

	stmsTotal, err := db.DB.Prepare("SELECT count(id) FROM tags")
	if err != nil {
		util.ErrorLog.Println(err)
		return ListResult{}, errors.New(fmt.Sprintf("获取标签列表失败"))
	}
	defer stmsTotal.Close()
	row := stmsTotal.QueryRow()

	var total int64
	err = row.Scan(&total)
	if err != nil {
		util.ErrorLog.Println(err)
		return ListResult{}, errors.New(fmt.Sprintf("获取标签列表失败"))
	}

	tagList := ListResult{
		int64(page), int64(size), total, result,
	}
	return tagList, nil
}
