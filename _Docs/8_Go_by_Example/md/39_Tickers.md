
## ⏲️ Go by Example: Ticker’lar (Tickers)

*Timer*’lar gelecekte bir kez bir şey yapmak istediğinizde kullanılır — *ticker*’lar ise düzenli aralıklarla tekrar tekrar bir şey yapmak istediğinizde. İşte durdurana kadar periyodik olarak “tick” üreten bir ticker örneği.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "time"
)
func main() {
```

---

## 🔁 Periyodik Tick Mekanizması

*Ticker*’lar, *timer*’lara benzer bir mekanizma kullanır: değer gönderilen bir *channel*. Burada `select` yerleşik yapısını kullanarak her `500ms`’de gelen değerleri bekleyeceğiz.

```go
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
```

---

## 🛑 Ticker Durdurma

*Ticker*’lar, *timer*’lar gibi durdurulabilir. Bir ticker durdurulduğunda kanalına artık değer gelmez. Bizimkini `1600ms` sonra durduracağız.

```go
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
```

---

## ✅ Program Çıktısı Beklentisi

Bu programı çalıştırdığımızda, ticker’ın durdurmadan önce 3 kez tick üretmesi gerekir.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
```

---

## ⏭️ Sonraki Örnek

Next example: Worker Pools.

