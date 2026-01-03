
## 🗃️ Channel Buffering

Varsayılan olarak channel’lar *unbuffered*’dır; yani yalnızca gönderilen değeri almaya hazır karşılık gelen bir alma (`<- chan`) varsa gönderimleri (`chan <-`) kabul ederler. *Buffered* channel’lar ise, bu değerler için karşılık gelen bir alıcı olmadan da sınırlı sayıda değeri kabul eder.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 🧩 main Fonksiyonu

Burada, en fazla 2 değeri buffer’layabilen bir `string` channel’ı oluşturuyoruz.

```go
func main() {
    messages := make(chan string, 2)
```

Bu channel buffer’lı olduğu için, eşzamanlı bir alma işlemi olmadan bu değerleri channel içine gönderebiliriz.

```go
    messages <- "buffered"
    messages <- "channel"
```

Daha sonra bu iki değeri her zamanki gibi alabiliriz.

```go
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run channel-buffering.go
buffered
channel
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Channel Synchronization*.

