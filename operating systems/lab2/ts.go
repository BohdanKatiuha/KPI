package main
import(
    "fmt"
     "math/rand"
//     "os"
    "time"
//     "runtime"
//     "strconv"
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
         
        coins[IndexOf(beginCoin)].amount += 1
        fmt.Println("exchanger do not count the amount of coins")
//         os.Exit(1)
       
     }
}



func TestAndSet (lock bool) bool { 
    initial := lock
    locked = true
    return initial
} 

// Reason: Nice implementation of Lock/Unlock: 






var locked bool
func main(){
    

var indB int
    var coins = []coin{
        coin{1,6},
        coin{2,50},
        coin{5,0},
        coin{10,20},
        coin{25,10},
        coin{50,4},
        coin{100,0}}
    locked = true
    
     beginCoin := make(chan int)
     
     
    go func(coins []coin){ 
        for true{
        for TestAndSet (locked) == true { 
            fmt.Println(coins)
            beginCoin <- RandCoin(len(mon))
            bc := int(<-beginCoin)
            indB = IndexOf(bc)
            
            fmt.Println("Begin coin: ",bc)
            locked = false
        
        }
        
        }
    }(coins)
   
   go func (coins []coin){
    for true{
        for TestAndSet (locked) == true {
            coins[indB].amount += 1
            exit := -1
            changCoin := 0
            bc := int(<-beginCoin)
            if IndexOf(bc) != 0{
                for true{
                    fmt.Print("select a number coin to be changed from ",)
                    printPart(IndexOf(bc))
                    fmt.Println(bc)
                    changCoin = RandCoin(IndexOf(bc))
                    fmt.Println("changCoin=",changCoin)
//                     fmt.Scanf("%d",&changCoin)
                    if IndexOf(changCoin) == -1 || IndexOf(changCoin) >=indB{
                        fmt.Println("\nexchanger can not change this coin ",changCoin)
                        changCoin = 0
                        fmt.Println("try again - input 0")
                        fmt.Println("exit - input 1")
                        fmt.Scanf("%d",&exit)
                        if exit == 0{
                            continue
                        }
                    }
            
           
            }
            }else{
                fmt.Println("Sorry, exchanger can not change 1 kopeyka")
            }
            if changCoin != 0{
                errChange(changCoin,bc,coins)
                fmt.Print(changCoin," * ",bc/changCoin)
                coins[IndexOf(changCoin)].amount -= bc/changCoin

                if bc%changCoin != 0 {
                    if coins[IndexOf(bc%changCoin)].amount<1 && bc%changCoin == 1 {
                        
                            coins[IndexOf(bc)].amount += 1
                            fmt.Println("\nexchanger do not count the amount of coins")
//                             os.Exit(1)
//                             break
                    }else if coins[IndexOf(bc%changCoin)].amount<1{
                        if coins[IndexOf(1)].amount<5{
                            coins[IndexOf(bc)].amount += 1
                            fmt.Println("\nexchanger do not count the amount of coins")
//                             os.Exit(1)
//                             break
                        }else{
                            fmt.Println(" + 1 * 5")
                            coins[IndexOf(1)].amount -= 5
                        }
                    }
                }
                fmt.Println()
                
            }else{
                fmt.Println(bc)
                coins[indB].amount += 1
            }
            fmt.Println(coins)
            locked = false 
        }
                
    }
}(coins)
    
    fmt.Println(coins)
     
}
