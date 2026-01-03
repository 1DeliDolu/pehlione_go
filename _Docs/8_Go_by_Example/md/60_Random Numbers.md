
## 🎲 Go by Example: Rastgele Sayılar

Go’nun `math/rand/v2` paketi, *sözde rastgele* sayı üretimi sağlar.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "math/rand/v2"
)

func main() {
    // Örneğin, rand.IntN 0 <= n < 100 olacak şekilde rastgele bir int n döndürür.
    fmt.Print(rand.IntN(100), ",")
    fmt.Print(rand.IntN(100))
    fmt.Println()

    // rand.Float64 0.0 <= f < 1.0 olacak şekilde bir float64 f döndürür.
    fmt.Println(rand.Float64())

    // Bu, diğer aralıklarda rastgele float üretmek için kullanılabilir; örneğin 5.0 <= f' < 10.0.
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // Bilinen bir seed istiyorsanız, yeni bir rand.Source oluşturup New kurucusuna verin.
    // NewPCG, iki adet uint64 sayıdan oluşan bir seed gerektiren yeni bir PCG kaynağı oluşturur.
    s2 := rand.NewPCG(42, 1024)
    r2 := rand.New(s2)
    fmt.Print(r2.IntN(100), ",")
    fmt.Print(r2.IntN(100))
    fmt.Println()

    s3 := rand.NewPCG(42, 1024)
    r3 := rand.New(s3)
    fmt.Print(r3.IntN(100), ",")
    fmt.Print(r3.IntN(100))
    fmt.Println()
}
```

### 🧾 Not

Örneği çalıştırdığınızda üretilen sayıların bazıları farklı olabilir.

### 💻 CLI

```bash
$ go run random-numbers.go
68,56
0.8090228139659177
5.840125017402497,6.937056298890035
94,49
94,49
```

### 📚 Ek Kaynak

Go’nun sağlayabildiği diğer rastgele nicelikler hakkında başvurular için `math/rand/v2` paket dokümanlarına bakın.

## ⏭️ Sonraki Örnek: Sayı Ayrıştırma

