package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	s := `
local cmd_rlt = {}
local key = KEYS[1]
local ts,room_id = ARGV[1],ARGV[2]
local r_zadd = redis.call("zadd",key,ts,room_id)
cmd_rlt[1] = r_zadd
local max_cnt = ARGV[3]
local r_zremrangebyrank = redis.call("zremrangebyrank",key,0,max_cnt)
cmd_rlt[2] = r_zremrangebyrank
local max_date = ARGV[4]
local r_zremrangebyscore = redis.call("zremrangebyscore",key,0,max_date)
cmd_rlt[3] = r_zremrangebyscore
local expire_date = ARGV[5]
local r_expire = redis.call("expire",key,expire_date)
cmd_rlt[4] = r_expire
return cmd_rlt
`
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	params := []interface{}{"iiii", 124, 456, -500, 10, 60}

	script := redis.NewScript(1, s)
	resp, err := redis.Ints(script.Do(c1, params...))
	//resp, err := script.Do(c1, params...)
	//rr := resp.([]uint8)
	//log.Printf("%s\n", string(rr))
	log.Printf("111:,%v\n", resp)
	log.Printf("222:%v %v %v %v %v \n", len(resp), resp[0], resp[1], resp[2], resp[3])
}
