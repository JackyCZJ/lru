# lru

usage:
```go
package main
import (
	"fmt"
	"github.com/jackyczj/lru"
)

func main(){
   sc :=  lru.NewSafeCache(lru.New(10))
   sc.Set("key","value")
   value := sc.Get("key")
   fmt.Println(value)  
}
```
