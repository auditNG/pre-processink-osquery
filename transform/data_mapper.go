package transform
import(
"github.com/boltdb/bolt"
"fmt"
"strings"
"time"
)
func NewDataMapper() DataMapper{
  return DataMapper{}
}
type DataMapper struct {
}
func (d DataMapper) mapper(input string,field string) {
  var fields []string=strings.Split(field,"\n")
  var data []string=strings.Split(input,"\n")
  // fmt.Println(input)
  // fmt.Println(field)
  db, _ := bolt.Open("my.db", 0600,&bolt.Options{Timeout: 1 * time.Second})
  db.Update(func(tx *bolt.Tx) error {
    for index,val:=range fields{
	     b, err := tx.CreateBucketIfNotExists([]byte(val))
	      if err != nil {
		        return fmt.Errorf("create bucket: %s", err)
	         }

    b.Put([]byte(data[index]),[]byte("0"))

    }

	return nil
})
db.Close()

}
