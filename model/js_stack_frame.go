package model

import "danci-api/global"

// js错误堆栈
type JsErrorStackFrame struct {
	global.GVA_MODEL
	ColumnNumber int    `json:"column_number"`
	FileName     string `json:"file_name"`
	FunctionName string `json:"function_name"`
	LineName     string `json:"line_name"`
	Source       string `json:"source" gorm:"type:text"` // 这个就是通过source map 解析出来的源码

	PageJsErrorId uint
}
