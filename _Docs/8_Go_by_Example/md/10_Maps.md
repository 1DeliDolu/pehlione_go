
## 🗺️ Go by Example: Haritalar (Maps)

*Map*’ler, Go’nun yerleşik ilişkisel (associative) veri türüdür (diğer dillerde bazen *hash* veya *dict* olarak da adlandırılır).

```go
package main

import (
    "fmt"
    "maps"
)

func main() {
```

Boş bir *map* oluşturmak için yerleşik `make` kullanın: `make(map[key-type]val-type)`.

```go
    m := make(map[string]int)
```

Anahtar/değer çiftlerini tipik `name[key] = val` söz dizimiyle ayarlayın.

```go
    m["k1"] = 7
    m["k2"] = 13
```

Bir *map*’i `fmt.Println` gibi bir şeyle yazdırmak, tüm anahtar/değer çiftlerini gösterir.

```go
    fmt.Println("map:", m)
```

Bir anahtarın değerini `name[key]` ile alın.

```go
    v1 := m["k1"]
    fmt.Println("v1:", v1)
```

Anahtar yoksa, değer türünün *sıfır değeri (zero value)* döndürülür.

```go
    v3 := m["k3"]
    fmt.Println("v3:", v3)
```

Yerleşik `len`, bir *map* üzerinde çağrıldığında anahtar/değer çiftlerinin sayısını döndürür.

```go
    fmt.Println("len:", len(m))
```

Yerleşik `delete`, bir *map*’ten anahtar/değer çiftlerini kaldırır.

```go
    delete(m, "k2")
    fmt.Println("map:", m)
```

Bir *map*’ten tüm anahtar/değer çiftlerini kaldırmak için yerleşik `clear` kullanın.

```go
    clear(m)
    fmt.Println("map:", m)
```

Bir *map*’ten değer alırken isteğe bağlı ikinci dönüş değeri, anahtarın *map*’te bulunup bulunmadığını belirtir. Bu, eksik anahtarlar ile `0` veya `""` gibi sıfır değerli anahtarları birbirinden ayırmak için kullanılabilir. Burada değerin kendisine ihtiyaç duymadık, bu yüzden boş tanımlayıcı `_` ile yok saydık.

```go
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
```

Bu söz dizimiyle aynı satırda yeni bir *map* de tanımlayıp başlatabilirsiniz.

```go
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
```

`maps` paketi, *map*’ler için bir dizi faydalı yardımcı fonksiyon içerir.

```go
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}
```

Dikkat: *map*’ler `fmt.Println` ile yazdırıldığında `map[k:v k:v]` biçiminde görünür.

```bash
$ go run maps.go 
map: map[k1:7 k2:13]
v1: 7
v3: 0
len: 2
map: map[k1:7]
map: map[]
prs: false
map: map[bar:2 foo:1]
n == n2
```

Sonraki örnek: *Functions*.

