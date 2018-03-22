package schema

type SpecialDao struct{}

func (SpecialDao) Insert(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	id, err := ExecInsert(sql, params...)
	if err != nil {
		return DBErrorLog("添加失败", err)
	}
	special := self.(*Special)
	special.ID = id
	return self, nil
}

func (SpecialDao) Update(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	return UpdateReturnByGraphql(sql, self, params...)
}

func (SpecialDao) Delete(sql string, params ...interface{}) (int64, error) {
	return ExecDelete(sql, params...)
}

func (SpecialDao) QueryRow(sql string, params ...interface{}) (interface{}, error) {
	row, err := QueryRow(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var id int64
	var name, cover string
	err = row.Scan(&id, &name, &cover)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	return Special{id, name, cover}, nil
}

func (SpecialDao) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := Query(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}

	var result []Special
	for rows.Next() {
		var id int64
		var name, cover string
		err = rows.Scan(&id, &name, &cover)
		if err != nil {
			return DBErrorLog("查询失败", err)
		}
		result = append(result, Special{id, name, cover})
	}
	return result, nil
}

var specialDao DaoBasic

func init() {
	specialDao = SpecialDao{}
}
