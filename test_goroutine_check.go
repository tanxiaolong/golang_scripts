package main

import (
	"fmt"
	"regexp"
	"sync"
)

func main() {
	uids := []string{"12345", "01234"}
	uids = []string{"12345", "54321"}
	fmt.Println(checkUIDs3(uids))

}

func checkUIDs3(uids []string) bool {
	wg := sync.WaitGroup{}
	r := sync.Map{}
	for i := range uids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if !IsNum(id) {
				r.Store(id, false)
			}
		}(uids[i])
	}
	wg.Wait()

	count := 0
	f := func(k, v interface{}) bool {
		if k != "" {
			count++
		}
		return true
	}
	r.Range(f)
	if count > 0 {
		return false
	}
	return true
}

func checkUIDs2(uids []string) bool {
	wg := sync.WaitGroup{}
	result := true
	for i := range uids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			result = result && IsNum(id)
		}(uids[i])
	}
	wg.Wait()
	return result
}
func IsNum(idStr string) bool {
	pattern := `^[1-9][0-9]+$`
	isNum, _ := regexp.MatchString(pattern, idStr)
	return isNum
}
func checkUIDs(uids []string) bool {
	wg := sync.WaitGroup{}
	result := make(chan bool, 1)
	result <- true
	for i := range uids {
		wg.Add(1)
		go func(id string, rlt chan bool) {
			defer wg.Done()
			IsNumWithChan(id, rlt)
		}(uids[i], result)
	}
	wg.Wait()
	return <-result
}

func IsNumWithChan(idStr string, rlt chan bool) {
	//	for {
	select {
	case store := <-rlt:
		fmt.Println("store:", store)
		pattern := `^[1-9][0-9]+$`
		_, err := regexp.MatchString(pattern, idStr)
		fmt.Println(idStr, "err:", err)
		isnum := true
		if err != nil {
			isnum = false
		}
		fmt.Println(idStr, store, isnum)
		rlt <- (store && isnum)
		return
	}
	//	}
}
