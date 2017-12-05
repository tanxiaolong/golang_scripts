package main
import "fmt"
import "time"
import "bytes"
import "io"
import "os"

func w(fileName string, content string){
	file,_ := os.Create(fileName)
	io.WriteString(file,content)
}

func main(){
	str := "hello world"
	var buf bytes.Buffer
	n := 1
	t0 := time.Now()
	for i:=0;i<n;i++{
		buf.WriteString(str)
	}
	_=buf.String()
	//fmt.Println(buf.String())
	t01 := time.Now()
	fmt.Println(t01.Sub(t0))
	//w("test_1buf",buf.String())

	a := ""
	t1 := time.Now()
	for i:=0;i<n;i++{
		a += str
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	//w("test_2plus",a)


	t3 := time.Now()
	b := []byte{}
	bstr := []byte(str)
	for i:=0;i<n;i++{
		b = append(b,bstr...)
	}
	_=string(b)
	t4 := time.Now()
	fmt.Println(t4.Sub(t3))
	//w("test_3byte_arr",string(b))

	t5 := time.Now()
	k := str
	for i:=0;i<n;i++{
		k = fmt.Sprintf("%s%s",k,str)
	}
	t6 := time.Now()
	fmt.Println(t6.Sub(t5))
	//w("test_4sprintf",k)
}
