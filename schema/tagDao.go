package schema


type TagDao struct{}

func (TagDao) Insert(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	id, err := ExecInsert(sql, params...)
	if err != nil {
		return DBErrorLog("添加失败", err)
	}
	tag := self.(*Tag)
	tag.ID = id
	return tag, nil
}

func (TagDao) Update(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	return UpdateReturnByGraphql(sql, self, params...)
}

func (TagDao) Delete(sql string, params ...interface{}) (int64, error) {
	return ExecDelete(sql, params...)
}

func (TagDao) QueryRow(sql string, params ...interface{}) (interface{}, error) {
	row, err := QueryRow(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var id int64
	var tag_name string
	err = row.Scan(&id, &tag_name)
	if err != nil {
		return DBErrorLog("标签不存在", err)
	}
	return Tag{id, tag_name}, nil
}

func (TagDao) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := Query(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	defer rows.Close()

	var result []Tag
	if rows.Next() {
		var id int64
		var tag_name string
		err = rows.Scan(&id, &tag_name)
		if err != nil {
			return DBErrorLog("标签不存在", err)
		}
		result = append(result, Tag{id, tag_name})
	}
	return result, nil
}

var tagDao DaoBasic

func init() {
	tagDao = TagDao{}
}
