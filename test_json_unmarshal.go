package main

import "fmt"
import "encoding/json"

// CommentInfo 评论详情结构体
type CommentInfoRsp struct {
	Code int64 `json:"code"`
	Data struct {
		// key是评论ID
		Comments map[string]CommentItem `json:"comments"`
	} `json:"data"`
}

// CommentItem 评论详情具体字段
type CommentItem struct {
	CheckHotScale string `json:"checkhotscale"`
	CheckStatus   string `json:"checkstatus"`
	CheckType     string `json:"checktype"`
	Content       string `json:"content"`
	Custom        string `json:"custom"`
	EmotionalTag  string `json:"emotionalTag"`
	Extend        struct {
		At string `json:"at"`
		Ut string `json:"ut"`
		Wt string `json:"wt"`
	} `json:"extend"`
	Highlight   string `json:"highlight"`
	ID          string `json:"id"`
	IsAuthor    string `json:"isauthor"`
	IsCity      string `json:"iscity"`
	IsDeleted   string `json:"isdeleted"`
	IsDown      string `json:"isdown"`
	IsFans      string `json:"isfans"`
	IsGod       string `json:"isgod"`
	IsHide      string `json:"ishide"`
	IsPick      string `json:"ispick"`
	OriReplyNum string `json:"orireplynum"`
	Parent      string `json:"parent"`
	PokeNum     string `json:"pokenum"`
	PUserID     string `json:"puserid"`
	RepNum      string `json:"repnum"`
	RichType    string `json:"richtype"`
	RootID      string `json:"rootid"`
	Source      string `json:"source"`
	TargetID    string `json:"targetid"`
	Time        string `json:"time"`
	Up          string `json:"up"`
	UserID      string `json:"userid"`
	VoteID      string `json:"voteid"`
}

func main() {

	data := `{
    "code": 0,
    "data": {
        "comments": { 
            "6805046640236494972": {
                "checkhotscale": "0", 
                "checkstatus": "0",
                "checktype": "1", 
                "content": "\u56de\u590d\u4e00\u5929\u4e86",
                "custom": "{\"uid\":\"865770910\"}",
                "emotionalTag": "3",
                "extend": {
                    "at": "0",
                    "ut": "0",
                    "wt": "0"
                },
                "highlight": "0",
                "id": "6805046640236494972", 
                "isauthor": "0",
                "iscity": "0",
                "isdeleted": "0",
                "isdown": "0",
                "isfans": "0",
                "isgod": "0",
                "ishide": "0",
                "ispick": "0",
                "orireplynum": "0",
                "parent": "6804716035868190772", 
                "pokenum": "0",
                "puserid": "1023231511", 
                "repnum": "0",
                "richtype": "0",
                "rootid": "6804716035868190772",
                "source": "162",
                "targetid": "6921368624", 
                "time": "1622449550",
                "up": "0",
                "userid": "584442166", 
                "voteid": "0"
            }
        }
    }
}`
	a := &CommentInfoRsp{}
	err := json.Unmarshal([]byte(data), a)
	fmt.Println(err)
	fmt.Println(a.Data.Comments["6805046640236494972"])

	fmt.Println("------------------")
	ttt := `{"flopcmtuser":"1310716683","flopcmtid":"6837317350949862287","name_alias":"","flopinfo":{"username":"iantan","vuid":"12345"}}`
	rlt := map[string]string{}
	json.Unmarshal([]byte(ttt), &rlt)
	fmt.Println(rlt, rlt["flopcmtid"])
}
