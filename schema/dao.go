package schema

type DaoBasic interface {
	Insert(sql string, self interface{}, params ...interface{}) (interface{}, error)
	Update(sql string, self interface{}, params ...interface{}) (interface{}, error)
	Delete(sql string, params ...interface{}) (int64, error)
	QueryRow(sql string, params ...interface{}) (interface{}, error)
	Query(sql string, params ...interface{}) (interface{}, error)
}

func UpdateReturnByGraphql(sql string, self interface{}, params ...interface{}) (interface{}, error){
	_, err := ExecUpdate(sql, params...)
	if err != nil {
		return DBErrorLog("修改失败", err)
	}
	return self, nil
}
