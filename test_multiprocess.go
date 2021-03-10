package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
	"time"
)

//处理结果
type MultiProcResult struct {
	Error  error       //子函数error
	Result interface{} //子函数结果
}

//子处理函数定义
type MultiProcessFunc = func(ctx context.Context, arg interface{}) (interface{}, error)

//并行处理多个子函数
func MultiProcess(ctx context.Context, procm map[string]MultiProcessFunc, arg interface{}) map[string]*MultiProcResult {
	resultMap := sync.Map{}

	wg := sync.WaitGroup{}
	for key, process := range procm {
		wg.Add(1)
		go func(key string, process MultiProcessFunc) {
			ret, err := process(ctx, arg)
			result := &MultiProcResult{
				Result: ret,
				Error:  err,
			}
			resultMap.Store(key, result)
			wg.Done()
		}(key, process)
	}

	wg.Wait()

	multiResult := map[string]*MultiProcResult{}
	resultMap.Range(
		func(key, value interface{}) bool {
			multiResult[key.(string)] = value.(*MultiProcResult)
			return true
		},
	)

	return multiResult
}

func Seq() {

	time.Sleep(time.Second * 10)
	time.Sleep(time.Second * 10)
}

func main() {
	process := map[string]MultiProcessFunc{}
	//定时12点和20点给新用户推送
	process["12-20-push-new"] = func(ctx context.Context, arg interface{}) (interface{}, error) {
		// do sth
		time.Sleep(time.Second * 10)
		fmt.Println("done," + "12-20-push-new")
		return nil, nil
	}
	process["12-30-push-new"] = func(ctx context.Context, args interface{}) (interface{}, error) {
		time.Sleep(time.Second * 10)
		fmt.Println("done," + "12-30-push-new")
		return nil, nil
	}
	ctx := context.Background()
	startAt := time.Now()
	MultiProcess(ctx, process, nil)
	duration := time.Since(startAt)
	fmt.Println(duration)
	startAt = time.Now()
	Seq()
	duration = time.Since(startAt)
	fmt.Println(duration)
}
