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
  var val []byte
  var key []byte
  var i int
  db, _ := bolt.Open("my.db", 0600,&bolt.Options{Timeout: 1 * time.Second})
  err:=db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte("features"))
     if err != nil {
         return fmt.Errorf("create bucket: %s", err)
        }
      c:=b.Cursor()
      for i=0;i<len(fields);i++{
      k,_:=c.First()
      if k!=nil{
        if b.Get([]byte(data[i]))==nil{
          key,val=c.Last()
          id, _ := b.NextSequence()
          var temp=strconv.Itoa(int(id))
          // fmt.Println(string(key),string(val))
          // var temp,_=strconv.Atoi(string(val))
          // temp=temp+1
          // var mapping=strconv.Itoa(temp)
          b.Put([]byte(data[i]),[]byte(temp))
        }
        }else {
        b.Put([]byte(data[0]),[]byte("0"))
      }
      }


	return nil
})
if err!=nil{
  fmt.Println("Data not saved")
}
err=db.View(func(tx *bolt.Tx)error{
var i int
b:=tx.Bucket([]byte("features"))
for i=0;i<len(fields);i++{
  fmt.Println(data[i],string(b.Get([]byte(data[i]))))
}
return nil
})
db.Close()
}
