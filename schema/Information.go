package schema

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

type Information struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Subtitle    string `json:"subtitle"`
	RealName    string `json:"realname"`
	Logo        string `json:"logo"`
	Email       string `json:"email"`
	QQ          string `json:"qq"`
	Telephone   string `json:"telephone"`
	Copyright   string `json:"copyright"`
	ICP         string `json:"icp"`
}

var InformationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Information",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"subtitle": &graphql.Field{
			Type: graphql.String,
		},
		"realname": &graphql.Field{
			Type: graphql.String,
		},
		"logo": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"qq": &graphql.Field{
			Type: graphql.String,
		},
		"telephone": &graphql.Field{
			Type: graphql.String,
		},
		"copyright": &graphql.Field{
			Type: graphql.String,
		},
		"icp": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var InformationInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "informationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"id": &graphql.InputObjectFieldConfig{
			Type: graphql.Int,
		},
		"title": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"subtitle": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"realname": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"logo": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"qq": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"telephone": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"copyright": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"icp": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var InformationQuery = &graphql.Field{
	Type: InformationType,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return GetInformation()
	},
}

var InformationMutation = &graphql.Field{
	Type: InformationType,
	Args: graphql.FieldConfigArgument{
		"information": &graphql.ArgumentConfig{
			Type: InformationInput,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		maps := params.Args["information"].(map[string]interface{})
		information := Information{}
		t := reflect.TypeOf(information)
		v := reflect.ValueOf(&information).Elem()

		for k := 0; k < t.NumField(); k++ {
			name := t.Field(k).Name
			lowerName := strings.ToLower(name)
			val, ok := maps[lowerName]
			if lowerName == "id" {
				val = int64(val.(int))
			}
			if ok {
				v.FieldByName(name).Set(reflect.ValueOf(val))
			}
		}

		return UpdateInformation(information)
	},
}

func GetInformation() (interface{}, error) {
	row, err := QueryRow("SELECT * FROM information order by id desc limit 0,1")
	if err != nil {
		return DBErrorLog("获取网站信息失败", err)
	}

	var id int64
	var title, description, icon, subtitle, realname, logo, email, qq, telephone, copyright, icp string
	err = row.Scan(&id, &title, &description, &icon, &subtitle, &realname, &logo, &email, &qq, &telephone, &copyright, &icp)
	if err != nil {
		return DBErrorLog("获取网站信息失败", err)
	}
	return Information{id, title, description, icon, subtitle, realname, logo, email, qq, telephone, copyright, icp}, nil
}

func UpdateInformation(information Information) (interface{}, error) {
	_, err := ExecUpdate("update information set title=?,logo=?,realname=?,subtitle=?,qq=?,icp=?,email=?,telephone=?,copyright=?,description=?,icon=? where id=?", information.Title, information.Logo, information.RealName, information.Subtitle, information.QQ, information.ICP, information.Email, information.Telephone, information.Copyright, information.Description, information.Icon, information.ID)
	if err != nil {
		return DBErrorLog("类型修改失败", err)
	}
	return information, nil
}
