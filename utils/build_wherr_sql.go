package utils

import (
	"reflect"
)

// 生成where 查询条件
func BuildWhereSql(model, sql string, params ...interface{}) (baseSql string, sqlParams []interface{}) {
	baseSql = "from_unixtime(" + model + ".happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and " + model + ".monitor_id = ? "
	if sql != "" {
		baseSql += sql
	}
	for _, value := range params {
		param := reflect.ValueOf(value)
		for i := 0; i < param.NumField(); i++ {
			sqlParams = append(sqlParams, param.Field(i).Interface())
		}
	}
	return
}
