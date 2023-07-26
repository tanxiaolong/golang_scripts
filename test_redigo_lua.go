package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	s := `
local recentView = {}
local actorList = ARGV
-- #actorList get table's size
table.remove(actorList,1)
for i, k in pairs(actorList) do
	local v = redis.call("zscore",KEYS[1],k)
	recentView[i]=v
end
return recentView
`
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	params:= []interface{}{"aaa",1,2,3,4,5,6,7,8}
	//params:= []interface{}{"aaa","g","a"}

	script := redis.NewScript(1, s)
	resp, err := redis.Ints(script.Do(c1, params...))
	log.Println(len(resp),resp)
}
