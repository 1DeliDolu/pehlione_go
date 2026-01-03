
## 🎛️ Select

Go’nun `select` yapısı, birden fazla channel işlemini aynı anda beklemenizi sağlar. Goroutine’leri ve channel’ları `select` ile birleştirmek, Go’nun güçlü özelliklerinden biridir.

---

## ▶️ Çalıştır

```go
package main

import (
    "fmt"
    "time"
)
```

---

## 🧩 main Fonksiyonu

Örneğimizde iki channel üzerinde `select` yapacağız.

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
```

Her channel, belirli bir süre sonra bir değer alacak; bu da örneğin eşzamanlı goroutine’lerde çalışan bloklayan RPC işlemlerini simüle etmek içindir.

```go
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
```

`select` kullanarak bu iki değeri aynı anda bekleyeceğiz ve her biri geldiğinde yazdıracağız.

```go
    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

---

## 🖨️ Program Çıktısı

Beklendiği gibi önce `"one"`, sonra `"two"` değerlerini alırız.

```bash
$ time go run select.go
received one
received two
```

Toplam çalışma süresinin yalnızca yaklaşık `~2` saniye olduğuna dikkat edin; çünkü `1` ve `2` saniyelik `Sleep` çağrıları eşzamanlı olarak çalışır.

```bash
real    0m2.245s
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Timeouts*.

