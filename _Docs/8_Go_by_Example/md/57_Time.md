
## 🧩 Go by Example: Zaman

Go, zamanlar ve süreler (durations) için kapsamlı destek sunar; işte bazı örnekler.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "time"
)
func main() {
    p := fmt.Println
```

---

## 🕒 Mevcut Zamanı Alma

Mevcut zamanı alarak başlayacağız.

```go
    now := time.Now()
    p(now)
```

---

## 🧱 time Struct Oluşturma

Yıl, ay, gün vb. bilgileri sağlayarak bir time struct oluşturabilirsiniz. Zamanlar her zaman bir `Location` (yani time zone) ile ilişkilidir.

```go
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)
```

---

## 🧩 Zaman Bileşenlerini Çekme

Zaman değerinin çeşitli bileşenlerini beklendiği gibi çıkarabilirsiniz.

```go
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())
```

Monday-Sunday aralığındaki `Weekday` değeri de mevcuttur.

```go
    p(then.Weekday())
```

---

## 🔍 Zamanları Karşılaştırma

Bu metotlar iki zamanı karşılaştırır; ilki ikinciden önce mi, sonra mı, yoksa aynı anda mı gerçekleşiyor, sırasıyla bunu test eder.

```go
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))
```

---

## ⏱️ İki Zaman Arasındaki Süreyi Bulma

`Sub` metodu, iki zaman arasındaki aralığı temsil eden bir `Duration` döndürür.

```go
    diff := now.Sub(then)
    p(diff)
```

Bu duration uzunluğunu çeşitli birimlerde hesaplayabiliriz.

```go
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())
```

---

## ➕➖ Duration Ekleyip Çıkarma

`Add` ile bir zamanı verilen duration kadar ileri alabilirsiniz; `-` ile ise duration kadar geri gidebilirsiniz.

```go
    p(then.Add(diff))
    p(then.Add(-diff))
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run time.go
2012-10-31 15:50:13.793654 +0000 UTC
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
25891h15m15.142266763s
25891.25420618521
1.5534752523711128e+06
9.320851514226677e+07
93208515142266763
2012-10-31 15:50:13.793654 +0000 UTC
2006-12-05 01:19:43.509120474 +0000 UTC
```

---

## 🔜 Sonraki Konu

Şimdi Unix epoch’a göre zaman kavramıyla ilgili fikre bakacağız.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Epoch.

