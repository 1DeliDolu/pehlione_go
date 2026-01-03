
## 🌐 HTTP İstemcisi

Go standart kütüphanesi, `net/http` paketinde HTTP istemcileri ve sunucuları için mükemmel destekle birlikte gelir. Bu örnekte, basit HTTP istekleri göndermek için bunu kullanacağız.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "bufio"
    "fmt"
    "net/http"
)
```

```go
func main() {
```

---

## 📥 HTTP GET İsteği Gönderme

Bir sunucuya HTTP GET isteği gönderin. `http.Get`, bir `http.Client` nesnesi oluşturup onun `Get` metodunu çağırmanın kullanışlı bir kısayoludur; yararlı varsayılan ayarlara sahip olan `http.DefaultClient` nesnesini kullanır.

```go
    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
```

---

## 🧾 Yanıt Durumunu Yazdırma

HTTP yanıt durumunu yazdırın.

```go
    fmt.Println("Response status:", resp.Status)
```

---

## 📄 Gövdenin İlk 5 Satırını Yazdırma

Yanıt gövdesinin (*response body*) ilk 5 satırını yazdırın.

```go
    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
```

---

## 🧪 Örnek Çıktı

```bash
$ go run http-clients.go
Response status: 200 OK
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Go by Example</title>
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **HTTP Server**.

