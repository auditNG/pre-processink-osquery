package transform
// package main
//
// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
//   "strings"
//   "encoding/csv"
//   "sort"
//   "github.com/boltdb/bolt"
//   "time"
//   "strconv"
//   // "io"
// )
// func NewTest() Test {
// 	return Test{}
// }
//
// type Test struct {
//
// }
// func(t Test) run() ([]string, error) {
// 	searchDir := "../"
//
// 	fileList := make([]string, 0)
//   files:=make([]string,0)
// 	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
//     if strings.Contains(path,".csv")!=false{
//     fileList = append(fileList, path)
//     }
// 		return err
// 	})
//
// 	if e != nil {
// 		panic(e)
// 	}
//
// 	// for _, file := range fileList {
//   //
//   //   // files=append(files,strings.TrimPrefix(file,"../"))
// 	// }
//   for _,val:=range files{
//     fmt.Println(val)
//   }
//
// 	return fileList, nil
// }
// type Pair struct {
// 	Key   string
// 	Value int
// }
//
// type PairList []Pair
//
// func (p PairList) Len() int           { return len(p) }
// func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
//
//
// func getdestmap(db *bolt.DB)map[string]int{
//   dest:=make(map[string]int)
//   tx, err := db.Begin(true)
//   if err != nil {
//     fmt.Println(err)
//   }
//   b, err := tx.CreateBucketIfNotExists([]byte("destmap"))
//   if err != nil {
//        fmt.Println("create bucket: %s", err)
//      }
//   b = tx.Bucket([]byte("destmap"))
//   c := b.Cursor()
//   key,_:=c.First()
//   if key!=nil{
//   for k, v := c.First(); k != nil; k, v = c.Next() {
//     var temp,_=strconv.Atoi(string(v))
//     dest[string(k)]=temp
// }
// }
// if err = tx.Commit(); err != nil {
//     fmt.Println(err)
// }
// return dest
// }
// func setdestmap(db *bolt.DB,input map[string]int,files []string)(map[string]map[string]int,map[string]map[int]string){
//   var count int
//   src:=make(map[string]map[string]int)
//   sourceCont:=make(map[string]map[int]string)
//   count=len(input)
//   tx, err := db.Begin(true)
//   if err != nil {
//     fmt.Println(err)
//   }
//   b := tx.Bucket([]byte("destmap"))
//   for _,val:=range files{
//   aMap:=make(map[string]int)
//   cMap:=make(map[int]string)
//   dest:=make(map[string]int)
//   read,_:=os.OpenFile(val,os.O_RDONLY,0644)
//   reader:=csv.NewReader(read)
//   record,_:=reader.Read()
//   for index,value:=range record{
//     aMap[value]=index
//     cMap[index]=value
//     ok := b.Get([]byte(value))
//     if ok==nil {
//       // fmt.Println(index,value,strconv.Itoa(count))
//         dest[value]=count
//         b.Put([]byte(value),[]byte(strconv.Itoa(count)))
//         count++
//       }
//
//     }
//     src[val]=aMap
//     sourceCont[val]=cMap
//     }
//   if err = tx.Commit(); err != nil {
//       fmt.Println(err)
//   }
//   return src,sourceCont
// }
//
//
// func main() {
//   test:=NewTest()
//   files:=make([]string,0)
//   files,_=test.run()
//    db, _ := bolt.Open("dest.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
//   dest:=make(map[string]int)
//   dest=getdestmap(db)
//   // src:=make(map[string]map[string]int)
//   sourceCont:=make(map[string]map[int]string)
//   var fieldNames=string("")
//   write, _ :=os.OpenFile("main.csv",os.O_WRONLY|os.O_CREATE, 0644)
//   _,sourceCont=setdestmap(db,dest,files)
//   dest=getdestmap(db)
//     p := make(PairList, len(dest))
//     i := 0
// 	for k, v := range dest {
// 		p[i] = Pair{k, v}
// 		i++
// 	}
//   sort.Sort(p)
//   for ind, k := range p {
//     if ind!=len(p)-1{
// 		fieldNames=fieldNames+k.Key+","
//     }else{
//       fieldNames=fieldNames+k.Key+"\n"
//     }
// 	}
//      write.Write([]byte(fieldNames))
//      write.Close()
//      write, _ =os.OpenFile("main.csv",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
//    for _,val:=range files{
//    read,_:=os.OpenFile(val,os.O_RDONLY,0644)
//    reader:=csv.NewReader(read)
//    record,_:=reader.ReadAll()
//
//    for rno,rows:=range record{
//      if rno==0{
//        continue
//      }
//   data:=make([]string,len(dest),len(dest))
//   for i:=0;i<len(dest);i++{
//     data[i]=" "
//   }
//    for index,value:=range rows{
//      data[dest[sourceCont[val][index]]]=value
//      // fmt.Println(data)
//    }
//    var outputLine =strings.Join(data,",")
//    outputLine=strings.TrimRight(outputLine,",")
//    outputLine=outputLine+"\n"
//     write.Write([]byte(outputLine))
//   //
//    }
//    }
//    // db.Close()
// //
//  }
