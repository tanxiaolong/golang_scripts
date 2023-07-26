package main

import "fmt"
func main(){
	call()
	fmt.Println("333 Helloworld")
}

func call()  {
	
	defer func(){
		fmt.Println("11111")
	}()
	defer func() {
		if r := recover(); r!= nil {
			fmt.Println("Recover from r : ",r)
		}
	}()
	defer func(){
		defer func(){
			fmt.Println("defer in 22222")
		}()
		defer func(){
			if r := recover(); r != nil {
				fmt.Println("recover in 22222")
			}
		}()
		fmt.Println("22222")
		panic("panic in 2222");
	}()

	defer func(){
		fmt.Println("33333")
	}()

	fmt.Println("111 Helloworld")
	panic("Panic 1!")
    // 这个panic和后面的代码已经不会执行，ide应该会有提示
	panic("Panic 2!")
	fmt.Println("222 Helloworld")
}
