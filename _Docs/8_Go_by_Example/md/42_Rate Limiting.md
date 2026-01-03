
## 🚦 Go by Example: Hız Sınırlama (Rate Limiting)

*Rate limiting* (hız sınırlama), kaynak kullanımını kontrol etmek ve hizmet kalitesini (quality of service) korumak için önemli bir mekanizmadır. Go, *goroutine*’ler, *channel*’lar ve *ticker*’lar ile rate limiting’i zarif biçimde destekler.

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

## 🧱 Temel Rate Limiting

Önce temel rate limiting’e bakacağız. Diyelim ki gelen istekleri (incoming requests) ele alma hızımızı sınırlamak istiyoruz. Bu istekleri aynı isimli bir kanal üzerinden servis edeceğiz.

```go
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)
```

Bu `limiter` kanalı her 200 milisaniyede bir değer alacak. Bu, rate limiting şemamızdaki düzenleyicidir (regulator).

```go
    limiter := time.Tick(200 * time.Millisecond)
```

Her isteği servis etmeden önce `limiter` kanalından alım yaparak bloklandığımız için, kendimizi her 200 milisaniyede 1 istekle sınırlarız.

```go
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
```

---

## 💥 Kısa Süreli Patlamalara İzin Verme (Burst)

Rate limiting şemamızda, genel oran sınırını korurken kısa süreli istek patlamalarına izin vermek isteyebiliriz. Bunu `limiter` kanalını buffer’layarak yapabiliriz. Bu `burstyLimiter` kanalı en fazla 3 olaya kadar patlamalara izin verir.

```go
    burstyLimiter := make(chan time.Time, 3)
```

Patlamaya izin verilen kapasiteyi temsil etmek için kanalı dolduralım.

```go
    for range 3 {
        burstyLimiter <- time.Now()
    }
```

Her 200 milisaniyede bir, `burstyLimiter`’a (3 olan limitine kadar) yeni bir değer eklemeyi deneyeceğiz.

```go
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()
```

Şimdi 5 tane daha gelen isteği simüle edelim. Bunların ilk 3’ü `burstyLimiter`’ın burst yeteneğinden faydalanacak.

```go
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
```

---

## ✅ Çalıştırma Sonucu Gözlemi

Programımızı çalıştırdığımızda, ilk istek grubunun istenildiği gibi yaklaşık her `~200ms`’de bir işlendiğini görürüz.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC
```

İkinci istek grubunda ise, burst edilebilir rate limiting sayesinde ilk 3 isteği hemen servis ederiz; kalan 2 isteği ise her biri için yaklaşık `~200ms` gecikmeyle servis ederiz.

```text
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```

---

## ⏭️ Sonraki Örnek

Next example: Atomic Counters.

