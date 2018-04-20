package schema

import (
	"blog-api-lvmingyin-com/util"
	"github.com/graphql-go/graphql"
)

var GetArticlesQuery = &graphql.Field{
	Type: ActListType,
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
		return FindArticles(ids, page, size)
	},
}

var GetArticleByTypeIdQuery = &graphql.Field{
	Type: ActListType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"page": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"size": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, idOK := params.Args["id"].(int)
		page, pOK := params.Args["page"].(int)
		size, sOK := params.Args["size"].(int)
		if !idOK {
			return DBError("请传入ID")
		}
		if !pOK {
			page = 1
		}
		if !sOK {
			size = 10
		}
		return FindArticleListByTypeId(int64(id), page, size)
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

func FindArticleById(queryId int64) (interface{}, error) {
	return articleDao.QueryRow("SELECT * FROM article where id = ?", queryId)
}

func FindArticleListByTagId(queryId int64) (interface{}, error) {
	return articleDao.Query("SELECT article.id,article.title,article.content,article.cover,article.type_id,article.create_at,article.update_at FROM`article`INNER JOIN actMappTag ON article.id=actMappTag.act_id INNER JOIN tags ON actMappTag.tag_id=tags.id WHERE tags.id=?", queryId)
}

func FindArticleListByTypeId(typeId int64, page, size int) (interface{}, error) {
	r, err := articleDao.Query("SELECT * FROM article WHERE type_id=? limit ?,?", typeId, (page-1)*size, size)
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}
	result := r.([]Article)

	row, err := QueryRow("SELECT count(id) FROM article")
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}

	var total int64
	err = row.Scan(&total)
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}

	actList := ListResult{
		int64(page), int64(size), total, result,
	}
	return actList, nil
}

func FindArticlesBySpecialId(speId int64) (interface{}, error) {
	return articleDao.Query("SELECT article.id,article.title,article.content,article.cover,article.type_id,article.create_at,article.update_at FROM`special`INNER JOIN speMappAct ON special.id=speMappAct.spe_id INNER JOIN article ON speMappAct.act_id=special.id WHERE special.id=?", speId)
}

func FindArticles(ids []interface{}, page, size int) (interface{}, error) {
	sql := util.GenInKeys("article", "id", ids, page, size)
	r, err := articleDao.Query(sql, ids...)
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}
	result := r.([]Article)

	row, err := QueryRow("SELECT count(id) FROM article")
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}

	var total int64
	err = row.Scan(&total)
	if err != nil {
		return DBErrorLog("获取文章列表失败", err)
	}

	actList := ListResult{
		int64(page), int64(size), total, result,
	}
	return actList, nil
}
