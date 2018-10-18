package main

import (
    "fmt"
)

type proc struct{
    name string
    size int
}

type memory struct{
    number int
    beginAdr int
    size int
    free bool
}

var(
    allSize = 60
    countSection = 1
    number = 0
)

func scan(processes []proc) (proc,[]proc){
    pr := processes[0]
    return pr,append(processes[:0],processes[1:]...)
}

func download(process proc, memSections []memory)[]memory{
    j:= 0
    for _,i:= range memSections{
        if i.free==true && i.size >= process.size{
            if i.size == process.size{
                memSections[j] =memory{i.number, i.beginAdr,i.size,false}
                return memSections
            }else{
                number++
                memSections = append(append(memSections[:j], []memory{memory{i.number,i.beginAdr + process.size, i.size - process.size,true}}...),memSections[j:]...)
                memSections[j] =memory{number,i.beginAdr,process.size,false}
                return memSections
            }
        }
        j++
    }
    fmt.Println("Error")
    return memSections
}

func free(n int, memSections []memory) []memory{
    j:= 0
    for _,i:= range memSections{
        if i.number == n {
            memSections[j].free = true
            return memSections
        }
        j++
    }
    return memSections
}


func deFragment(memSections []memory) []memory{
     freeSize := 0
     for i:=0;i<len(memSections);i++{
        if memSections[i].free == true{
              freeSize += memSections[i].size
              memSections = append(memSections[:i],memSections[i+1:]...)
         }         
    }
    memSections = append(memSections, memory{0,memSections[len(memSections)-1].beginAdr+memSections[len(memSections)-1].size,freeSize,true})
    return memSections
}


func main(){
    
    fmt.Println("allSize= ",allSize)
    
    var processes = []proc{
        proc{"1",1},
        proc{"2",2},
        proc{"3",3},
        proc{"4",4},
        proc{"5",5},
        proc{"6",6},
        proc{"7",7}}
    fmt.Println(processes)    
    
    var memSections =[]memory{
        memory{0,0,allSize,true}}
        
     fmt.Println(memSections)
     var pr proc
     for  len(processes)>0 {
         
         pr,processes = scan(processes)
         fmt.Println("pr =" ,pr)
         memSections = download(pr,memSections)
     
     }
     
     fmt.Println(memSections,"\n\n")
    
    memSections = free(2,memSections)
    fmt.Println(memSections)
    memSections = download(proc{"8",1},memSections)
    fmt.Println(memSections)
    memSections = deFragment(memSections)
    fmt.Println("deFragment: ",memSections)
}
