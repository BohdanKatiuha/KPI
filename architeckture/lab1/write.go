package main
import(
//     "bufio"
//     "io"
//     "os"
    "fmt"
    "sort"
//     "encoding/json" 
    "strings"
    "log"
    "io/ioutil"
    "regexp"
    "strconv"
//     "encoding/csv"
)
 
// import 

type strukt struct{
        val string 
        diff int 
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

func main() {
    file, err := ioutil.ReadFile("eTest.txt")
    checkError("Cannot write to file", err)
    
    re := regexp.MustCompile("[[:punct:]]")
    final := re.ReplaceAllString(string(file), "")
 
    filll := strings.Split(strings.Join(strings.Fields(final), " "), " ")

    sss:= ""
    for _, i := range sortWord( "the", filll){
       sss  += i.val + " --- " + strconv.Itoa(i.diff) + "\n"
    }
    
    err = ioutil.WriteFile("g.txt", []byte(sss), 0777)
        checkError("Cannot write to file", err)
    
 

}

