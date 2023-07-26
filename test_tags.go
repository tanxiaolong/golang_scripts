package main

import (
	"encoding/json"
	"fmt"
)

type tag struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	AppID       int    `json:"appid"`
	OrderIdx    int    `json:"order_idx"`
	ExtraID     string `json:"extra_id"`
	ExtraData   string `json:"extra_data"`
	ParentID    uint32 `json:"parent_id"`
	IsShowOrder int    `json:""`
}

func main() {
	tagsIlive := `[{"id":40,"name":"娱乐" , "level":1 , "appid":1017 , "order_idx":4,  "extra_id":"test"  ,"extra_data":"ttt"} ,{"id":42,"name":"生活"  ,"level":1 , "appid":1017 , "order_idx":4 },{"id":43,"name":"游戏"  ,"level":1  ,"appid":1017,  "order_idx":1 },{"id":44,"name":"其他"  ,"level":1  ,"appid":1017  ,"order_idx":4 },{"id":45,"name":"音乐"  ,"level":2  ,"parent_id":40 ,"appid":1017  ,"order_idx":3},{"id":46,"name":"舞蹈"  ,"level":2  ,"parent_id":40 ,"appid":1017  ,"order_idx":3},{"id":47,"name":"颜值"  ,"level":2  ,"parent_id":40 ,"appid":1017  ,"order_idx":3 },{"id":48,"name":"脱口秀4","level":2  ,"parent_id":40  ,"appid":1017  ,"order_idx":2 },{"id":49,"name":"二次元","level":2  ,"parent_id":40  ,"appid":1017  ,"order_idx":1 },{"id":53,"name":"美食" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":54,"name":"萌宠" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":55,"name":"情感" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":56,"name":"时尚" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":57,"name":"手艺" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":58,"name":"运动" , "level":2  ,"parent_id":42  ,"appid":1017 },{"id":59,"name":"知识2", "level":2  ,"parent_id":44  ,"appid":1017 },{"id":60,"name":"英雄联盟","level":2  ,"parent_id":43  ,"appid":1017  ,"order_idx":3 },{"id":61,"name":"绝地求生","level":2  ,"parent_id":43  ,"appid":1017  ,"order_idx":3 },{"id":62,"name":"王者荣耀","level":2  ,"parent_id":43  ,"appid":1017  ,"order_idx":3 },{"id":63,"name":"和平精英","level":2  ,"parent_id":43  ,"appid":1017 },{"id":64,"name":"暴雪游戏","level":2  ,"parent_id":43  ,"appid":1017 },{"id":65,"name":"单机热游","level":2  ,"parent_id":43  ,"appid":1017 },{"id":66,"name":"综合手游","level":2  ,"parent_id":43  ,"appid":1017 },{"id":67,"name":"综合网游","level":2  ,"parent_id":43  ,"appid":1017 },{"id":68,"name":"其他游戏","level":2  ,"parent_id":43  ,"appid":1017 },{"id":72,"name":"新增一级分类","level":1  ,"parent_id":1  ,"appid":1017 },{"id":73,"name":"新增二级分类","level":2  ,"parent_id":72 ,"appid":1017 },{"id":74,"name":"一级分类" , "level":1 ,"appid":1017 },{"id":75,"name":"二级分类1" ,"level":2  ,"parent_id":74  ,"appid":1017 },{"id":78,"name":"三级测试标签"  ,"level":3  ,"parent_id":45  ,"appid":1017 },{"id":79,"name":"脱口秀三级标签"  ,"level":3 ,"parent_id":48  ,"appid":1017 },{"id":80,"name":"脱口秀三级标签2" ,"level":3 ,"parent_id":48  ,"appid":1017 },{"id":81,"name":"脱口秀三级标签3" ,"level":3 ,"parent_id":48  ,"appid":1017 },{"id":82,"name":"脱口秀三级标签4" ,"level":3,"parent_id":48  ,"appid":1017 },{"id":83,"name":"可见标签1"  ,"level":1  ,"appid":1017 },{"id":84,"name":"可见标签2"  ,"level":83 ,"appid":1017 },{"id":85,"name":"可见标签4"  ,"level":4  ,"parent_id":60  ,"appid":1017 , "order_idx":5}]`
	var tags []*tag
	json.Unmarshal([]byte(tagsIlive), &tags)
	fmt.Println(tags)

	validParent := consValidTags(tags)
	fmt.Println("valid tags", validParent)

	var filteredTagsInfo []*tag
	for i := range tags {
		tag := tags[i]
		if isTagValid(tag, validParent) {
			filteredTagsInfo = append(filteredTagsInfo, tag)
		}
	}

	for i := range filteredTagsInfo {
		fmt.Println("filteredTagsInfo", filteredTagsInfo[i])
	}

	a := map[int][]*tag{}
	fmt.Println("a", a[0])
}

func consValidTags(tags []*tag) map[uint32]uint32 {
	// map[uint32]uint32 维护了tag.ID到tag.ParentID的关系
	validTags := map[uint32]uint32{0: 0}
	for i := range tags {
		tag := tags[i]
		validTags[tag.ID] = tag.ParentID
	}
	return validTags
}

func isTagValid(tag *tag, validParent map[uint32]uint32) bool {
	i, parent := 0, tag.ParentID
	for true {
		grandParent, exists := validParent[parent]
		if !exists {
			break
		}
		if grandParent == 0 {
			return true
		}
		if grandParent == parent {
			fmt.Printf("tag's parent recur, tag_id: %d, parent: %d, grandparent: %d", tag.ID, parent, grandParent)
			break
		}
		parent = grandParent
		// 防止死循环
		if i == 5 {
			break
		}
		i++
	}
	return false
}
