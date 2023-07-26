package main
import "fmt"
func main(){
c1:=make(chan int)        //无缓冲
c2:=make(chan int,1)      //有缓冲
close(c1)
rlt := <-c1
fmt.Printf("1:%+v\n", rlt)
rlt = <-c1
fmt.Printf("2:%+v\n", rlt)
rlt = <-c1
fmt.Printf("3:%+v\n", rlt)



c2<-1
close(c2)
rltC2, ok :=<- c2
fmt.Printf("c2 %+v, ok: %v\n",rltC2,ok)
rltC2, ok =<- c2
fmt.Printf("c2 %+v, ok: %v\n",rltC2,ok)

c3 := make(chan int,1)
c3<-1
<-c3
fmt.Println("hello c3")

//c4 := make(chan int)
//c4<-1
//<-c4
/*
// 重复关闭导致的panic可以被recover到
defer func(){
	r := recover()
	fmt.Println(r)
}()
close(c2)
*/

c2 = nil
rltC2, ok =<- c2
fmt.Println("nil c2", rltC2,ok)


}
