package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	s := `
	local cmd = {}
	local key = KEYS[1]
	local key2 = KEYS[2]
	print(key,key2)
	local pid = ARGV[1]
	print(pid, type(pid))
	local srem = redis.call("srem", key, pid)
	cmd[1] = srem
	local del = redis.call("del",key2)
	cmd[2] = del
	-- hget result need to json unmarshal
	local hgetall = redis.call("hget", "txl","b")
	local lua_object = cjson.decode(hgetall)
	for k, v in pairs(lua_object) do
	    	print("k = " .. k)
		print("v = " .. v)
		print(v==tonumber(0))
	end
	print("-----------------------")
     	-- hgetall result native table, no need to json unmarshal
	local hkeys = redis.call("hkeys", "txl")
        for k, v in ipairs(hkeys) do
		local d = redis.call("hget","txl",v)
		print(type(d))
		local d_config = cjson.decode(d)
		print(type(d_config["a"]))
		if (v == "b") then
			d_config["a"] = tonumber(1)
		else
			d_config["a"] = tonumber(0)
		end
		local new_config = cjson.encode(d_config)
		print(type(new_config))
		local rlt = redis.call("hset","txl",v,new_config)
		print(rlt)
        end
	--if (tonumber("102") == 102)
	--then
	--  print("1")
	--  local rlt = redis.call("del", "txl")
	--  cmd[3] = rlt
	--else
	--  print("2")
	--  lua_object["c_prop_status"] =tonumber("102")
	--  local new_object = cjson.encode(lua_object)
	--  local rlt = redis.call("hset","txl","b", new_object)
	--  cmd[3] = rlt
	--end
	--if hgetall then
	--	local hgetall = cjson.decode(hgetall)
	--	print(hgetall["a"])
	--end
	return cmd
`
	c1, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer c1.Close()
	params := []interface{}{"living_pid","prepare_push:203706","203706"}

	script := redis.NewScript(2, s)
	resp, err := redis.Ints(script.Do(c1, params...))
	//resp, err := script.Do(c1, params...)
	//rr := resp.([]uint8)
	//log.Printf("%s\n", string(rr))
	log.Printf("111: %v \n", resp)
	//log.Printf("222: %v %v %v %v\n", len(resp), resp[0], resp[1], resp[2])
}
