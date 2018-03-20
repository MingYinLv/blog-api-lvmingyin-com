package schema

import (
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

var AddTagMutation = &graphql.Field{
	Type: TagType,
	Args: graphql.FieldConfigArgument{
		"tag_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		tag_name, tOK := params.Args["tag_name"].(string)
		if !tOK || strings.TrimSpace(tag_name) == "" {
			return ArticleType{}, errors.New("请输入标签名称")
		}

		tagType := Tag{TagName: tag_name}

		return AddTag(&tagType)
	},
}

var UpdateTagMutation = &graphql.Field{
	Type: TagType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"tag_name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		tag_name, tOK := p.Args["tag_name"].(string)
		id, idOK := p.Args["id"].(int)
		if !tOK || strings.TrimSpace(tag_name) == "" {
			return Tag{}, errors.New("请输入标签名称")
		} else if !idOK {
			return Tag{}, errors.New("请输入id")
		}
		return UpdateTag(&Tag{int64(id), tag_name})
	},
}

var DeleteTagMutation = &graphql.Field{
	Type: graphql.Int,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if id, idOK := params.Args["id"].(int); idOK {
			return DeleteTag(int64(id))
		} else {
			return 0, errors.New("请输入要删除的的标签id")
		}
	},
}

func AddTag(tag *Tag) (*Tag, error) {
	_, err := FindTagByName(tag.TagName)
	if err == nil {
		return &Tag{}, errors.New(fmt.Sprintf("标签 %s 已存在", tag.TagName))
	}

	id, err := ExecInsert("INSERT INTO tags(tag_name) values(?)", tag.TagName)
	if err != nil {
		util.ErrorLog.Println(err)
		return &Tag{}, errors.New("标签创建失败")
	}
	tag.ID = id
	return tag, nil

}

func DeleteTag(idQuery int64) (int64, error) {
	return ExecDelete("DELETE FROM tags WHERE id = ?", idQuery)
}

func UpdateTag(tag *Tag) (*Tag, error) {
	_, err := FindTagById(tag.ID)
	if err != nil {
		return &Tag{}, errors.New("修改的标签不存在")
	}

	actType, err := FindTagByName(tag.TagName)
	if err == nil && tag.ID != actType.(Tag).ID {
		// 能查到数据，并且id和当前修改的id不一样，不允许冲突
		return &Tag{}, errors.New(fmt.Sprintf("标签 %s 已存在", tag.TagName))
	}

	_, err = ExecUpdate("UPDATE tags SET tag_name = ? WHERE id = ?", tag.TagName, tag.ID)
	if err != nil {
		util.ErrorLog.Println(err)
		return &Tag{}, errors.New("标签修改失败")
	}
	return tag, nil
}
