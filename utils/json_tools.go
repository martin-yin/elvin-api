package utils

import (
	"encoding/json"
	"fmt"
)

func JSONToStruct(str string, v interface{}) error {
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		return err
	}
	return nil
}

func StructToJSON(v interface{}) (str string, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func StructToJsonToStruct(v interface{}, o interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	str := string(data)
	err = json.Unmarshal([]byte(str), &o)
	if err != nil {
		return err
	}
	return nil
}

// interface 转 json 在转换成 struct
func InterfaceToJsonToStruct(v interface{}, o interface{}) error {
	resByre, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(resByre), &o)
	if err != nil {
		return err
	}
	fmt.Println(o, "oooooooooooooooooooo")
	return nil
}
