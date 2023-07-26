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
	local data = ARGV[2]
	local length = ARGV[3]
	local llen = redis.call("llen", key)
	rlt[1] = llen
	if llen >= tonumber(length) then
		redis.call("rpop", key)
	end
	local push = redis.call(cmd, key, data)
	rlt[2] = push
	return rlt
`
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	params := []interface{}{"a","rpush","999",4}

	script := redis.NewScript(1, s)
	resp, err := redis.Ints(script.Do(c1, params...))
	log.Printf("111: %v \n", resp)
}
