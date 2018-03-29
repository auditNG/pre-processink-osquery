import(
  "fmt"
	"os"
	"path/filepath"
  "strings"
)

func main() {
  test:=NewTest()
  files:=make([]string,0)
  files=test.run()
}
