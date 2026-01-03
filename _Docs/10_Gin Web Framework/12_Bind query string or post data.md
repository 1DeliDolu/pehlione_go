
## 🔗 Query String veya POST Data Bind Etmek

Detaylı bilgiye bakın.

```go
package main

import (
  "log"
  "time"

  "github.com/gin-gonic/gin"
)

type Person struct {
  Name     string    `form:"name"`
  Address  string    `form:"address"`
  Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
  route := gin.Default()
  route.GET("/testing", startPage)
  route.Run(":8085")
}

func startPage(c *gin.Context) {
  var person Person
  // `GET` ise yalnızca `Form` binding engine (`query`) kullanılır.
  // `POST` ise önce `content-type` içindeki `JSON` veya `XML` kontrol edilir, ardından `Form` (`form-data`) kullanılır.
  // Daha fazlası için bkz. https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
  err := c.ShouldBind(&person)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  log.Println(person.Name)
  log.Println(person.Address)
  log.Println(person.Birthday)

  c.String(200, "Success")
}
```

---

## 🧪 Şununla Test Edin

Terminal penceresi

```bash
curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15"
```

