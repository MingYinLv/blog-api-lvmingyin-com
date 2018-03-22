package schema

import (
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
			Type: graphql.NewNonNull(graphql.String),
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

func FindAdminByUserName(username string) (interface{}, error) {
	row, err := QueryRow("SELECT * FROM admin WHERE username = ?", username)
	if err != nil {
		return DBErrorLog(fmt.Sprintf("用户 %s 不存在", username), err)
	}

	var id int64
	var userResult, pwdResult, salt string
	err = row.Scan(&id, &userResult, &pwdResult, &salt)
	if err != nil {
		return DBErrorLog(fmt.Sprintf("用户 %s 不存在", username), err)
	}
	return Admin{id, userResult, pwdResult, salt}, nil
}

func AdminLogin(username, password string) (interface{}, error) {
	a, err := FindAdminByUserName(username)
	if err != nil {
		// 其他错误
		return DBError("登录失败")
	}

	admin := a.(Admin)

	if admin.Password == util.GetSha256Password(password, admin.Salt) {
		// 密码加密有一致
		return Admin{ID: admin.ID, Username: admin.Username}, nil
	} else {
		// 密码错误
		return DBError("密码错误")
	}

}

func AdminCreate(username, password string) (interface{}, error) {

	admin, err := FindAdminByUserName(username)
	if err == nil && admin.(Admin).ID > 0 {
		return DBError("该用户名已注册")
	}

	salt := util.GetRandomString()
	pwd := util.GetSha256Password(password, salt)

	id, err := ExecInsert("INSERT INTO admin(username,password,salt) values(?,?,?)", username, pwd, salt)

	if err != nil {
		return DBErrorLog("用户创建失败", err)
	}
	return Admin{ID: id, Username: username}, nil
}
