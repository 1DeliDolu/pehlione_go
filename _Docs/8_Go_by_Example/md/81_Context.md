
## 🧠 Context

Önceki örnekte basit bir HTTP sunucusu kurmaya baktık. HTTP sunucuları, iptali (*cancellation*) kontrol etmek için `context.Context` kullanımını göstermek açısından faydalıdır. Bir *Context*; API sınırları ve goroutine’ler boyunca deadline’ları, iptal sinyallerini ve istek kapsamlı (*request-scoped*) diğer değerleri taşır.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "fmt"
    "net/http"
    "time"
)
```

---

## 🧩 Handler İçinde Context Kullanımı

```go
func hello(w http.ResponseWriter, req *http.Request) {
```

`net/http` mekanizması her istek için bir `context.Context` oluşturur ve bu, `Context()` metodu ile erişilebilir.

```go
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")
```

İstemciye yanıt göndermeden önce birkaç saniye bekleyin. Bu, sunucunun yaptığı bir işi simüle edebilir. Çalışırken, işi iptal edip mümkün olan en kısa sürede dönmemiz gerektiğine dair bir sinyal için context’in `Done()` kanalını izleyin.

```go
    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
```

Context’in `Err()` metodu, `Done()` kanalının neden kapandığını açıklayan bir hata döndürür.

```go
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}
```

---

## 🛣️ Route Kaydı ve Sunucuyu Başlatma

```go
func main() {
```

Önceki gibi, handler’ımızı `"/hello"` rotasına kaydeder ve sunucuyu çalıştırırız.

```go
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}
```

---

## 🧪 Sunucuyu Çalıştırma ve İptali Simüle Etme

Sunucuyu arka planda çalıştırın.

```bash
$ go run context.go &
```

`/hello` için bir istemci isteğini simüle edin; iptali işaretlemek için başladıktan kısa süre sonra `Ctrl+C`’ye basın.

```bash
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Spawning Processes**.

