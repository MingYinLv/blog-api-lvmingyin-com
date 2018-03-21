package schema

import (
	"blog-api-lvmingyin-com/util"
	"errors"
)

func DBErrorLog(text string, err error) (interface{}, error) {
	DBLog(err)
	return DBError(text)
}

func DBLog(err error) {
	util.ErrorLog.Println(err)
}

func DBError(text string) (interface{}, error) {
	return nil, errors.New(text)
}
