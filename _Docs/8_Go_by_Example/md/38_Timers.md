
## ⏱️ Go by Example: Zamanlayıcılar (Timers)

Gelecekte belirli bir noktada veya belirli aralıklarla tekrar ederek Go kodu çalıştırmak isteriz. Go’nun yerleşik *timer* ve *ticker* özellikleri bu iki işi de kolaylaştırır. Önce *timer*’lara, ardından *ticker*’lara bakacağız.

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

## ⏳ Tek Seferlik Olay: Timer Mantığı

*Timer*’lar gelecekte gerçekleşecek tek bir olayı temsil eder. Zamanlayıcıya ne kadar beklemek istediğinizi söylersiniz ve o süre dolduğunda bilgilendirilecek bir *channel* sağlar. Bu timer 2 saniye bekleyecek.

```go
    timer1 := time.NewTimer(2 * time.Second)
```

`<-timer1.C`, timer’ın `C` kanalında bloklanır; timer tetiklendiğinde (fired) bir değer gönderene kadar bekler.

```go
    <-timer1.C
    fmt.Println("Timer 1 fired")
```

---

## 🛑 Timer İptali: Stop Kullanımı

Sadece beklemek isteseydiniz `time.Sleep` kullanabilirdiniz. Bir timer’ın faydalı olmasının nedenlerinden biri, tetiklenmeden önce iptal edebilmenizdir. İşte bunun bir örneği.

```go
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
```

Timer2’nin gerçekten durdurulduğunu göstermek için, eğer tetiklenecek olsaydı tetiklenmesine yetecek kadar süre verelim.

```go
    time.Sleep(2 * time.Second)
}
```

---

## 🧾 Davranış Özeti

İlk timer programı başlattıktan yaklaşık `~2s` sonra tetiklenecek, ancak ikincisi tetiklenme şansı bulamadan durdurulmalıdır.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run timers.go
Timer 1 fired
Timer 2 stopped
```

---

## ⏭️ Sonraki Örnek

Next example: Tickers.

