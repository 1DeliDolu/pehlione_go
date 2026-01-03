
## 🧩 Go by Example: Zaman Biçimlendirme / Parse Etme

Go, pattern tabanlı layout’lar üzerinden zaman biçimlendirmeyi ve parse etmeyi destekler.

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

## 🧾 RFC3339 ile Zaman Biçimlendirme

RFC3339’a göre bir zamanı biçimlendirmeye dair temel bir örnek; ilgili layout sabitini kullanıyoruz.

```go
    t := time.Now()
    p(t.Format(time.RFC3339))
```

---

## 🧩 RFC3339 ile Zaman Parse Etme

Zaman parse etme, `Format` ile aynı layout değerlerini kullanır.

```go
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)
```

---

## 🧱 Örnek Tabanlı Layout’lar ve Özel Layout Tanımlama

`Format` ve `Parse`, örnek tabanlı layout’lar kullanır. Genellikle bu layout’lar için `time` paketindeki bir sabiti kullanırsınız, ancak özel layout’lar da sağlayabilirsiniz. Layout’lar, belirli bir zaman/string’i formatlamak/parse etmek için kullanılacak pattern’i göstermek üzere referans zaman olan `Mon Jan 2 15:04:05 MST 2006` değerini kullanmak zorundadır. Örnek zaman, gösterildiği şekilde birebir olmalıdır: yıl `2006`, saat için `15`, haftanın günü için Monday, vb.

```go
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)
```

---

## 🔢 Sayısal Gösterimler İçin Standart String Formatlama

Tamamen sayısal gösterimler için, zaman değerinin bileşenlerini çıkarıp standart string formatlamayı da kullanabilirsiniz.

```go
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())
```

---

## ⚠️ Hatalı Parse Girdilerinde Hata Döndürme

`Parse`, hatalı biçimlendirilmiş girdi için parse problemine dair açıklama içeren bir hata döndürür.

```go
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run time-formatting-parsing.go 
2014-04-15T18:00:15-07:00
2012-11-01 22:08:41 +0000 +0000
6:00PM
Tue Apr 15 18:00:15 2014
2014-04-15T18:00:15.161182-07:00
0000-01-01 20:41:00 +0000 UTC
2014-04-15T18:00:15-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Random Numbers.

