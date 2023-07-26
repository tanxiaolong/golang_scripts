package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	s := `
	local rlt = {}
	local key = KEYS[1]
	local cmd = ARGV[1]
	local tmp = ARGV[2]
	local status = redis.call("hget", key, cmd)
	rlt[1] = status
	print(type(status))
	print(type(tmp))
	if status == tonumber(tmp) then
		print("hello")
	end
	return rlt
`
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	params := []interface{}{"txl","a","1"}

	script := redis.NewScript(1, s)
	resp, err := redis.Ints(script.Do(c1, params...))
	log.Printf("111: %v \n", resp)
}
