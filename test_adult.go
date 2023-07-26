package main
import "time"
import "fmt"
func main(){
id := "440106200505314023"
ymd := id[6:14]
fmt.Println(ymd)
ymdTime, _ := time.Parse("20060102", ymd)
fmt.Println(ymdTime)
	d := time.Since(ymdTime)
	y := d.Hours() / 8760
fmt.Println(y)
}
