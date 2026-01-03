
## 🔎 Sadece Query String Bağlama

`ShouldBindQuery` fonksiyonu yalnızca **query parametrelerini** bağlar, **post data**’yı bağlamaz. Detaylı bilgiye bakın.

```go
package main

import (
  "log"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name    string `form:"name"`
  Address string `form:"address"`
}

func main() {
  route := gin.Default()
  route.Any("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  if c.ShouldBindQuery(&person) == nil {
    log.Println("====== Sadece Query String ile Bağla ======")
    log.Println(person.Name)
    log.Println(person.Address)
  }
  c.String(200, "Success")
}
```


