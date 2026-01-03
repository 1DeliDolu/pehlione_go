
## 🔤 Go by Example: String Fonksiyonları (String Functions)

Standart kütüphanedeki `strings` paketi, string’lerle ilgili birçok faydalı fonksiyon sağlar. İşte pakete dair fikir vermesi için bazı örnekler.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    s "strings"
)
```

---

## 🧷 Kısaltma: fmt.Println için Alias

Aşağıda çok kullanacağımız için `fmt.Println`’i daha kısa bir isme alias’lıyoruz.

```go
var p = fmt.Println
func main() {
```

---

## 🧰 strings Paketindeki Fonksiyonlardan Örnekler

İşte `strings` içinde mevcut fonksiyonlardan bir örnek seçkisi. Bunlar string nesnesinin üzerinde metodlar değil, paket fonksiyonları olduğu için; ilgili string’i fonksiyona ilk argüman olarak vermemiz gerekir. Daha fazla fonksiyon için `strings` paket dokümanlarına bakabilirsiniz.

```go
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
}
```

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run string-functions.go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST
```

---

## ⏭️ Sonraki Örnek

Next example: String Formatting.

