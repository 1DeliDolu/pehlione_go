
## 🧾 Multipart/Urlencoded Bağlama

```go
package main

import (
  "github.com/gin-gonic/gin"
)

type LoginForm struct {
  User     string `form:"user" binding:"required"`
  Password string `form:"password" binding:"required"`
}

func main() {
  router := gin.Default()
  router.POST("/login", func(c *gin.Context) {
    // multipart form’u açık bir binding bildirimiyle bağlayabilirsiniz:
    // c.ShouldBindWith(&form, binding.Form)
    // veya ShouldBind metodu ile otomatik bağlamayı (autobinding) kullanabilirsiniz:
    var form LoginForm
    // bu durumda uygun binding otomatik olarak seçilecektir
    if c.ShouldBind(&form) == nil {
      if form.User == "user" && form.Password == "password" {
        c.JSON(200, gin.H{"status": "you are logged in"})
      } else {
        c.JSON(401, gin.H{"status": "unauthorized"})
      }
    }
  })
  router.Run(":8080")
}
```

## 🧪 Test Etme

Terminal penceresi

```bash
$ curl -v --form user=user --form password=password http://localhost:8080/login
```

