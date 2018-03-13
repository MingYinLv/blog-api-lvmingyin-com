package schema

import (
	"blog-api-lvmingyin-com/db"
	"blog-api-lvmingyin-com/util"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"strings"
)

type Admin struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"_"`
	Salt     string `json:"_"`
}

var AdminType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Admin",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"salt": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AdminLoginQuery = &graphql.Field{
	Type: AdminType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		username, _ := params.Args["username"].(string)
		password, _ := params.Args["password"].(string)
		return AdminLogin(username, password)
	},
}

var AdminCreateMutation = &graphql.Field{
	Type: AdminType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		username, uOK := params.Args["username"].(string)
		password, pOK := params.Args["password"].(string)
		if !uOK || strings.TrimSpace(username) == "" {
			return Admin{}, errors.New("请输入用户名")
		} else if !pOK || strings.TrimSpace(password) == "" {
			return Admin{}, errors.New("请输入密码")
		} else if len(strings.TrimSpace(password)) < 6 {
			return Admin{}, errors.New("密码长度必须必须大于6位数")
		}
		return AdminCreate(username, password)
	},
}

func FindAdminByUserName(username string) (Admin, error) {
	stms, err := db.DB.Prepare("SELECT * FROM admin WHERE username = ?")
	if err != nil {
		util.ErrorLog.Println(err)
		return Admin{}, errors.New(fmt.Sprintf("用户 %s 不存在", username))
	}

	row := stms.QueryRow(username)
	stms.Close()
	var id int64
	var userResult, pwdResult, salt string
	err = row.Scan(&id, &userResult, &pwdResult, &salt)
	if err != nil {
		return Admin{}, errors.New(fmt.Sprintf("用户 %s 不存在", username))
	}
	return Admin{id, userResult, pwdResult, salt}, nil
}

func AdminLogin(username, password string) (Admin, error) {
	admin, err := FindAdminByUserName(username)
	if err != nil {
		// 其他错误
		return admin, err
	}

	if admin.Password == util.GetSha256Password(password, admin.Salt) {
		// 密码加密有一致
		return Admin{ID: admin.ID, Username: admin.Username}, nil
	} else {
		// 密码错误
		return Admin{}, errors.New("密码错误")
	}

}

func AdminCreate(username, password string) (Admin, error) {

	admin, err := FindAdminByUserName(username)
	if err == nil && admin.ID > 0 {
		return Admin{}, errors.New("该用户名已注册")
	}

	stms, err := db.DB.Prepare("INSERT INTO admin(username,password,salt) values(?,?,?)")
	if err != nil {
		util.ErrorLog.Println(err)
	}
	salt := util.GetRandomString()
	pwd := util.GetSha256Password(password, salt)
	result, err := stms.Exec(username, pwd, salt)
	stms.Close()
	if err != nil {
		util.ErrorLog.Println(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Admin{}, errors.New("用户创建失败")
	}
	return Admin{ID: id, Username: username}, nil
}
