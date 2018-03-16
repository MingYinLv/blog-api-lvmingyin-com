package util

import "fmt"

func GenInKeys(tableName string, keyName string, keys *[]interface{}, page, size int) string {
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	if len(*keys) > 0 {
		sql = fmt.Sprintf("%s WHERE %s in (", sql, keyName)
		for i := range *keys {
			sql = fmt.Sprintf("%s ?", sql)
			if i < len(*keys)-1 {
				sql = fmt.Sprintf("%s,", sql)
			}
		}
		sql = fmt.Sprintf("%s)", sql)
	}
	sql = fmt.Sprintf("%s limit %d,%d", sql, (page - 1) * size, size)
	fmt.Println(sql)
	return sql
}
