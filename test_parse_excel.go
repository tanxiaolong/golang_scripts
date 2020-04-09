package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	path := `/Users/tanxiaolong/Downloads/20200303102650302_null_185736_exec.xlsx`
	parseExcel(path)
}

func parseExcel(path string) {
	xlFile, err := xlsx.OpenFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//遍历行读取
		for i, row := range sheet.Rows {
			fmt.Printf("line num:", i)
			// 遍历每行的列读取
			postData := map[string]interface{}{}
			data := []interface{}{}
			for _, cell := range row.Cells {
				text := cell.String()
				data = append(data, text)
				fmt.Printf("%20s", text)
			}
			fmt.Println()
			postData["uid"], _ = strconv.Atoi(data[0].(string))
			postData["duration"], _ = strconv.Atoi(data[1].(string))
			postData["date_day"] = data[2]
			bytes, _ := json.Marshal(postData)

			fmt.Println("post data:", string(bytes))
			curl(string(bytes))
			//break
		}
	}
	fmt.Println("\n\nimport success")
}

func curl(p string) {
	url := `http://10.104.4.19:10284/api/v1/mediacenter/data/mic_duration_report`
	method := "POST"

	payload := strings.NewReader(p)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("uberctx-_namespace_appkey_", "haokan")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println("response from server: ", string(body))
}
