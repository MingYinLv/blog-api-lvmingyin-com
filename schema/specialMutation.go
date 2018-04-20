package schema

import (
	"blog-api-lvmingyin-com/db"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

var AddSpecialMutation = &graphql.Field{
	Type: SpecialType,
	Args: graphql.FieldConfigArgument{
		"special_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cover": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		specialName, tOK := params.Args["special_name"].(string)
		cover, _ := params.Args["cover"].(string)
		if !tOK || strings.TrimSpace(specialName) == "" {
			return nil, errors.New("请输入专题名称")
		}

		special := Special{SpecialName: specialName, Cover: cover}

		return AddSpecial(&special)
	},
}

var UpdateSpecialMutation = &graphql.Field{
	Type: SpecialType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"special_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"cover": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		specialName, tOK := p.Args["special_name"].(string)
		cover, _ := p.Args["cover"].(string)
		id, idOK := p.Args["id"].(int)
		if !tOK || strings.TrimSpace(specialName) == "" {
			return nil, errors.New("请输入专题名称")
		} else if !idOK {
			return nil, errors.New("请输入id")
		}
		return UpdateSpecial(&Special{int64(id), specialName, cover})
	},
}

var DeleteSpecialMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if id, idOK := params.Args["id"].(int); idOK {
			return DeleteSpecial(int64(id))
		} else {
			return 0, errors.New("请输入要删除的的专题id")
		}
	},
}

var DeleteSpecialActsMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"ids": &graphql.ArgumentConfig{
			Type: graphql.NewList(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if ids, idOK := params.Args["ids"].([]int); idOK {
			return DeleteSpecialActs(ids)
		} else {
			return 0, errors.New("请输入要删除的的专题id")
		}
	},
}

func AddSpecial(special *Special) (interface{}, error) {
	_, err := FindSpecialByName(special.SpecialName)
	if err == nil {
		return nil, errors.New(fmt.Sprintf("专题 %s 已存在", special.SpecialName))
	}

	return specialDao.Insert("INSERT INTO special(special_name, cover) values(?, ?)", special, special.SpecialName, special.Cover)
}

func UpdateSpecial(special *Special) (interface{}, error) {
	_, err := FindSpecialById(special.ID)
	if err != nil {
		return nil, errors.New("修改的专题不存在")
	}
	spe, err := FindSpecialByName(special.SpecialName)
	if err == nil && special.ID != spe.(Special).ID {
		// 能查到数据，并且id和当前修改的id不一样，不允许冲突
		return nil, errors.New(fmt.Sprintf("专题 %s 已存在", special.SpecialName))
	}
	return specialDao.Update("UPDATE special SET special_name = ?, cover = ? WHERE id = ?", special, special.SpecialName, special.Cover, special.ID)
}

func DeleteSpecial(speId int64) (int64, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}

	stms, err := tx.Prepare("DELETE FROM special WHERE id = ?")
	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}
	defer stms.Close()
	result, err := stms.Exec(speId)
	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}
	row, err := result.RowsAffected()
	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}

	mappStmt, err := tx.Prepare("DELETE FROM speMappAct WHERE spe_id = ?")
	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}

	result, err = mappStmt.Exec(speId)

	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}

	mappRow, err := result.RowsAffected()

	if err != nil {
		DBLog(err)
		tx.Rollback()
		return 0, errors.New("删除失败")
	}

	tx.Commit()
	return mappRow + row, nil
}

func DeleteSpecialActs(ids []int) (int64, error) {
	length := len(ids)
	if length >= 0 {
		return 0, DBNewTextError("请选择要删除的文章")
	}
	sql := "DELETE FROM speMappAct WHERE"
	var params []interface{}
	for i, v := range ids {
		if i > 0 {
			sql = fmt.Sprintf(" %s OR", sql)
		}
		sql = fmt.Sprintf(" %s act_id = ?", sql)
		params = append(params, v)
	}
	return specialDao.Delete(sql, params...)
}
