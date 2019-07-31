package aPackage

import (
    "fmt"
)

func A() {
    fmt.Println("This is function A!")
}
```

`aPackage.go` 的第二部分代码如下：

```go
func B() {
    fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21