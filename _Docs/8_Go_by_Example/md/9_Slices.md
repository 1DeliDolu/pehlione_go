
## 🧩 Go by Example: Dilimler (Slices)

*Slice*’lar Go’da önemli bir veri türüdür; dizilere (arrays) göre ardışık öğeler üzerinde daha güçlü bir arayüz sağlar.

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
```

Dizilerden farklı olarak *slice*’lar yalnızca içerdikleri öğe türüne göre tiplenir (öğe sayısına göre değil). Başlatılmamış bir *slice*, `nil`’e eşittir ve uzunluğu `0`’dır.

```go
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)
```

Sıfır olmayan uzunlukta bir *slice* oluşturmak için yerleşik `make` kullanılır. Burada uzunluğu `3` olan (başlangıçta *zero-valued*) bir `string` dilimi oluşturuyoruz. Varsayılan olarak yeni bir *slice*’ın kapasitesi, uzunluğuna eşittir; *slice*’ın ileride büyüyeceğini önceden biliyorsak, `make`’e ek bir parametre olarak kapasiteyi açıkça geçirmek mümkündür.

```go
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
```

Dizilerde olduğu gibi değer atayıp okuyabiliriz.

```go
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
```

Beklendiği gibi `len`, *slice*’ın uzunluğunu döndürür.

```go
    fmt.Println("len:", len(s))
```

Bu temel işlemlere ek olarak *slice*’lar, onları dizilerden daha zengin yapan birkaç özelliği daha destekler. Bunlardan biri yerleşik `append`’dir; bir veya daha fazla yeni değer içeren bir *slice* döndürür. `append` sonucunu mutlaka bir dönüş değeri olarak almak gerekir, çünkü yeni bir *slice* değeri dönebilir.

```go
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
```

*Slice*’lar `copy` ile kopyalanabilir. Burada `s` ile aynı uzunlukta boş bir `c` *slice*’ı oluşturup `s` içeriğini `c`’ye kopyalıyoruz.

```go
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
```

*Slice*’lar `slice[low:high]` söz dizimine sahip bir “dilimleme” operatörünü destekler. Örneğin bu, `s[2]`, `s[3]` ve `s[4]` öğelerini alan bir *slice* elde eder.

```go
    l := s[2:5]
    fmt.Println("sl1:", l)
```

Bu, `s[5]`’e kadar (ama `s[5]` hariç) dilimler.

```go
    l = s[:5]
    fmt.Println("sl2:", l)
```

Ve bu da `s[2]`’den başlayarak (dahil) sonuna kadar dilimler.

```go
    l = s[2:]
    fmt.Println("sl3:", l)
```

Bir *slice* değişkenini tek satırda tanımlayıp başlatabiliriz.

```go
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
```

`slices` paketi, *slice*’lar için bir dizi faydalı yardımcı fonksiyon içerir.

```go
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    }
```

*Slice*’lar çok boyutlu veri yapıları oluşturmak için birleştirilebilir. Çok boyutlu dizilerin (arrays) aksine, iç *slice*’ların uzunluğu değişken olabilir.

```go
    twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

Dikkat: *slice*’lar dizilerden farklı türler olsa da, `fmt.Println` ile benzer şekilde gösterilirler.

```bash
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]
```

Go ekibinin, Go’da *slice*’ların tasarımı ve uygulanışı hakkında daha fazla detay içeren bu harika blog yazısına da göz atın.

Artık dizileri (arrays) ve *slice*’ları gördüğümüze göre, Go’nun diğer önemli yerleşik veri yapısına bakacağız: *map*’ler.

Sonraki örnek: *Maps*.

