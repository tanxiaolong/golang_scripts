package main

import "log"

func main(){
  defer func(){
    if err := recover();err!=nil{
      log.Println("panic err:",err)
    }
  }()

  panic("123")
}
