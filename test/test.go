package main

import (
  "fmt"
  "test.com/greetings"
  "rsc.io/quote"
)

func main(){
//  fmt.Println("hello,world!")
  fmt.Println(quote.Go())
  message := greetings.Hello("daifu")
  fmt.Println(message)
}
