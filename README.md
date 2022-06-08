# zincapi

> zincapi 是 zinc 库的一个 Go 客户端实现.

### 使用

```go
package main

import (
    "gitter.top/tk/zincapi"
)

func main() {
    driver := zincapi.New("127.0.0.1:4080", "admin", "admin")
    // search for article index
    driver.Index().SetIndexName("article").Search()
}
```
