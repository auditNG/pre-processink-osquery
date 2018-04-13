package transform
import(
"github.com/boltdb/bolt"
"fmt"
"strings"
"time"
"strconv"
)
func NewDataMapper() DataMapper{
  return DataMapper{}
}
type DataMapper struct {
}
func (d DataMapper) mapper(input string,field string) {
  var fields []string=strings.Split(field,"\n")
  var data []string=strings.Split(input,"\n")
  // var val []byte
  // var key []byte
  // var i int
  db,err := bolt.Open("my.db", 0600,&bolt.Options{Timeout: 1 * time.Second})

  if err != nil {
    fmt.Println(err)
  }
  for i:=0;i<len(fields)-1;i++{
    tx, err := db.Begin(true)
    b, err := tx.CreateBucketIfNotExists([]byte(fields[i]))
    if err != nil {
         fmt.Println("create bucket: %s", err)
       }
       id, _ := b.NextSequence()
       var temp=strconv.Itoa(int(id))
       b.Put([]byte(data[i]),[]byte(temp))
       if err := tx.Commit(); err != nil {
         fmt.Println(err)
  }
}
for i:=0;i<len(fields)-1;i++{
  tx, err := db.Begin(true)
  if err != nil {
    fmt.Println(err)
  }

  // b:=tx.Bucket([]byte(fields[i]))
  // v := b.Get([]byte(data[i]))
	// fmt.Printf("The answer is: %s\n", v)
  // fmt.Println(data[i])
  tx.Commit()
}



//       c:=b.Cursor()
//       for i=0;i<len(fields);i++{
//       k,_:=c.First()
//       if k!=nil{
//         if b.Get([]byte(data[i]))==nil{
//           key,val=c.Last()
//           id, _ := b.NextSequence()
//           var temp=strconv.Itoa(int(id))
//           // fmt.Println(string(key),string(val))
//           // var temp,_=strconv.Atoi(string(val))
//           // temp=temp+1
//           // var mapping=strconv.Itoa(temp)
//
//         }
//         }else {
//         b.Put([]byte(data[0]),[]byte("0"))
//       }
//       }
//
//
// 	return nil
// })
// if err!=nil{
//   fmt.Println("Data not saved")
// }
// err=db.View(func(tx *bolt.Tx)error{
// var i int
// b:=tx.Bucket([]byte("features"))
// for i=0;i<len(fields);i++{
//   fmt.Println(data[i],string(b.Get([]byte(data[i]))))
// }
// return nil
// })
db.Close()
}
