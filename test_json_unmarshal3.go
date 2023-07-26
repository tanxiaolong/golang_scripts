package main

import "fmt"
import "encoding/json"

type DbTabInfos struct {
	Before []*DbTabInfo `protobuf:"bytes,1,rep,name=before,proto3" json:"before,omitempty"` // 直播前配置信息
	Ing    []*DbTabInfo `protobuf:"bytes,2,rep,name=ing,proto3" json:"ing,omitempty"`       // 直播中配置
	After  []*DbTabInfo `protobuf:"bytes,3,rep,name=after,proto3" json:"after,omitempty"`   // 直播后配置
}

type DbTabInfo struct {
	TabId     string     `protobuf:"bytes,1,opt,name=tabId,proto3" json:"tabId,omitempty"`                                                               // 页卡id
	TabType   string     `protobuf:"bytes,2,opt,name=tabType,proto3" json:"tabType,omitempty"`                                                           // 页卡类型，前中后页卡类型相同
	Title     string     `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                                               // 页卡标题
	Order     string     `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`                                                               // 页卡排序
	H5Url     string     `protobuf:"bytes,5,opt,name=H5Url,proto3" json:"H5Url,omitempty"`                                                               // h5 url
	Client    ClientType `protobuf:"varint,6,opt,name=Client,proto3,enum=trpc.video_app_live.live_console_ving_read.ClientType" json:"Client,omitempty"` // 端类型
	Roomid    string     `protobuf:"bytes,7,opt,name=roomid,proto3" json:"roomid,omitempty"`                                                             // 轮播台ID
	ExtraInfo string     `protobuf:"bytes,8,opt,name=extraInfo,proto3" json:"extraInfo,omitempty"`                                                       // 扩展信息，为h5时可填入跳转链接，图片则上传图片
	Selected  int32      `protobuf:"varint,9,opt,name=selected,proto3" json:"selected,omitempty"`                                                        // 页卡被默认选中状态 0|未选中、1|选中
}

type ClientType int32

func main() {
	str := `{"before":[{"tabId":"65","tabType":"25","title":"提问区","order":"0","Client":2},{"tabId":"33","tabType":"12","title":"预约","order":"0","Client":1},{"tabId":"34","tabType":"21","title":"精彩回顾2222","order":"1","Client":1},{"tabId":"71","tabType":"27","title":"相关视频","order":"2","Client":1,"extraInfo":"[{"part_name":"vid","part_type":2,"resource":[{"id":"a33080lx4tx"},{"id":"a33080lx4tx"}]},{"part_name":"cid","part_type":1,"resource":[{"id":"mzc00200tsv8yfb"}]}]"}],"ing":[{"tabId":"0","tabType":"1","title":"主持人","order":"0","Client":2},{"tabId":"1","tabType":"2","title":"聊天室","order":"0","Client":1},{"tabId":"3","tabType":"5","title":"榜单","order":"1","Client":1}],"after":[{"tabId":"43","tabType":"1","title":"主持人是谁","order":"0","Client":2},{"tabId":"54","tabType":"3","title":"发电666","order":"0","H5Url":"www.tencent.com","Client":1}]}
`
	dst := &DbTabInfos{}
	err := json.Unmarshal([]byte(str), dst)
	fmt.Println(dst)
	fmt.Println(err)

}
