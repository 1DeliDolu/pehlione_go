
## 🔢 Go by Example: Sayı Ayrıştırma

String’lerden sayı ayrıştırmak, birçok programda temel ama yaygın bir iştir; Go’da bunu nasıl yapacağınız aşağıdadır.

### ▶️ Çalıştırma

```go
package main

// Yerleşik strconv paketi sayı ayrıştırma işlemlerini sağlar.
import (
    "fmt"
    "strconv"
)

func main() {
    // ParseFloat ile, buradaki 64 kaç bitlik hassasiyetle ayrıştırılacağını belirtir.
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // ParseInt için 0, tabanın string’den çıkarılacağı (infer) anlamına gelir.
    // 64 ise sonucun 64 bit içine sığmasını gerektirir.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // ParseInt, hex formatındaki sayıları tanır.
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // ParseUint de mevcuttur.
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // Atoi, temel base-10 int ayrıştırma için bir kolaylık fonksiyonudur.
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // Parse fonksiyonları hatalı girdi için error döndürür.
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
```

### 💻 CLI

```bash
$ go run number-parsing.go 
1.234
123
456
789
135
strconv.ParseInt: parsing "wat": invalid syntax
```

### 🧩 Sonraki Adım

Sırada başka bir yaygın ayrıştırma görevi var: URL’ler.

## ⏭️ Sonraki Örnek: URL Ayrıştırma

