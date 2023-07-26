package main

import "log"

func main(){
	log.Print(diStringMatch("DDD"))
}

func diStringMatch(S string) []int {
    i := []int{0}
    d := []int{}
    k := 1
    l := len(S)
    for _,c := range S {
        if string(c) == "I" {
            i = append(i,k)
	    k++
        }else if string(c) == "D" {
            d = append(d,l)
	    l--
        }
    }
    li := len(i)
    ld := len(d)
    if ld == 0 {
	return i
    }
    m,n := 0,0
    rlt := []int{}
    for _,c := range S {
	if string(c) == "I" {
	    rlt = append(rlt,i[m])
	    m++
	}else if string(c) == "D" {
	    rlt = append(rlt,d[n])
	    n++
	}
    }
    if m<li {
	rlt = append(rlt,i[m:]...)
    }
    if n<ld {
	rlt = append(rlt,d[n:]...)
    }
    return rlt
}
