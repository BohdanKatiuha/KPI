package main
import(
    "io/ioutil"
    "bufio"
    "io"
    "os"
    "fmt"
    "sort" 
    "strings"
//     "log"
    "regexp"
    "strconv"
   "time"
)

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


func sortStructs(str string,arr []string)  strukts {
    n := len(arr)
    arrDifs := make([]strukt, n)
    for i:=0; i< n; i++ {
       arrDifs[i].val, arrDifs[i].diff = arr[i], distance(str, arr[i])
    }
    sort.Sort(strukts(arrDifs))
    return arrDifs 
}

func checkError(err error) {
    if err != nil {
        panic(err)
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

 func readChunk(buf []byte, r *bufio.Reader ) int{
     n, err := r.Read(buf)
     if err != nil && err != io.EOF {
         panic(err)
     }
     return n
 }

func readWord(r *bufio.Reader)string{
    str,err := r.ReadString('\n')
        if err != nil && err != io.EOF {
            panic(err)
        }
    return str    
}

func writeFile(fileBuf []string, j int  ){
    foutput, err := os.Create("files/f"+strconv.Itoa(j)+".txt")
    checkError(err)
    defer closeFile(foutput)
    wrBuf := bufio.NewWriter(foutput)
    for _, i:=range sortStructs( "the", fileBuf) {
        if _, err := wrBuf.WriteString(i.val+" "+strconv.Itoa(i.diff) + " "+"\n"); err != nil {
            panic(err)
        }
    }
    if err = wrBuf.Flush(); err != nil {
        panic(err)
    }
    return
}

func parse(buf []byte, r *bufio.Reader) []string{
        str := string(buf)+ readWord(r)
        reg := regexp.MustCompile("[[:punct:]]")
        fina := reg.ReplaceAllString(string(str), "")
        f :=strings.Join(strings.Fields(fina), " ")
        fil := strings.Split(f, " ")
        return fil
}

func fileOpen(name string)*os.File{
    file, err := os.Open(name)
    checkError(err)
    return file
}

func wrAllFile(buf []byte, r *bufio.Reader)int{
    j:=0
    for {
        n:=readChunk(buf, r )
        if n == 0 {
            break
        }
        buffer := buf[:n] 
        fileBuf := parse(buffer,r)
        writeFile(fileBuf, j)
        j++
    }
    return j
}

func RWfile(name string, sizeBuff int) int{
    finput := fileOpen(name)
    defer closeFile(finput)
    r := bufio.NewReader(finput)
    buf := make([]byte, sizeBuff)
    return wrAllFile(buf,r)
}


func removeFiles(amountFiles int){
    for i:=0; i<amountFiles; i++{    
        os.Remove("files/f"+strconv.Itoa(i)+".txt")
    }
}


func writeLineRes(lbuf []*bufio.Scanner,b []int, w *bufio.Writer){
    if _, err := w.WriteString(lbuf[minInd(b)].Text()+"\n"); err != nil {
            panic(err)
    }
}

func MakeLineRes(lbuf []*bufio.Scanner, b []int, w *bufio.Writer){
    var err error 
    b[minInd(b)], err = strconv.Atoi(strings.Split(lbuf[minInd(b)].Text()," ")[1])
    checkError(err)
    writeLineRes(lbuf,b, w)
}

func writeInFile(lbuf []*bufio.Scanner, b []int, w *bufio.Writer,amountFiles int){

    for i:=0;i<amountFiles;i++ {
        MakeLineRes(lbuf , b , w )
        for lbuf[minInd(b)].Scan() {
            MakeLineRes(lbuf , b , w)
        }

        if i!=0{
            lbuf =append(lbuf[:minInd(b)], lbuf[minInd(b)+1:]...)  
            b =append(b[:minInd(b)], b[minInd(b)+1:]...)         
        }
    }
    if err := w.Flush(); err != nil {
        panic(err)
    }
}

func wrFinalFile(w *bufio.Writer,amountFiles int){
    var err error
    lbuf := make([]*bufio.Scanner, amountFiles)
    b := make([]int, amountFiles)
     for i:=0; i<amountFiles; i++{
         file := fileOpen("files/f"+strconv.Itoa(i)+".txt")
         defer file.Close()
         lbuf[i] = bufio.NewScanner(file)
         lbuf[i].Scan() 
         b[i], err = strconv.Atoi(strings.Split(lbuf[i].Text()," ")[1])
         checkError(err)
     }
    writeInFile(lbuf, b, w ,amountFiles)
}

func WriteResult(file string, amountFiles int){
    fileRes, err := os.Create(file)
    checkError( err)
    defer closeFile(fileRes)
    
    w := bufio.NewWriter(fileRes)
    wrFinalFile(w, amountFiles)
    removeFiles(amountFiles)
}


func WholeFile(fileInput, fileOutput string){
    file, err := ioutil.ReadFile(fileInput)
    checkError( err)
    re := regexp.MustCompile("[[:punct:]]") 
    final := strings.Split(strings.Join(strings.Fields(re.ReplaceAllString(string(file), "")), " "), " ")
    strRes:= ""
    for _, i := range sortStructs( "the", final){
       strRes  += i.val + " --- " + strconv.Itoa(i.diff) + "\n"
    }
    err = ioutil.WriteFile(fileOutput, []byte(strRes), 0777)
        checkError( err)
}

func main() {
    f:= "eTest.txt"
	start :=time.Now()
    amountFiles := RWfile(f, 1024*32)
    fmt.Println(amountFiles)
    WriteResult("buffRes.txt",amountFiles)
    end := time.Now()
    fmt.Println(end.Sub(start))
    
    start2 :=time.Now()
    WholeFile(f, "fulRes.txt")
    end2 := time.Now()
    fmt.Println(end2.Sub(start2))
}

