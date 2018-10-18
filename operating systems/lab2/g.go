
package main

import "fmt"
import "time"
import "runtime"
import "math/rand"


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

var (locked = 1
 
)
func main() {
    var coins = []coin{
        coin{1,60},
        coin{2,50},
        coin{5,0},
        coin{10,20},
        coin{25,110},
        coin{50,42},
        coin{100,0}}
        
    numcpu := runtime.NumCPU()
    fmt.Println("NumCPU", numcpu)
    //runtime.GOMAXPROCS(numcpu)
    runtime.GOMAXPROCS(1)
	ch1 := make(chan int)
	ch2 := make(chan float64)

	go func() { 
         
        for {
        local1 := 1
		for local1 ==1 {
            local1=TestAndSet()
            fmt.Println("start: ",coins)
            ch1 <- RandCoin(len(mon))
			locked = 0
			ch1 <- -1
		}
        ch2 <- 0.0
		
	}}()
	go func() {
       
		for {
   			local2 := 1
			for local2 == 1 {
                 local2 = TestAndSet()
                beginCoin := <-ch1
                 
                 
                 
                if IndexOf(beginCoin) != -1{
               
               fmt.Println("Begin coin: ",beginCoin)
                changCoin := 0
                
                if IndexOf(beginCoin) != 0{
                    coins[IndexOf(beginCoin)].amount += 1
                    fmt.Print("select a number coin to be changed from ",)
                    printPart(IndexOf(beginCoin))
                    changCoin = RandCoin(IndexOf(beginCoin))
                     fmt.Scanf("%d",&changCoin)
                    if IndexOf(changCoin) == -1 || IndexOf(changCoin) >=IndexOf(beginCoin){
                        fmt.Println("\nexchanger can not change this coin ",changCoin)
                        if beginCoin == -1{
                            ch2 <- 1.1
//                             fmt.Println("end: ",coins)
                        }
                        
                        locked = 0
                        break
                    }
                }else{
                    fmt.Println("Sorry, exchanger can not change 1 kopiyka")
                     if beginCoin == -1{
                        ch2 <- 1.1
//                         fmt.Println("end: ",coins)
                     }
                        
                        locked = 0
                        break
                     
                    
                }
                    if beginCoin%changCoin != 0 {
                        if coins[IndexOf(beginCoin%changCoin)].amount >= 1  {
                            errChange(changCoin,beginCoin,coins)
                            fmt.Println(" + 1 *",beginCoin%changCoin)
                            coins[IndexOf(beginCoin%changCoin)].amount -= 1
                        }else if beginCoin%changCoin == 5 && coins[IndexOf(1)].amount>=5 {
                            errChange(changCoin,beginCoin,coins)
                            fmt.Println(" + 1 * 5")
                            coins[IndexOf(1)].amount -= 5
                        }else{
                            coins[IndexOf(beginCoin)].amount -= 1
                            fmt.Println("\nexchanger do not count the amount of coins")
                             if beginCoin == -1{
                                ch2 <- 1.1
//                                 fmt.Println("end: ",coins)
                             }
                            
                            locked = 0
                            break
                        }
                    }else{
                        errChange(changCoin,beginCoin,coins)
                    }
                    fmt.Println()
                }                
                                
                if beginCoin == -1{
                    ch2 <- 1.1
//                     fmt.Println("end: ",coins)
                }
                
				locked = 0
			}
		}
	}()
    for{
        <-ch2
        fmt.Println("\nend\n\n: ",coins)
    }
}
