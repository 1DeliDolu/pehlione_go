
## 🧩 Go by Example: Düzenli İfadeler

Go, düzenli ifadeler (regular expressions) için yerleşik destek sunar. Aşağıda Go’da regexp ile ilgili yaygın görevlerin bazı örnekleri yer alır.

---

## ▶️ Çalıştırma

```go
package main
import (
    "bytes"
    "fmt"
    "regexp"
)
func main() {
```

---

## ✅ Bir Pattern’in String ile Eşleşmesini Test Etme

Bu, bir pattern’in bir string ile eşleşip eşleşmediğini test eder.

```go
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)
```

Yukarıda bir string pattern’i doğrudan kullandık, ancak diğer regexp görevleri için optimize edilmiş bir `Regexp` struct’ı derlemek (`Compile`) gerekir.

```go
    r, _ := regexp.Compile("p([a-z]+)ch")
```

---

## 🧪 Regexp Struct Üzerindeki Metotlar

Bu struct’lar üzerinde birçok metot vardır. İşte daha önce gördüğümüze benzer bir eşleşme testi.

```go
    fmt.Println(r.MatchString("peach"))
```

Bu, regexp için eşleşen metni bulur.

```go
    fmt.Println(r.FindString("peach punch"))
```

Bu da ilk eşleşmeyi bulur ancak eşleşen metin yerine, eşleşmenin başlangıç ve bitiş indekslerini döndürür.

```go
    fmt.Println("idx:", r.FindStringIndex("peach punch"))
```

---

## 🧷 Submatch ve İndeks Bilgileri

`Submatch` varyantları, hem tüm pattern eşleşmesi hem de bu eşleşme içindeki alt eşleşmeler (submatch) hakkında bilgi içerir. Örneğin bu, hem `p([a-z]+)ch` hem de `([a-z]+)` için bilgi döndürür.

```go
    fmt.Println(r.FindStringSubmatch("peach punch"))
```

Benzer şekilde bu, eşleşmelerin ve alt eşleşmelerin indeksleri hakkında bilgi döndürür.

```go
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
```

---

## 🔎 Tüm Eşleşmeleri Bulma

Bu fonksiyonların `All` varyantları, girdideki yalnızca ilk eşleşmeye değil, tüm eşleşmelere uygulanır. Örneğin bir regexp için tüm eşleşmeleri bulmak:

```go
    fmt.Println(r.FindAllString("peach punch pinch", -1))
```

Bu `All` varyantları, yukarıda gördüğümüz diğer fonksiyonlar için de mevcuttur.

```go
    fmt.Println("all:", r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))
```

Bu fonksiyonlara ikinci argüman olarak negatif olmayan bir tamsayı vermek, eşleşme sayısını sınırlar.

```go
    fmt.Println(r.FindAllString("peach punch pinch", 2))
```

---

## 🧱 []byte ile Çalışma

Yukarıdaki örnekler string argümanlarla ve `MatchString` gibi isimlerleydi. `[]byte` argümanlar da verebilir ve fonksiyon adından `String` kısmını kaldırabiliriz.

```go
    fmt.Println(r.Match([]byte("peach")))
```

---

## 🌍 Global Regexp Değişkenleri: MustCompile

Düzenli ifadelerle global değişkenler oluştururken `Compile` yerine `MustCompile` varyantını kullanabilirsiniz. `MustCompile`, hata döndürmek yerine `panic` eder; bu da global değişkenler için daha güvenli kullanım sağlar.

```go
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println("regexp:", r)
```

---

## 🔁 Metin Değiştirme

`regexp` paketi, string’in alt kümelerini başka değerlerle değiştirmek için de kullanılabilir.

```go
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
```

`Func` varyantı, eşleşen metni verilen bir fonksiyonla dönüştürmenizi sağlar.

```go
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run regular-expressions.go
true
true
peach
idx: [0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
regexp: p([a-z]+)ch
a <fruit>
a PEACH
```

---

## 📚 Referans

Go düzenli ifadeleri hakkında tam bir referans için `regexp` paketi dokümanlarına bakın.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: JSON.

