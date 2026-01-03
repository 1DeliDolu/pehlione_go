
## 🚫⛓️ Bloklamayan Channel İşlemleri

Channel’larda temel gönderme (*send*) ve alma (*receive*) işlemleri bloklayıcıdır (*blocking*). Ancak `select` yapısını bir `default` dalı ile kullanarak bloklamayan gönderme, alma ve hatta bloklamayan çoklu (*multi-way*) `select` işlemleri uygulayabiliriz.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 🧩 main Fonksiyonu

```go
func main() {
    messages := make(chan string)
    signals := make(chan bool)
```

---

## 📥 Bloklamayan Alma

Bu bir bloklamayan alma örneğidir. `messages` üzerinde bir değer mevcutsa, `select` o değerle birlikte `<-messages` durumunu seçer. Eğer yoksa, hemen `default` durumunu seçer.

```go
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }
```

---

## 📤 Bloklamayan Gönderme

Bloklamayan gönderme benzer şekilde çalışır. Burada `msg`, `messages` channel’ına gönderilemez; çünkü channel buffer’sızdır ve alıcı yoktur. Bu nedenle `default` durumu seçilir.

```go
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
```

---

## 🔀 Bloklamayan Çoklu Select

Yukarıda `default` dalının üstünde birden fazla `case` kullanarak çok yönlü (*multi-way*) bloklamayan `select` uygulayabiliriz. Burada hem `messages` hem de `signals` üzerinde bloklamayan alma işlemlerini deniyoruz.

```go
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run non-blocking-channel-operations.go
no message received
no message sent
no activity
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Closing Channels*.

