# Usage
```go
package main

import "fmt"
import "github.com/otiai10/rodeo"

func main() {
    vaquero := rodeo.TheVaquero(conf)

    _ = vaquero.Set("my_key", 12345)
    myVal, _ := vaquero.Get("my_key")
    fmt.Pritnln(myVal)
    // 12345
}
```
