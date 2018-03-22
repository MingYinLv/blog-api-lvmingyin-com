package schema

type ArticleDao struct{}

func (ArticleDao) Insert(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	id, err := ExecInsert(sql, params...)
	if err != nil {
		return DBErrorLog("添加失败", err)
	}
	article := self.(*Article)
	article.ID = id
	return article, nil
}

func (ArticleDao) Update(sql string, self interface{}, params ...interface{}) (interface{}, error) {
	return UpdateReturnByGraphql(sql, self, params...)
}

func (ArticleDao) Delete(sql string, params ...interface{}) (int64, error) {
	return ExecDelete(sql, params...)
}

func (ArticleDao) QueryRow(sql string, params ...interface{}) (interface{}, error) {
	row, err := QueryRow(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var id, type_id, create_at, update_at int64
	var title, content, cover string
	err = row.Scan(&id, &title, &content, &cover, &type_id, &create_at, &update_at)
	if err != nil {
		return DBErrorLog("文章不存在", err)
	}
	return Article{id, title, content, cover, type_id, create_at, update_at}, nil
}

func (ArticleDao) Query(sql string, params ...interface{}) (interface{}, error) {
	rows, err := Query(sql, params...)
	if err != nil {
		return DBErrorLog("查询失败", err)
	}
	var result []Article
	for rows.Next() {
		var id, type_id, create_at, update_at int64
		var title, content, cover string
		err = rows.Scan(&id, &title, &content, &cover, &type_id, &create_at, &update_at)
		if err != nil {
			return DBErrorLog("查询失败", err)
		}
		result = append(result, Article{id, title, content, cover, type_id, create_at, update_at})
	}
	return result, nil
}

var articleDao DaoBasic

func init(){
	articleDao = ArticleDao{}
}
