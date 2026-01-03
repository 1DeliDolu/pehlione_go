
## 🧩 Middleware (Gelişmiş)

Bu örnek, Go’da middleware’in daha gelişmiş bir sürümünün nasıl oluşturulacağını gösterir.

Bir middleware, parametrelerinden biri olarak bir **`http.HandlerFunc`** alır, onu sarar (*wrap*) ve sunucunun çağırması için yeni bir **`http.HandlerFunc`** döndürür.

Burada, birden fazla middleware’i zincirlemek (*chain*) için işleri kolaylaştıran yeni bir **`Middleware`** türü tanımlıyoruz. Bu fikir, Mat Ryers’ın *Building APIs* konuşmasından esinlenmiştir. Daha detaylı bir açıklamayı ve konuşmayı burada bulabilirsiniz.

---

## 🧱 Yeni Bir Middleware Oluşturma Mantığı

Bu parça, yeni bir middleware’in nasıl oluşturulduğunu ayrıntılı biçimde açıklar. Aşağıdaki tam örnekte, bu versiyon bazı *boilerplate* kodları azaltılarak kullanılır.

```go
func createNewMiddleware() Middleware {

    // Yeni bir Middleware oluştur
    middleware := func(next http.HandlerFunc) http.HandlerFunc {

        // Sunucu tarafından nihai olarak çağrılan http.HandlerFunc'u tanımla
        handler := func(w http.ResponseWriter, r *http.Request) {

            // ... middleware işleri yap

            // Zincirdeki bir sonraki middleware/handler'ı çağır
            next(w, r)
        }

        // Yeni oluşturulan handler'ı döndür
        return handler
    }

    // Yeni oluşturulan middleware'i döndür
    return middleware
}
```

---

## 🧪 Tam Örnek

```go
// advanced-middleware.go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging tüm istekleri path'i ve işlenmesi için geçen süreyle birlikte loglar
func Logging() Middleware {

    // Yeni bir Middleware oluştur
    return func(f http.HandlerFunc) http.HandlerFunc {

        // http.HandlerFunc'u tanımla
        return func(w http.ResponseWriter, r *http.Request) {

            // Middleware işleri yap
            start := time.Now()
            defer func() { log.Println(r.URL.Path, time.Since(start)) }()

            // Zincirdeki bir sonraki middleware/handler'ı çağır
            f(w, r)
        }
    }
}

// Method URL'nin yalnızca belirli bir method ile istenebilmesini sağlar, aksi halde 400 Bad Request döner
func Method(m string) Middleware {

    // Yeni bir Middleware oluştur
    return func(f http.HandlerFunc) http.HandlerFunc {

        // http.HandlerFunc'u tanımla
        return func(w http.ResponseWriter, r *http.Request) {

            // Middleware işleri yap
            if r.Method != m {
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
                return
            }

            // Zincirdeki bir sonraki middleware/handler'ı çağır
            f(w, r)
        }
    }
}

// Chain middlewares'leri bir http.HandlerFunc'a uygular
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, m := range middlewares {
        f = m(f)
    }
    return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
}

func main() {
    http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
    http.ListenAndServe(":8080", nil)
}
```

```bash
$ go run advanced-middleware.go
2017/02/11 00:34:53 / 0s
```

```bash
$ curl -s http://localhost:8080/
hello world
```

```bash
$ curl -s -XPOST http://localhost:8080/
Bad Request
```

