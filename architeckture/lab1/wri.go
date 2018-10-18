package main
import(
    "bufio"
    "io"
    "os"
    "fmt"
    "sort" 
    "strings"
    "log"
    "regexp"
    "strconv"
)

type strukt struct{
        val string `json:"value"`
        diff int `json:"difference"`
    }
    
type strukts [] strukt
    
    
    
func (s strukts) Len() int {
    return len(s)
}
    
func (s strukts) Less(i, j int) bool {
    return s[i].diff < s[j].diff
}

func (s strukts) Swap(i, j int) {  
    s[i], s[j] = s[j], s[i]
}

func out(text strukts) {
        for i := range text {
            fmt.Println(text[i])
        }
        return
}



 func min(add, del, change int) int {
		m := add
		if m > del{
            m = del
        }
        if m > change{
            m = change
        }
		return m
}


func distance(s1, s2 string) int {
	r1, r2 := []int32(s1), []int32(s2)
	n, m := len(r1), len(r2)
	if n > m {
		r1, r2 = r2, r1
		n, m = m, n
	}
	current, previous := make([]int, n+1), make([]int, n+1)
 
	for i := 1; i <= m; i++ {
		for j := range current {
			previous[j] = current[j]
			if j == 0 {
                current[j] = i
				continue
			} else {
				current[j] = 0
			}
			add, del, change := previous[j]+1, current[j-1]+1, previous[j-1]
			if r1[j-1] != r2[i-1] {
				change++
			}
			current[j] = min(add, del, change)
		}
	}
	return current[n]
}


func sortWord(str string,arr []string)  strukts {
    n := len(arr)
    arrDifs := make([]strukt, n)
    for i:=0; i< n; i++ {
       arrDifs[i].val, arrDifs[i].diff = arr[i], distance(str, arr[i])
    }
    sort.Sort(strukts(arrDifs))
    return arrDifs 
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

 func minInd (arr []int) int{
         min_i := 0
         for p:=0;p<len(arr);p++{
             if arr[min_i]>arr[p]{
                 min_i = p
             }
         }
         return min_i
 }
 
func closeFile(f *os.File){
    if err := f.Close(); err != nil {
        panic(err)
    }
}



func main() {

    finput, err := os.Open("sTest.txt")
    if err != nil {
        panic(err)
    }
    // close finput on exit and check for its returned error
    defer closeFile(finput)
    // make a read buffer
    r := bufio.NewReader(finput)
    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024*32)
    j:=0
    
    for {
        // read a chunk
        n, err := r.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        str,err := r.ReadString('\n')
        if err != nil && err != io.EOF {
            panic(err)
        }
        
        
        if n == 0 {
            break
        }
        
        buffer := buf[:n] 
//         fmt.Println(str)
        str = string(buffer)+ str
        reg := regexp.MustCompile("[[:punct:]]")
        fina := reg.ReplaceAllString(string(str), "")

//         fina = concat(fina,1)   
        fil := strings.Split(strings.Join(strings.Fields(fina), " "), " ")
        
         if fil[0] ==""{
             break            
         }
         
//          fmt.Println(fil)
        // write a chunk
        
        
        foutput, err := os.Create("files/f"+strconv.Itoa(j)+".txt")
            if err != nil {
                panic(err)
            }
            defer closeFile(foutput)
            // close foutput on exit and check for its returned error
            
            w := bufio.NewWriter(foutput)
        for _, i:=range sortWord( "the", fil) {
            if _, err := w.WriteString(i.val+" "+strconv.Itoa(i.diff) + " "+"\n"); err != nil {
                panic(err)
            }
        }
        
        if err = w.Flush(); err != nil {
            panic(err)
        }
        j++
    }

    foo, err := os.Create("res.txt")
             if err != nil {
                 checkError("1-----------------------", err)
             }
             defer func() {
                 if err := foo.Close(); err != nil {
                     checkError("2-----------------------", err)
                 }
             }()
             // close foo on exit and check for its returned error
             
    w := bufio.NewWriter(foo)
      

     lbuf := make([]*bufio.Scanner, j)

     b := make([]int, j)
    
    for i:=0; i<j; i++{
        bb, err := os.Open("files/f"+strconv.Itoa(i)+".txt")
        if err != nil {
            log.Fatal(err)
        }
        defer bb.Close()

        lbuf[i] = bufio.NewScanner(bb)
        lbuf[i].Scan() 

        b[i], err = strconv.Atoi(strings.Split(lbuf[i].Text()," ")[1])
        checkError("6----------------------", err)
    }
    
    
    b[minInd(b)],err= strconv.Atoi(strings.Split(lbuf[minInd(b)].Text()," ")[1])
            checkError("4------------------", err)
//             fmt.Println(lbuf[minInd(b)].Text())
            if _, err := w.WriteString(lbuf[minInd(b)].Text()+"\n"); err != nil {
            
                checkError("3---------------", err)
            }
//             fmt.Println(b)  
    
    for g:=0;g<j-1;g++ {
        for lbuf[minInd(b)].Scan() {
            b[minInd(b)],err= strconv.Atoi(strings.Split(lbuf[minInd(b)].Text()," ")[1])
            checkError("4------------------", err)
//             fmt.Println(lbuf[minInd(b)].Text())
            if _, err := w.WriteString(lbuf[minInd(b)].Text()+"\n"); err != nil {
                checkError("3---------------", err)
            }
        }
        lbuf =append(lbuf[:minInd(b)], lbuf[minInd(b)+1:]...)  
        b =append(b[:minInd(b)], b[minInd(b)+1:]...)         
        if _, err := w.WriteString(lbuf[minInd(b)].Text()+"\n"); err != nil {
            panic(err)
        }
    }
    if err = w.Flush(); err != nil {
            panic(err)
    }
        
    /* for i:=0; i<j; i++{    
         os.Remove("files/f"+strconv.Itoa(i)+".txt")
     }*/
}

