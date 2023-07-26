package main

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

import "strconv"
import "fmt"
import "math"

func main(){
  fmt.Println(reverse2(1234567789))
}


func reverse2(x int) int {
    rev := 0
    for x != 0 {
        if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
            return 0
        }
        digit := x % 10
        x /= 10
        rev = rev*10 + digit
    }
    return rev
}

func reverse(x int32) int32 {
  xStr := fmt.Sprintf("%d", x)
  revStr := rev(xStr)
  aim, err := strconv.ParseInt(revStr, 10, 32)
  if err != nil {
    return 0
  }
  return int32(aim)
}

func rev(s string) string {
  b := []byte(s)
  l := len(b)
  mid := l/2
  for i:=0;i<mid;i++ {
    b[i],b[l-i-1] = b[l-i-1],b[i]
  }
  return string(b)
}
