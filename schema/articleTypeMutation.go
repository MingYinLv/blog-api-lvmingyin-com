package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

var InsertActTypeMutation = &graphql.Field{
	Type: ActTType,
	Args: graphql.FieldConfigArgument{
		"type_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"show_menu": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"logo": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		type_name, tOK := params.Args["type_name"].(string)
		show_menu, _ := params.Args["show_menu"].(int)
		logo, _ := params.Args["logo"].(string)
		if !tOK || strings.TrimSpace(type_name) == "" {
			return ArticleType{}, errors.New("请输入类型名称")
		}

		actType := ArticleType{TypeName: type_name, ShowMenu: int64(show_menu), Logo: logo}

		return AddArticleType(&actType)
	},
}

var UpdateActTypeMutation = &graphql.Field{
	Type: ActTType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"type_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"show_menu": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"logo": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		type_name, tOK := params.Args["type_name"].(string)
		show_menu, _ := params.Args["show_menu"].(int)
		id, idOK := params.Args["id"].(int)
		logo, _ := params.Args["logo"].(string)
		if !tOK || strings.TrimSpace(type_name) == "" {
			return ArticleType{}, errors.New("请输入类型名称")
		} else if !idOK {
			return ArticleType{}, errors.New("请输入id")
		}

		actType := ArticleType{ID: int64(id), TypeName: type_name, ShowMenu: int64(show_menu), Logo: logo}

		return UpdateArticleType(&actType)
	},
}

var DeleteActTypeMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if id, idOK := params.Args["id"].(int); idOK {
			return DeleteArticleType(int64(id))
		} else {
			return 0, errors.New("请输入要删除的的类型id")
		}
	},
}

func AddArticleType(articleType *ArticleType) (*ArticleType, error) {
	_, err := FindActTypeByName(articleType.TypeName)
	if err == nil {
		return &ArticleType{}, errors.New(fmt.Sprintf("类型 %s 已存在", articleType.TypeName))
	}
	stms, err := db.DB.Prepare("INSERT INTO articleType(typeName,showMenu,logo) values(?,?,?)")

	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型创建失败")
	}
	defer stms.Close()

	result, err := stms.Exec(articleType.TypeName, articleType.ShowMenu, articleType.Logo)
	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型创建失败")
	}

	id, err := result.LastInsertId()
	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型创建失败")
	}
	articleType.ID = id
	return articleType, nil

}

func UpdateArticleType(articleType *ArticleType) (*ArticleType, error) {
	_, err := FindActTypeById(articleType.ID)
	if err != nil {
		return &ArticleType{}, errors.New("修改的类型不存在")
	}

	actType, err := FindActTypeByName(articleType.TypeName)
	if err == nil && articleType.ID != actType.ID {
		// 能查到数据，并且id和当前修改的id不一样，不允许冲突
		return &ArticleType{}, errors.New(fmt.Sprintf("类型 %s 已存在", articleType.TypeName))
	} else if err != nil {
		return &ArticleType{}, errors.New("类型修改失败")
	}

	stms, err := db.DB.Prepare("UPDATE articleType SET typeName = ?,showMenu = ?,logo = ? WHERE id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型修改失败")
	}
	defer stms.Close()
	result, err := stms.Exec(articleType.TypeName, articleType.ShowMenu, articleType.Logo, articleType.ID)
	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型修改失败")
	}
	_, err = result.RowsAffected()
	if err != nil {
		util.ErrorLog.Println(err)
		return &ArticleType{}, errors.New("类型修改失败")
	}
	return articleType, nil
}

func DeleteArticleType(idQuery int64) (int64, error) {
	stms, err := db.DB.Prepare("DELETE FROM articleType WHERE id = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return 0, errors.New("类型删除失败")
	}
	defer stms.Close()

	result, err := stms.Exec(idQuery)
	if err != nil {
		util.ErrorLog.Println(err)
		return 0, errors.New("类型删除失败")
	}

	return result.RowsAffected()
}
