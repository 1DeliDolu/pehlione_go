
## 🧱 Go by Example: Diziler (Arrays)

Go’da bir *dizi (array)*, belirli bir uzunluğa sahip, numaralandırılmış bir öğe dizisidir. Tipik Go kodunda *slice*’lar çok daha yaygındır; diziler ise bazı özel senaryolarda faydalıdır.

```go
package main

import "fmt"

func main() {
```

Burada tam olarak **5 adet int** tutacak bir `a` dizisi oluşturuyoruz. Öğelerin türü ve uzunluk, dizinin türünün bir parçasıdır. Varsayılan olarak bir dizi *sıfır değerli (zero-valued)* olur; `int` için bu, `0` değerleridir.

```go
    var a [5]int
    fmt.Println("emp:", a)
```

Bir indeksteki değeri `array[index] = value` söz dizimiyle ayarlayabilir, `array[index]` ile alabilirsiniz.

```go
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
```

Yerleşik `len`, bir dizinin uzunluğunu döndürür.

```go
    fmt.Println("len:", len(a))
```

Bu söz dizimiyle bir diziyi tek satırda tanımlayıp başlatabilirsiniz.

```go
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
```

Ayrıca derleyicinin öğe sayısını `...` ile sizin yerinize saymasını sağlayabilirsiniz.

```go
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
```

İndeksi `:` ile belirtirseniz, aradaki öğeler sıfırlanır.

```go
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)
```

Dizi türleri tek boyutludur, ancak türleri birleştirerek çok boyutlu veri yapıları oluşturabilirsiniz.

```go
    var twoD [2][3]int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
```

Çok boyutlu dizileri tek seferde oluşturup başlatabilirsiniz.

```go
    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}
```

Dikkat: Diziler `fmt.Println` ile yazdırıldığında `[v1 v2 v3 ...]` biçiminde görünür.

```bash
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
dcl: [1 2 3 4 5]
idx: [100 0 0 400 500]
2d:  [[0 1 2] [1 2 3]]
2d:  [[1 2 3] [1 2 3]]
```

Sonraki örnek: *Slices*.

