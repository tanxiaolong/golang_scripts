package main

import (
	"fmt"
	"github.com/huichen/sego"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("/Users/tanxiaolong/workbench/go/src/github.com/huichen/sego/data/dictionary.txt")

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	var segments []sego.Segment
	for i := 0; i < 1694115; i++ {
		text = []byte("前几天我的一个大学同学负责的网站出现了严重的性能瓶颈")
		segments = segmenter.Segment(text)
	}
	fmt.Println(segments)
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	//fmt.Println(sego.SegmentsToString(segments, false))
}
