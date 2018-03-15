package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
)

var GetArticles = &graphql.Field{
	Type: graphql.NewList(ActType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return FindArticles()
	},
}

var GetArticleByIdQuery = &graphql.Field{
	Type: ActType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id := int64(params.Args["id"].(int))
		return FindArticleById(id)
	},
}

func FindArticleById(queryId int64) (Article, error) {
	stms, err := db.DB.Prepare("SELECT * FROM article where id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return Article{}, errors.New(fmt.Sprintf("获取文章信息失败"))
	}

	row := stms.QueryRow(queryId)
	stms.Close()
	var id, type_id, create_at, update_at int64
	var title, content, cover string
	err = row.Scan(&id, &title, &content, &cover, &type_id, &create_at, &update_at)
	if err != nil {
		return Article{}, errors.New(fmt.Sprintf("没有该文章"))
	}
	return Article{id, title, content, cover, type_id, create_at, update_at}, nil
}

func FindArticleListByTagId(queryId int64) ([]Article, error) {
	stms, err := db.DB.Prepare("SELECT article.id,article.title,article.content,article.cover,article.type_id,article.create_at,article.update_at FROM`article`INNER JOIN actMappTag ON article.id=actMappTag.act_id INNER JOIN tags ON actMappTag.tag_id=tags.id WHERE tags.id=? GROUP BY article.id")
	if err != nil {
		util.ErrorLog.Println(err)
		return []Article{}, errors.New(fmt.Sprintf("获取文章信息失败"))
	}
	defer stms.Close()

	rows, err := stms.Query(queryId)
	if err != nil {
		util.ErrorLog.Println(err)
		return []Article{}, errors.New(fmt.Sprintf("获取文章信息失败"))
	}

	var result []Article
	for rows.Next() {
		var id, type_id, create_at, update_at int64
		var title, content, cover string
		err = rows.Scan(&id, &title, &content, &cover, &type_id, &create_at, &update_at)
		if err != nil {
			return []Article{}, errors.New(fmt.Sprintf("没有该文章"))
		}
		result = append(result, Article{id, title, content, cover, type_id, create_at, update_at})
	}
	return result, nil
}


func FindArticles() ([]Article, error) {
	stms, err := db.DB.Prepare("SELECT * FROM article")
	if err != nil {
		util.ErrorLog.Println(err)
		return []Article{}, errors.New(fmt.Sprintf("获取文章列表失败"))
	}
	defer stms.Close()

	rows, err := stms.Query()
	if err != nil {
		util.ErrorLog.Println(err)
		return []Article{}, errors.New(fmt.Sprintf("获取文章列表失败"))
	}

	var result []Article
	for rows.Next() {
		var id, type_id, create_at, update_at int64
		var title, content, cover string
		err = rows.Scan(&id, &title, &content, &cover, &type_id, &create_at, &update_at)
		if err != nil {
			return []Article{}, errors.New(fmt.Sprintf("获取文章列表失败"))
		}
		result = append(result, Article{id, title, content, cover, type_id, create_at, update_at})
	}
	return result, nil
}
