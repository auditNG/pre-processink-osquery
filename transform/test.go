// package transform
package main

import (
	"fmt"
	"os"
	"path/filepath"
  "strings"
  "encoding/csv"
  "sort"
)
func NewTest() Test {
	return Test{}
}

type Test struct {

}

func(t Test) run() ([]string, error) {
	searchDir := "../"

	fileList := make([]string, 0)
  files:=make([]string,0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
    if strings.Contains(path,".csv")!=false{
    fileList = append(fileList, path)
    }
		return err
	})

	if e != nil {
		panic(e)
	}

	// for _, file := range fileList {
  //
  //   // files=append(files,strings.TrimPrefix(file,"../"))
	// }
  for _,val:=range files{
    fmt.Println(val)
  }

	return fileList, nil
}
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }


func main() {
  test:=NewTest()
  files:=make([]string,0)
  files,_=test.run()
  var count int
  count=0
  dest:=make(map[string]int)
  src:=make(map[string]map[string]int)
  sourceCont:=make(map[string]map[int]string)
  var fieldNames=string("")
  write, _ :=os.OpenFile("main.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
  for _,val:=range files{
  aMap:=make(map[string]int)
  cMap:=make(map[int]string)
  read,_:=os.OpenFile(val,os.O_RDONLY,0644)
  reader:=csv.NewReader(read)
  record,_:=reader.Read()
  for index,value:=range record{
    aMap[value]=index
    cMap[index]=value
    if _, ok := dest[value]; !ok {
        dest[value]=count
        count++
      }
    }
    src[val]=aMap
    sourceCont[val]=cMap
    }
    p := make(PairList, len(dest))
    i := 0
	for k, v := range dest {
		p[i] = Pair{k, v}
		i++
	}
  sort.Sort(p)
  for ind, k := range p {
    if ind!=len(p)-1{
		fieldNames=fieldNames+k.Key+","
    }else{
      fieldNames=fieldNames+k.Key+"\n"
    }
	}
     write.Write([]byte(fieldNames))

//   var space int
//   space=0
//   var temp int
//   temp=0
   for _,val:=range files{
   read,_:=os.OpenFile(val,os.O_RDONLY,0644)
   reader:=csv.NewReader(read)
   record,_:=reader.ReadAll()
//   for i:=0;i<space;i++{
//     fieldNames=fieldNames+","
//   }
   for rno,rows:=range record{
     if rno==0{
       continue
     }
  data:=make([]string,len(dest),len(dest))
   for index,value:=range rows{
     data[dest[sourceCont[val][index]]]=value

   }
   var outputLine =strings.Join(data,",")
   outputLine=strings.TrimRight(outputLine,",")
   outputLine=outputLine+"\n"
    write.Write([]byte(outputLine))

    }
   }
//
 }
