
## 🧾 Go by Example: String Biçimlendirme (String Formatting)

Go, `printf` geleneğinde string biçimlendirme için mükemmel destek sunar. İşte yaygın string biçimlendirme görevlerine dair bazı örnekler.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "os"
)
type point struct {
    x, y int
}
func main() {
```

---

## 🧱 Genel Değerleri Biçimlendirme Verb’leri

Go, genel Go değerlerini biçimlendirmek için tasarlanmış çeşitli yazdırma “verb”’leri sunar. Örneğin, bu satır `point` struct’ımızın bir örneğini yazdırır.

```go
    p := point{1, 2}
    fmt.Printf("struct1: %v\n", p)
```

Değer bir struct ise, `%+v` varyantı struct alan adlarını da içerir.

```go
    fmt.Printf("struct2: %+v\n", p)
```

`%#v` varyantı, değerin bir Go sözdizimi (Go syntax) temsiliğini yazdırır; yani o değeri üretecek kaynak kod parçasını.

```go
    fmt.Printf("struct3: %#v\n", p)
```

Bir değerin tipini yazdırmak için `%T` kullanın.

```go
    fmt.Printf("type: %T\n", p)
```

---

## ✅ Boolean Biçimlendirme

Boolean biçimlendirmek düz bir işlemdir.

```go
    fmt.Printf("bool: %t\n", true)
```

---

## 🔢 Tamsayı Biçimlendirme

Tamsayıları biçimlendirmek için birçok seçenek vardır. Standart, 10’luk taban biçimlendirme için `%d` kullanın.

```go
    fmt.Printf("int: %d\n", 123)
```

Bu, ikili (binary) gösterimi yazdırır.

```go
    fmt.Printf("bin: %b\n", 14)
```

Bu, verilen tamsayıya karşılık gelen karakteri yazdırır.

```go
    fmt.Printf("char: %c\n", 33)
```

`%x` hex kodlaması sağlar.

```go
    fmt.Printf("hex: %x\n", 456)
```

---

## 🌊 Float Biçimlendirme

Float’lar için de çeşitli biçimlendirme seçenekleri vardır. Temel ondalık biçimlendirme için `%f` kullanın.

```go
    fmt.Printf("float1: %f\n", 78.9)
```

`%e` ve `%E`, float’ı (biraz farklı versiyonlarda) bilimsel gösterimde biçimlendirir.

```go
    fmt.Printf("float2: %e\n", 123400000.0)
    fmt.Printf("float3: %E\n", 123400000.0)
```

---

## 🔤 String Biçimlendirme

Temel string yazdırma için `%s` kullanın.

```go
    fmt.Printf("str1: %s\n", "\"string\"")
```

Go kaynak kodundaki gibi çift tırnakla (double-quote) string yazdırmak için `%q` kullanın.

```go
    fmt.Printf("str2: %q\n", "\"string\"")
```

Daha önce tamsayılarda gördüğümüz gibi, `%x` string’i 16’lık tabanda gösterir; girişin her baytı için iki çıktı karakteri üretir.

```go
    fmt.Printf("str3: %x\n", "hex this")
```

---

## 📍 Pointer Biçimlendirme

Bir pointer temsili yazdırmak için `%p` kullanın.

```go
    fmt.Printf("pointer: %p\n", &p)
```

---

## 📐 Genişlik ve Hassasiyet (Width & Precision)

Sayıları biçimlendirirken genellikle ortaya çıkan değerin genişliğini ve hassasiyetini kontrol etmek istersiniz. Bir tamsayının genişliğini belirtmek için verb içindeki `%`’den sonra bir sayı kullanın. Varsayılan olarak sonuç sağa yaslanır ve boşluklarla doldurulur.

```go
    fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
```

Basılmış float’ların genişliğini de belirtebilirsiniz; ancak genellikle `width.precision` sözdizimi ile aynı anda ondalık hassasiyeti de sınırlandırmak istersiniz.

```go
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)
```

Sola yaslamak (left-justify) için `-` bayrağını kullanın.

```go
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)
```

Özellikle tablo benzeri çıktılarda hizalamayı garantilemek için string biçimlendirirken de genişliği kontrol etmek isteyebilirsiniz. Temel sağa yaslı genişlik:

```go
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")
```

Sola yaslamak için, sayılarda olduğu gibi `-` bayrağını kullanın.

```go
    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")
```

---

## 🧵 Sprintf: Yazdırmadan String Üretme

Şimdiye kadar biçimlendirilmiş string’i `os.Stdout`’a yazdıran `Printf`’i gördük. `Sprintf` ise biçimlendirir ve hiçbir yere yazdırmadan bir string döndürür.

```go
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)
```

---

## 🧾 Fprintf: os.Stdout Dışındaki Writer’lara Yazma

`os.Stdout` dışındaki `io.Writer`’lara biçimlendirip yazdırmak için `Fprintf` kullanabilirsiniz.

```go
    fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
```

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run string-formatting.go
struct1: {1 2}
struct2: {x:1 y:2}
struct3: main.point{x:1, y:2}
type: main.point
bool: true
int: 123
bin: 1110
char: !
hex: 1c8
float1: 78.900000
float2: 1.234000e+08
float3: 1.234000E+08
str1: "string"
str2: "\"string\""
str3: 6865782074686973
pointer: 0xc0000ba000
width1: |    12|   345|
width2: |  1.20|  3.45|
width3: |1.20  |3.45  |
width4: |   foo|     b|
width5: |foo   |b     |
sprintf: a string
io: an error
```

---

## ⏭️ Sonraki Örnek

Next example: Text Templates.

