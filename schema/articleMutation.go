package schema

import (
	"blog-api-lvmingyin-com/db"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
	"time"
)

var AddArticleMutation = &graphql.Field{
	Type: ActType,
	Args: graphql.FieldConfigArgument{
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cover": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"tags": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		title, tOK := params.Args["title"].(string)
		content, cOK := params.Args["content"].(string)
		cover, _ := params.Args["cover"].(string)
		type_id, tiOK := params.Args["type_id"].(int)
		tags, _ := params.Args["tags"].([]interface{})
		if !tOK || strings.TrimSpace(title) == "" {
			return ArticleType{}, errors.New("请输入文章标题")
		} else if !cOK || strings.TrimSpace(content) == "" {
			return ArticleType{}, errors.New("请输入文章内容")
		} else if !tiOK || type_id < 1 {
			return ArticleType{}, errors.New("请选择文章分类")
		}

		article := Article{
			Title:   title,
			Content: content,
			Cover:   cover,
			TypeId:  int64(type_id),
		}

		return AddArticle(&article, &tags)
	},
}

var UpdateArticleMutation = &graphql.Field{
	Type: ActType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cover": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"type_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		title, tOK := params.Args["title"].(string)
		content, cOK := params.Args["content"].(string)
		cover, _ := params.Args["cover"].(string)
		type_id, tiOK := params.Args["type_id"].(int)
		id, idOK := params.Args["id"].(int)

		if !idOK {
			return ArticleType{}, errors.New("请输入id")
		} else if !tOK || strings.TrimSpace(title) == "" {
			return ArticleType{}, errors.New("请输入文章标题")
		} else if !cOK || strings.TrimSpace(content) == "" {
			return ArticleType{}, errors.New("请输入文章内容")
		} else if !tiOK || type_id < 1 {
			return ArticleType{}, errors.New("请选择文章分类")
		}

		article := Article{
			ID:      int64(id),
			Title:   title,
			Content: content,
			Cover:   cover,
			TypeId:  int64(type_id),
		}

		return UpdateArticle(&article)
	},
}

var DeleteArticleMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if id, idOK := params.Args["id"].(int); idOK {
			return DeleteArticle(int64(id))
		} else {
			return 0, errors.New("请输入要删除的的文章id")
		}
	},
}

func AddArticle(article *Article, tags *[]interface{}) (interface{}, error) {

	tx, err := db.DB.Begin()
	stms, err := tx.Prepare("INSERT INTO article(title,content,cover,type_id,update_at,create_at) values(?,?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return DBErrorLog("文章创建失败",err)
	}
	defer stms.Close()

	result, err := stms.Exec(article.Title, article.Content, article.Cover, article.TypeId, time.Now().Unix(), time.Now().Unix())

	if err != nil {
		tx.Rollback()
		return DBErrorLog("文章创建失败",err)
	}

	id, err := result.LastInsertId()
	article.ID = id

	if err != nil {
		tx.Rollback()
		return DBErrorLog("文章创建失败",err)
	}

	if len(*tags) > 0 {
		var execParam []interface{}
		sql := "INSERT INTO actMappTag(act_id,tag_id) VALUES"
		for i, v := range *tags {
			sql = fmt.Sprintf("%s(?, ?)", sql)
			if i < len(*tags)-1 {
				sql = fmt.Sprintf("%s,", sql)
			}
			execParam = append(execParam, id, v)
		}
		stmsTag, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return DBErrorLog("文章创建失败",err)
		}

		_, err = stmsTag.Exec(execParam...)
		if err != nil {
			tx.Rollback()
			return DBErrorLog("文章创建失败",err)
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return DBErrorLog("文章创建失败",err)
	}
	return *article, nil

}

func UpdateArticle(article *Article) (interface{}, error) {
	_, err := FindArticleById(article.ID)
	if err != nil {
		return DBErrorLog("修改的文章不存在", err)
	}

	stms, err := db.DB.Prepare("UPDATE article SET title = ?, content = ?, cover=?,type_id=?,update_at=? WHERE id = ?")
	if err != nil {
		return DBErrorLog("文章修改失败", err)
	}
	defer stms.Close()

	timeNow := time.Now().Unix()
	_, err = articleDao.Update("UPDATE article SET title = ?, content = ?, cover=?,type_id=?,update_at=? WHERE id = ?", article, article.Title, article.Content, article.Cover, article.TypeId, timeNow, article.ID)

	if err != nil {
		return DBErrorLog("文章修改失败", err)
	}
	article.UpdateAt = timeNow
	return *article, nil
}

func DeleteArticle(articleId int64) (int64, error) {

	tx, err := db.DB.Begin()
	stms, err := tx.Prepare("DELETE FROM article WHERE id=?")
	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}
	defer stms.Close()

	result, err := stms.Exec(articleId)

	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}

	row, err := result.RowsAffected()

	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}

	sql := "DELETE FROM actMappTag WHERE act_id = ?"
	stmsTag, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}

	result, err = stmsTag.Exec(articleId)
	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}

	tagRow, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		DBLog(err)
		return 0, DBNewTextError("删除失败")
	}
	return row + tagRow, nil
}
