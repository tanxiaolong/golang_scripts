package main

import (
	"fmt"
	"golang.org/x/net/context"
	"sync"
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

func main() {
	process := map[string]MultiProcessFunc{}
	//定时12点和20点给新用户推送
	pushNew := "12-20-push-new"
	process[pushNew] = func(ctx context.Context, arg interface{}) (interface{}, error) {
		// do sth
		fmt.Println(pushNew)
		return nil, nil
	}
	ctx := context.Background()
	MultiProcess(ctx, process, nil)
}
