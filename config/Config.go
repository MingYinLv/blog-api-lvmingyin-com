package config

import (
	"blog-api-lvmingyin-com/db"
)

type Config struct {
	Mysql      db.Mysql
	ListenPort int64
}

