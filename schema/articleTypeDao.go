package schema

type ArticleTypeDao struct{}

func (ArticleTypeDao) Insert(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	id, err := ExecInsert(sql, params...)
	if err != nil {
		return DBErrorLog("添加失败", err)
	}
	actType := self.(*ArticleType)
	actType.ID = id
	return actType, nil
}

func (ArticleTypeDao) Update(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	return UpdateReturnByGraphql(sql, self, params...)
}

func (ArticleTypeDao) Delete(sql string, params ...interface{}) (int64, error) {
	return ExecDelete(sql, params...)
}

func (ArticleTypeDao) QueryRow(sql string, params ...interface{}) (interface{}, error) {
	row, err := QueryRow(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var id, show_menu int64
	var type_name, logo string
	err = row.Scan(&id, &type_name, &show_menu, &logo)
	if err != nil {
		return DBErrorLog("分类不存在", err)
	}
	return ArticleType{id, type_name, show_menu, logo}, nil
}

func (ArticleTypeDao) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := Query(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	defer rows.Close()

	var result []ArticleType
	if rows.Next() {
		var id, show_menu int64
		var type_name, logo string
		err = rows.Scan(&id, &type_name, &show_menu, &logo)
		if err != nil {
			return DBErrorLog("查询失败", err)
		}
		result = append(result, ArticleType{id, type_name, show_menu, logo})
	}
	return result, nil
}

var articleTypeDao DaoBasic

func init(){
	articleTypeDao = ArticleTypeDao{}
}
