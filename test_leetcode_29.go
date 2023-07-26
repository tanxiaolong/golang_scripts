package main

import "log"

func main(){
    log.Print(float64(3*0.1/2))
    log.Print(divideLong(3,-2))
}

func divideLong(dividend int, divisor int) int {
    negative := false
    if (dividend < 0 && divisor > 0) || (dividend >0 && divisor < 0){
        negative = true
    }
    if dividend < 0 {
        dividend = -dividend
    }
    if divisor < 0 {
        divisor = -divisor
    }
    if dividend < divisor {
        return 0
    }
    sum := divisor
    divide := 1
    for sum+sum <= dividend {
        sum += sum
        divide += divide
    }
    if negative {
        return -(divide + divideLong((dividend-sum), divisor))
    } 
    return divide + divideLong((dividend-sum), divisor)
}
