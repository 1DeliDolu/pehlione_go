
## 🔄 Channel Senkronizasyonu

Goroutine’ler arasında yürütmeyi senkronize etmek için channel’ları kullanabiliriz. Aşağıda, bir goroutine’in bitmesini beklemek için bloklayan bir alma (*blocking receive*) kullanan bir örnek yer alıyor. Birden fazla goroutine’in bitmesini beklerken, `WaitGroup` kullanmayı tercih edebilirsiniz.

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

## 🧑‍🏭 worker Fonksiyonu

Bu, bir goroutine içinde çalıştıracağımız fonksiyondur. `done` channel’ı, bu fonksiyonun işi bittiğinde başka bir goroutine’e bildirim yapmak için kullanılacaktır.

```go
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
```

Bittiğimizi bildirmek için bir değer gönderin.

```go
    done <- true
}
```

---

## 🧩 main Fonksiyonu

Bir worker goroutine’i başlatın ve bildirim yapacağı channel’ı ona verin.

```go
func main() {
    done := make(chan bool, 1)
    go worker(done)
```

Channel üzerinden worker’dan bir bildirim alana kadar bloklayın.

```go
    <-done
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run channel-synchronization.go
working...done
```

---

## ⚠️ Not

Bu programdan `<-done` satırını kaldırırsanız, program worker işini bitirmeden önce çıkabilir; bazı durumlarda worker daha başlamadan bile program sonlanabilir.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Channel Directions*.

