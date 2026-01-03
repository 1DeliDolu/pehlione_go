
## 🧘‍♂️ Graceful Restart veya Stop

Web sunucunuzu *graceful* şekilde yeniden başlatmak veya durdurmak mı istiyorsunuz? Bunu yapmanın bazı yolları vardır.

Varsayılan `ListenAndServe` yerine **fvbock/endless** kullanarak bunu gerçekleştirebiliriz. Daha fazla detay için **issue #296**’ya bakın.

```go
[...]
router := gin.Default()
router.GET("/", handler)
endless.ListenAndServe(":4242", router)
```

## 🔁 endless Alternatifleri

**endless** için bir alternatif:

* **manners:** Nazikçe (*polite*) kapanan bir Go HTTP sunucusu.
* **graceful:** `http.Handler` sunucularında *graceful shutdown* sağlayan bir Go paketi.
* **grace:** *Graceful restart* ve *zero downtime deploy* için Go sunucuları.

Go **1.8** ve sonrası kullanıyorsanız bu kütüphaneleri kullanmanız gerekmeyebilir. *Graceful shutdown* için `http.Server`’ın yerleşik `Shutdown()` metodunu kullanmayı değerlendirin. Gin ile tam *graceful-shutdown* örneğine bakın.

## 🧪 http.Server Shutdown() ile Graceful Shutdown Örneği (Go 1.8+)

```go
//go:build go1.8
// +build go1.8

package main

import (
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    time.Sleep(5 * time.Second)
    c.String(http.StatusOK, "Welcome Gin Server")
  })

  srv := &http.Server{
    Addr:    ":8080",
    Handler: router.Handler(),
  }

  go func() {
    // bağlantılara hizmet ver
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("listen: %s\n", err)
    }
  }()

  // Sunucuyu 5 saniyelik timeout ile graceful kapatmak için
  // interrupt sinyalini bekle.
  quit := make(chan os.Signal, 1)
  // kill (parametresiz) varsayılan olarak syscall.SIGTERM gönderir
  // kill -2 syscall.SIGINT’tir
  // kill -9 syscall.SIGKILL’dir ama yakalanamaz, bu yüzden eklemeye gerek yok
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("Shutdown Server ...")

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := srv.Shutdown(ctx); err != nil {
    log.Println("Server Shutdown:", err)
  }
  log.Println("Server exiting")
}
```

