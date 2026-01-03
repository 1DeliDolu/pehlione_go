
## 🔁 Go by Example: For

`for`, Go’nun tek döngü yapısıdır. İşte bazı temel `for` döngüsü türleri.

---

## 🧾 Kaynak Kod

```go
package main

import "fmt"

func main() {
    // En temel tür: tek bir koşul ile.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // Klasik başlangıç/koşul/sonrası (init/condition/after) for döngüsü.
    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    // Temel “bunu N kez yap” yinelemesini yapmanın başka bir yolu: bir tamsayı üzerinde range kullanmak.
    for i := range 3 {
        fmt.Println("range", i)
    }

    // Koşulsuz for, döngüden break ile çıkana ya da saran fonksiyondan return edene kadar tekrar eder.
    for {
        fmt.Println("loop")
        break
    }

    // Bir sonraki yinelemeye geçmek için continue da kullanabilirsiniz.
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

---

## ▶️ Çalıştırma

```bash
$ go run for.go
1
2
3
0
1
2
range 0
range 1
range 2
loop
1
3
5
```

---

## 📌 Not

`range` deyimlerine, kanallara (*channels*) ve diğer veri yapılarına baktığımızda daha sonra başka `for` biçimleri de göreceğiz.

**Sonraki örnek:** *If/Else.*

