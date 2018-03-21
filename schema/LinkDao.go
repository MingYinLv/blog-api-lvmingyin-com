package schema

type LinkDao struct{}

func (*LinkDao) QueryRow(sql string, params ...interface{}) (interface{}, error) {
	row, err := QueryRow(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var id, linkType int64
	var icon, url, name string
	err = row.Scan(&id, &icon, &linkType, &url, &name)
	if err != nil {
		return DBErrorLog("链接不存在", err)
	}

	if err != nil {
		return DBErrorLog("链接不存在", err)
	}
	return Link{id, icon, linkType, url, name}, nil
}

func (*LinkDao) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := Query(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	defer rows.Close()

	var result []Link
	if rows.Next() {
		var id, linkType int64
		var icon, url, name string
		err = rows.Scan(&id, &icon, &linkType, &url, &name)
		if err != nil {
			return DBErrorLog("查询失败", err)
		}
		result = append(result, Link{id, icon, linkType, url, name})
	}
	return result, nil
}

func (*LinkDao) Delete(sql string, params ...interface{}) (int64, error) {
	return ExecDelete(sql, params...)
}

func (*LinkDao) Update(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	_, err := ExecUpdate(sql, params...)

	if err != nil {
		return DBErrorLog("修改失败", err)
	}
	return self, nil
}

func (*LinkDao) Insert(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	id, err := ExecInsert(sql, params...)
	if err != nil {
		return DBErrorLog("添加失败", err)
	}
	link := self.(Link)
	link.ID = id
	return link, nil
}

var linkDao = LinkDao{}
