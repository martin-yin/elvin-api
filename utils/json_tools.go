package utils

import "encoding/json"

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
