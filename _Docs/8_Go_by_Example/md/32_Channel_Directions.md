
## 🧭 Channel Yönleri

Channel’ları fonksiyon parametresi olarak kullanırken, bir channel’ın yalnızca değer göndermek mi yoksa yalnızca değer almak mı için kullanılacağını belirtebilirsiniz. Bu tür bir belirtme, programın *type-safety* seviyesini artırır.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 📤 Yalnızca Gönderim Yapan Channel Parametresi

Bu `ping` fonksiyonu yalnızca değer göndermek için bir channel kabul eder. Bu channel üzerinden değer almaya çalışmak derleme zamanında (*compile-time*) hata olur.

```go
func ping(pings chan<- string, msg string) {
    pings <- msg
}
```

---

## 📥📤 Alma ve Gönderme Yapan Channel Parametreleri

`pong` fonksiyonu, biri almak için (`pings`) ve ikincisi göndermek için (`pongs`) olmak üzere iki channel kabul eder.

```go
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```

---

## 🧩 main Fonksiyonu

```go
func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "passed message")
    pong(pings, pongs)

    fmt.Println(<-pongs)
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run channel-directions.go
passed message
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Select*.

