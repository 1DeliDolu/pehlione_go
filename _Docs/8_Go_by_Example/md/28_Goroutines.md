
## 🧵 Goroutine’ler

Bir *goroutine*, yürütmenin hafif (*lightweight*) bir iş parçacığı (*thread*) türüdür.

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

## 🔁 Örnek Fonksiyon

```go
func f(from string) {
    for i := range 3 {
        fmt.Println(from, ":", i)
    }
}
```

---

## 🧩 main Fonksiyonu

`f(s)` şeklinde bir fonksiyon çağrımız olduğunu varsayalım. İşte bunu normal şekilde, senkron olarak çalıştırarak nasıl çağırırız:

```go
func main() {
    f("direct")
```

Bu fonksiyonu bir goroutine içinde çağırmak için `go f(s)` kullanın. Bu yeni goroutine, çağıran goroutine ile eşzamanlı (*concurrently*) olarak çalışır.

```go
    go f("goroutine")
```

Anonim bir fonksiyon çağrısı için de goroutine başlatabilirsiniz.

```go
    go func(msg string) {
        fmt.Println(msg)
    }("going")
```

Artık iki fonksiyon çağrımız ayrı goroutine’lerde asenkron olarak çalışıyor. Bitmelerini bekleyin (daha sağlam bir yaklaşım için `WaitGroup` kullanın).

```go
    time.Sleep(time.Second)
    fmt.Println("done")
}
```

---

## 🖨️ Program Çıktısı

Bu programı çalıştırdığımızda, önce bloklayan çağrının çıktısını, ardından iki goroutine’in çıktısını görürüz. Goroutine’lerin çıktıları iç içe geçebilir, çünkü goroutine’ler Go runtime tarafından eşzamanlı olarak çalıştırılır.

```bash
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done
```

---

## ⏭️ Sonraki Konu

Sonraki olarak, eşzamanlı Go programlarında goroutine’leri tamamlayan bir yapıya bakacağız: *channels*.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Channels*.

