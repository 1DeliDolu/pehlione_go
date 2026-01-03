
## 🧭 Go by Example: Yerleşik Türler Üzerinde Range

`range`, çeşitli yerleşik veri yapılarındaki elemanlar üzerinde yineleme yapar. Daha önce öğrendiğimiz veri yapılarının bazılarıyla `range` kullanımına bakalım.

### ▶️ Çalıştır

```go
package main
import "fmt"
func main() {
```

## ➕ Slice Üzerinde Toplama

Burada bir *slice* içindeki sayıları toplamak için `range` kullanıyoruz. *Array*’ler de aynı şekilde çalışır.

```go
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
```

## 🔢 Index ve Değer

Array’ler ve slice’lar üzerinde `range`, her giriş için hem *index* hem de *değer* sağlar. Yukarıda index’e ihtiyacımız olmadığı için, boş tanımlayıcı `_` ile onu yok saydık. Ancak bazen index’leri gerçekten isteriz.

```go
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
```

## 🗺️ Map Üzerinde Key/Value

`map` üzerinde `range`, *key/value* çiftleri üzerinde yineleme yapar.

```go
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
```

## 🔑 Sadece Key’ler Üzerinde Dönme

`range`, bir `map`’in sadece *key*’leri üzerinde de yineleme yapabilir.

```go
    for k := range kvs {
        fmt.Println("key:", k)
    }
```

## 🔤 String Üzerinde Unicode Rune’ları

String’ler üzerinde `range`, Unicode *code point*’leri üzerinde yineleme yapar. İlk değer, *rune*’un başladığı *byte index*’idir; ikinci değer ise *rune*’un kendisidir. Daha fazla ayrıntı için *Strings and Runes* bölümüne bakın.

```go
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

## 💻 CLI Çıktısı

```bash
$ go run range-over-built-in-types.go
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
```

## ➡️ Sonraki Örnek

Next example: Pointers.

