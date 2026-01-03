
## 🧩 Go by Example: Epoch

Programlarda sık karşılaşılan bir gereksinim, Unix epoch’tan itibaren geçen saniye, milisaniye veya nanosaniye sayısını almaktır. Go’da bunu nasıl yapacağınız aşağıda gösterilmiştir.

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

## ⏱️ Unix Epoch’tan İtibaren Geçen Süreyi Alma

Unix epoch’tan itibaren geçen süreyi sırasıyla saniye, milisaniye veya nanosaniye cinsinden almak için `time.Now` ile `Unix`, `UnixMilli` veya `UnixNano` kullanın.

```go
    now := time.Now()
    fmt.Println(now)
    fmt.Println(now.Unix())
    fmt.Println(now.UnixMilli())
    fmt.Println(now.UnixNano())
```

---

## 🔁 Epoch Değerini time’a Dönüştürme

Epoch’tan itibaren geçen tamsayı saniyeleri veya nanosaniyeleri, karşılık gelen time değerine de dönüştürebilirsiniz.

```go
    fmt.Println(time.Unix(now.Unix(), 0))
    fmt.Println(time.Unix(0, now.UnixNano()))
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run epoch.go 
2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC
```

---

## 🔜 Sonraki Konu

Şimdi zamanla ilgili başka bir göreve bakacağız: zaman parse etme ve biçimlendirme.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Time Formatting / Parsing.

