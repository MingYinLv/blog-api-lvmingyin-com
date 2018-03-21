package schema

import (
	"blog-api-lvmingyin-com/db"
	"database/sql"
	"errors"
)

func ExecInsert(sql string, params ...interface{}) (int64, error) {

	stms, err := db.DB.Prepare(sql)
	if err != nil {
		DBLog(err)
		return 0, errors.New("添加失败")
	}
	defer stms.Close()

	result, err := stms.Exec(params...)
	if err != nil {
		DBLog(err)
		return 0, errors.New("添加失败")
	}

	return result.LastInsertId()
}

func ExecUpdate(sql string, params ...interface{}) (int64, error) {
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		DBLog(err)
		return 0, errors.New("修改失败")
	}
	defer stms.Close()

	result, err := stms.Exec(params...)
	if err != nil {
		DBLog(err)
		return 0, errors.New("修改失败")
	}

	return result.RowsAffected()
}

func ExecDelete(sql string, params ...interface{}) (int64, error) {
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		DBLog(err)
		return 0, errors.New("删除失败")
	}
	defer stms.Close()

	result, err := stms.Exec(params...)
	if err != nil {
		DBLog(err)
		return 0, errors.New("删除失败")
	}

	return result.RowsAffected()
}

func Query(sql string, params ...interface{}) (*sql.Rows, error) {
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		DBLog(err)
		return nil, errors.New("查询失败")
	}
	defer stms.Close()

	rows, err := stms.Query(params...)
	if err != nil {
		DBLog(err)
		return nil, errors.New("删除失败")
	}

	return rows, nil
}

func QueryRow(sql string, params ...interface{}) (*sql.Row, error) {
	stms, err := db.DB.Prepare(sql)
	if err != nil {
		DBLog(err)
		return nil, errors.New("查询失败")
	}
	defer stms.Close()

	result := stms.QueryRow(params...)
	return result, nil
}
