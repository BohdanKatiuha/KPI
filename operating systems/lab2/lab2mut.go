
package main

import (
    "fmt"
    "time"
    "runtime"
    "math/rand"
    "sync"
)


type coin struct{
        value int
        amount int 
}

var mon = []int{1,2,5,10,25,50,100}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandCoin(l int) int {
    return mon[rand.Intn(l)]
}

func IndexOf(val int)int{
    
    for i:=0;i<len(mon);i++{
        if val == mon[i]{
            return i
        }
    }
    return -1
}

func printPart(n int){
    for i:=0;i< n;i++{
        fmt.Print(mon[i]," ")
    }
    fmt.Print(":")
}

func errChange(changCoin, beginCoin int,coins []coin){
     if coins[IndexOf(changCoin)].amount-beginCoin/changCoin<0{
         
        coins[IndexOf(beginCoin)].amount -= 1
        fmt.Println("exchanger do not count the amount of coins")
     }else{
        fmt.Print(changCoin," * ",beginCoin/changCoin)
        coins[IndexOf(changCoin)].amount -= beginCoin/changCoin
    }
}

func TestAndSet () int { 
    initial := locked
    locked = 1
    return initial
}

var(
    locked = 1
    identLock sync.Mutex
    exchangLock sync.Mutex
)

var locked = 0
func main() {
   
}




// package main
/* producer-consumer problem in Go */

/*
import ("fmt")

var done = make(chan bool)
var msgs = make(chan int)

func produce () {
    for i := 0; i < 10; i++ {
        msgs <- i
    }
    done <- true
}

func consume () {
    for {
      msg := <-msgs
      fmt.Println(msg)
   }
}

func main () {
   go produce()
   go consume()
   <- done
}*/
