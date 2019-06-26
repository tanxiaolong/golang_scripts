package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	filename := "/Users/tanxiaolong/t"
	content, err := readAllIntoMemory(filename)
	fmt.Println("err", err)
	m := []map[string]interface{}{}
	json.Unmarshal(content, &m)
	arr, mm := MapColumn(m, "ad_id", "id")
	fmt.Println("arr", arr)
	fmt.Println("mm", mm)
	arr, mm = MapColumn(m, "name", "id")
	fmt.Println("arr", arr)
	fmt.Println("mm", mm)
}

func MapColumn(m interface{}, want_col string, key_col string) ([]interface{}, map[interface{}]interface{}) {
	mBytes, _ := json.Marshal(m)
	var mMap []map[string]interface{}
	json.Unmarshal(mBytes, &mMap)
	var rlt []interface{}
	tmp := map[interface{}]interface{}{}
	for _, v := range mMap {
		val := v[want_col]
		fmt.Println("want_col", want_col, "val", val)
		//if keyVal,exists := v[key_col];exists {
		//}
		if _, exists := tmp[v[key_col]]; !exists {
			//fmt.Println("not exists", val)
			rlt = append(rlt, val)
			tmp[v[key_col]] = val
			//} else {
			//	fmt.Println("exists", val)
		}
	}
	return rlt, tmp
}

func readAllIntoMemory(filename string) (content []byte, err error) {
	fp, err := os.Open(filename) // 获取文件指针
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer) // 文件内容读取到buffer中
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
