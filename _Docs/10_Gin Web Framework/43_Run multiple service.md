
## 🧩 Birden Fazla Servis Çalıştırma (Run multiple service)

### 🧠 Açıklama

Aşağıdaki örnek, aynı Go uygulaması içinde **iki ayrı HTTP sunucusunu** (iki farklı portta) aynı anda çalıştırmayı gösterir. Bunu yapmak için `golang.org/x/sync/errgroup` kullanılır.

* `router01()` → `:8080` portunda çalışır
* `router02()` → `:8081` portunda çalışır
* `g.Go(...)` ile her sunucu ayrı bir goroutine içinde başlatılır
* `g.Wait()` ile hepsinin bitmesi beklenir; herhangi biri hata verirse uygulama hata ile sonlandırılır

```go
package main

import (
  "log"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
  "golang.org/x/sync/errgroup"
)

var (
  g errgroup.Group
)

func router01() http.Handler {
  e := gin.New()
  e.Use(gin.Recovery())
  e.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code":  http.StatusOK,
        "message": "Welcome server 01",
      },
    )
  })

  return e
}

func router02() http.Handler {
  e := gin.New()
  e.Use(gin.Recovery())
  e.GET("/", func(c *gin.Context) {
    c.JSON(
      http.StatusOK,
      gin.H{
        "code":  http.StatusOK,
        "message": "Welcome server 02",
      },
    )
  })

  return e
}

func main() {
  server01 := &http.Server{
    Addr:         ":8080",
    Handler:      router01(),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  server02 := &http.Server{
    Addr:         ":8081",
    Handler:      router02(),
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  g.Go(func() error {
    return server01.ListenAndServe()
  })

  g.Go(func() error {
    return server02.ListenAndServe()
  })

  if err := g.Wait(); err != nil {
    log.Fatal(err)
  }
}
```

### 🧪 Test Etme

```bash
curl http://localhost:8080/
```

```bash
curl http://localhost:8081/
```

