
## 🧷 Go by Example: Constants

Go; karakter, string, boolean ve sayısal değerler için sabitleri (*constants*) destekler.

---

## 🧾 Kaynak Kod

```go
package main

import (
    "fmt"
    "math"
)

// const, bir sabit değer bildirir.
const s string = "constant"

func main() {
    fmt.Println(s)

    // const deyimi bir fonksiyon gövdesi içinde de yer alabilir.
    const n = 500000000

    // Sabit ifadeler, keyfi hassasiyetle (*arbitrary precision*) aritmetik yapar.
    const d = 3e20 / n
    fmt.Println(d)

    // Sayısal bir sabitin, açık bir dönüşüm gibi bir yolla tür verilene kadar bir türü yoktur.
    fmt.Println(int64(d))

    // Bir sayı, tür gerektiren bir bağlamda kullanılarak tür alabilir; ör. değişken ataması veya fonksiyon çağrısı.
    // Örneğin burada math.Sin bir float64 bekler.
    fmt.Println(math.Sin(n))
}
```

---

## ▶️ Çalıştırma

```bash
$ go run constant.go 
constant
6e+11
600000000000
-0.28470407323754404
```

