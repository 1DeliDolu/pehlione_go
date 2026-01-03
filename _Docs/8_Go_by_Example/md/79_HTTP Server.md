
## 🖧 HTTP Sunucusu

`net/http` paketini kullanarak temel bir HTTP sunucusu yazmak kolaydır.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "fmt"
    "net/http"
)
```

---

## 🧩 Handler Kavramı

`net/http` sunucularında temel bir kavram *handler*’lardır. Bir handler, `http.Handler` arayüzünü (*interface*) uygulayan bir nesnedir. Bir handler yazmanın yaygın bir yolu, uygun imzaya (*signature*) sahip fonksiyonlarda `http.HandlerFunc` adaptörünü kullanmaktır.

```go
func hello(w http.ResponseWriter, req *http.Request) {
```

Handler olarak görev yapan fonksiyonlar, argüman olarak bir `http.ResponseWriter` ve bir `http.Request` alır. Response writer, HTTP yanıtını doldurmak için kullanılır. Burada basit yanıtımız sadece `"hello\n"`’dur.

```go
    fmt.Fprintf(w, "hello\n")
}
```

```go
func headers(w http.ResponseWriter, req *http.Request) {
```

Bu handler, tüm HTTP istek başlıklarını (*request headers*) okuyup onları yanıt gövdesine (*response body*) geri yazarak biraz daha gelişmiş bir şey yapar.

```go
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}
```

---

## 🛣️ Route’lara Handler Kaydetme

```go
func main() {
```

Handler’larımızı sunucu rotalarına (*routes*) `http.HandleFunc` kolaylık fonksiyonuyla kaydederiz. Bu fonksiyon, `net/http` paketindeki varsayılan router’ı kurar ve argüman olarak bir fonksiyon alır.

```go
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
```

Son olarak port ve bir handler ile `ListenAndServe` çağırırız. `nil`, az önce kurduğumuz varsayılan router’ı kullanmasını söyler.

```go
    http.ListenAndServe(":8090", nil)
}
```

---

## 🧪 Sunucuyu Çalıştırma

Sunucuyu arka planda çalıştırın.

```bash
$ go run http-server.go &
```

`/hello` rotasına erişin.

```bash
$ curl localhost:8090/hello
hello
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **TCP Server**.

