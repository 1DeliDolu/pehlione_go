
## 🧪 Channel’lar

*Channel*’lar, eşzamanlı (*concurrent*) goroutine’leri birbirine bağlayan borulardır (*pipes*). Bir goroutine’den channel içine değer gönderebilir ve bu değerleri başka bir goroutine’de alabilirsiniz.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 🧩 main Fonksiyonu

`make(chan val-type)` ile yeni bir channel oluşturun. Channel’lar, taşıdıkları değerlerin tipine göre türlendirilir (*typed*).

```go
func main() {
    messages := make(chan string)
```

`channel <-` sözdizimini kullanarak bir channel içine değer gönderin. Burada, yukarıda oluşturduğumuz `messages` channel’ına yeni bir goroutine’den `"ping"` gönderiyoruz.

```go
    go func() { messages <- "ping" }()
```

`<-channel` sözdizimi bir değeri channel’dan alır (*receive*). Burada yukarıda gönderdiğimiz `"ping"` mesajını alıp ekrana yazdıracağız.

```go
    msg := <-messages
    fmt.Println(msg)
}
```

---

## 🖨️ Program Çıktısı

Programı çalıştırdığımızda `"ping"` mesajı channel aracılığıyla bir goroutine’den diğerine başarıyla aktarılır.

```bash
$ go run channels.go
ping
```

---

## ⛔ Varsayılan Bloklama Davranışı

Varsayılan olarak, gönderme (*send*) ve alma (*receive*) işlemleri hem gönderen hem de alan hazır olana kadar bloklanır. Bu özellik sayesinde, programın sonunda `"ping"` mesajını beklemek için başka bir senkronizasyon mekanizması kullanmamıza gerek kalmadı.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Channel Buffering*.

