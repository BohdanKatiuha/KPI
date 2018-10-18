package main
import(
    "fmt"
    "strconv"
)

type process struct{
        name string 
        arrivalTime int 
        burstTime int
}

func fifo(processes []process){
    n:= len(processes)
    fmt.Println(processes)
    fmt.Println("number processes = ",n)

//     startTime := make([]int, n)
//     responseTime := make([]int, n)
    
    stime := 0
    fulltime, delaytime := 0,0
    fmt.Printf("Process\t|  Arrival time\t|  Burst Time\t|  Start time\t|  Finish Time\t|  Delay time\t|  Full Time\t|\n");
    fmt.Printf("------------------------------------------------------------------------------------------------------\n");
 for i:=0;i<n;i++{
     if processes[0].arrivalTime > stime{
        stime = processes[0].arrivalTime
    }
     fmt.Printf("%s\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\n", processes[0].name, processes[0].arrivalTime, processes[0].burstTime,stime, stime+processes[0].burstTime, stime,stime+processes[0].burstTime);
     fulltime += stime+processes[0].burstTime
     delaytime+= stime
     stime += processes[0].burstTime;
     processes = append(processes[:0],processes[1:]...)   
     if n == 0{
         break
    }
}
fmt.Println("average full time ",float64(fulltime)/float64(n))
fmt.Println("average delay time ",float64(delaytime)/float64(n))
}

func round(processes []process){
      n:= len(processes)
    fmt.Println(processes)
    fmt.Println("number processes = ",n)
    q := 2
    fmt.Println("quant =",q)

    startTime := make([]int, n)
    responseTime := make([]int, n)
 
    fulltime, delaytime := 0,0
    for i:=0;i<n;i++{
        responseTime[i] = processes[i].burstTime;
    }
    
    fmt.Printf("Process\t|  Arrival time\t|  Burst Time\t|  Start time\t|  Finish Time\t|  Delay time\t|  Full Time\t|\n");
    fmt.Printf("------------------------------------------------------------------------------------------------------\n");
    flag:=0
     fl:=0
    i, time := 0,0;
    count := n
    str:="0 "
    for count != 0 {
       
    if processes[i].arrivalTime <= time{
        if responseTime[i]==processes[i].burstTime{
            startTime[i] = time
        }
    }else {
       
        for true{
            
            i = 0
            for j:=0;j<n;j++{
                if processes[j].arrivalTime > time{
                    i++
                }else{
                    fl = 1
                    break
                }
            }
            if fl == 1{
                if responseTime[i]==processes[i].burstTime{
                    startTime[i] = time
                }
                break
            }else{
                time++
            }
            
        }

    }
     
    
    if responseTime[i] <= q && responseTime[i]>0 {
        time += responseTime[i];
        responseTime[i] = 0;
        flag = 1;
        str+= processes[i].name +" "+ strconv.Itoa(time)+"; "
    }else if responseTime[i] > 0   {
        responseTime[i]-= q;
        time += q;
        str+= processes[i].name +" "+ strconv.Itoa(time)+"; "
    }
    
    if responseTime[i] == 0 && flag == 1    {
        
         
        count--;
        fmt.Printf("%s\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\t%d\t|\n", processes[i].name, processes[i].arrivalTime, processes[i].burstTime,startTime[i], time, time - processes[i].burstTime,time -startTime[i]);
        flag = 0;
        fulltime += time -startTime[i]
        delaytime += time - processes[i].burstTime
    }

    if i == n-1      {
        i = 0;
    }else if processes[i].arrivalTime <= time{
        i++
        
    } 
}  
fmt.Println(str)
fmt.Println("average full time ",float64(fulltime)/float64(n))
fmt.Println("average delay time ",float64(delaytime)/float64(n))
}

func main() {
//     s := person{name: "Sean", age: 50}
    
    var processes = []process{
        process{"s1",2,8},
        process{"s2",2,5},
        process{"s3",4,1},
        process{"s4",5,2}}
   fmt.Println("round")
   round(processes)
   fmt.Println("\n\n\nfifo")
   fifo(processes)
}
