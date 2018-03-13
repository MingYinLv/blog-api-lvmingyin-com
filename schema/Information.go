package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
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

var InformationQuery = &graphql.Field{
	Type: InformationType,
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return GetInformation()
	},
}

func GetInformation() (Information, error) {
	stms, err := db.DB.Prepare("SELECT * FROM information order by id desc limit 0,1")
	if err != nil {
		util.ErrorLog.Println(err)
		return Information{}, errors.New(fmt.Sprintf("获取网站信息失败"))
	}

	row := stms.QueryRow()
	stms.Close()
	var id int64
	var title, description, icon, subtitle, realname, logo, email, qq, telephone, copyright, icp string
	err = row.Scan(&id, &title, &description, &icon, &subtitle, &realname, &logo, &email, &qq, &telephone, &copyright, &icp)
	if err != nil {
		return Information{}, errors.New(fmt.Sprintf("获取网站信息失败"))
	}
	return Information{id, title, description, icon, subtitle, realname, logo, email, qq, telephone, copyright, icp}, nil
}
