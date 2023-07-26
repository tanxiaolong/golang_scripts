package main

import (
	_ "bufio"
	"errors"
	"fmt"
	"io"
	_ "io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 判断文件是否存在
// 存在返回true，否则false
func checkFileExist(fileName string) bool {
	exist := true
	fileInfo, err := os.Stat(fileName)
	fmt.Printf("%+v\n", fileInfo)
	if err != nil {
		exist = false
	}
	return exist
}

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	for tryTimes := 10; err != nil; tryTimes-- {
		file, err = os.Create(fileName)
	}
	if err != nil {
		return nil, errors.New("internal err")
	}
	return file, nil
}

func writeString(file io.Writer, content string) (bool, error) {
	count, err := io.WriteString(file, content)
	isSucc := true
	if count == 0 {
		isSucc = false
	}
	return isSucc, err
}

func main() {
	fileName := ""
	fmt.Println("请输入文件名：")
	fmt.Scanf("%s", &fileName)
	fmt.Printf("输入的是：%s\n", fileName)
	exist := checkFileExist(fileName)
	fmt.Printf("文件 %s 不存在，正在创建\n", fileName)
	if !exist {
		file, err := createFile(fileName)
		if file == nil {
			fmt.Println("文件 %s 创建失败，原因：%s", fileName, err.Error())
			return
		}
		fmt.Println("文件 %s 创建成功，要写东西吗？(y/n)")
		yesOrNo := ""
		fmt.Scanf("%s\n", &yesOrNo)
		fmt.Println("你要写什么")
		content := ""
		fmt.Scanf("%s\n", &content)
		if yesOrNo == "y" {
			isSucc, err := writeString(file, content)
			if isSucc {
				fmt.Println("请前去查看")
				return
			}
			fmt.Println("写失败，失败原因是：%s\n", err)
		}
	}

}
