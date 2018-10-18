package main
import(
    "fmt"
    "time"
    "math/rand"
    "os"
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

func RandCoin() int {
    return mon[rand.Intn(len(mon))]
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
        os.Exit(1)
     }
}


func TestAndSet (n *int) int { 
// semantics: 
// if *n == 1 { 
//   return 1 
// } 
// *n = 1 
// return 0 
} 

// Reason: Nice implementation of Lock/Unlock: 


func Lock() { 
  for TestAndSet (locked) == 1 { 
    // do nothing, maybe sleep some usecs 
  } 

func Unlock() { 
  locked = 0 
} 


func main(){
    
    var coins = []coin{
        coin{1,6},
        coin{2,50},
        coin{5,0},
        coin{10,20},
        coin{25,10},
        coin{50,4},
        coin{100,0}}
    common := 0
//     common := true
    
        fmt.Println(coins)
        beginCoin:=RandCoin()
        indB := IndexOf(beginCoin)
        coins[indB].amount += 1
        fmt.Println("Begin coin: ",beginCoin)
//         common := 0
        
    
    exit := -1
    changCoin := 0
    if IndexOf(beginCoin) != 0{
         for true{
            fmt.Print("select a number coin to be changed from ",)
            printPart(IndexOf(beginCoin))
            
            fmt.Scanf("%d",&changCoin)
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
            
            break
         }
    }else{
        fmt.Println("Sorry, exchanger can not change 1 kopeyka")
    }
    if changCoin != 0{
        errChange(changCoin,beginCoin,coins)
        fmt.Print(changCoin," * ",beginCoin/changCoin)
        coins[IndexOf(changCoin)].amount -= beginCoin/changCoin

        if beginCoin%changCoin != 0 {
            if coins[IndexOf(beginCoin%changCoin)].amount<1 && beginCoin%changCoin == 1 {
                
                    coins[IndexOf(beginCoin)].amount += 1
                    fmt.Println("\nexchanger do not count the amount of coins")
                    os.Exit(1)
            }else if coins[IndexOf(beginCoin%changCoin)].amount<1{
                if coins[IndexOf(1)].amount<5{
                    coins[IndexOf(beginCoin)].amount += 1
                    fmt.Println("\nexchanger do not count the amount of coins")
                    os.Exit(1)
                }else{
                    fmt.Println(" + 1 * 5")
                    coins[IndexOf(1)].amount -= 5
                }
            }
        }
        fmt.Println()
        
    }else{
        fmt.Println(beginCoin)
        coins[indB].amount += 1
    }
    fmt.Println(coins)
}




















