package main
import "fmt"
import "strconv"
func main(){
	buf := &[]byte{}
	a := 12345
	itoa(buf, a, 3)
	fmt.Printf("%s\n",string(*buf))
	fmt.Println(strconv.Itoa(a))
}

func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
