
## 🧷 Go by Example: Kapanışlar (Closures)

Go, *closure* oluşturabilen anonim fonksiyonları destekler. Anonim fonksiyonlar, bir fonksiyonu isim vermek zorunda kalmadan satır içinde tanımlamak istediğinizde kullanışlıdır.

```go
package main

import "fmt"
```

Bu `intSeq` fonksiyonu başka bir fonksiyon döndürür; bunu `intSeq` gövdesinde anonim olarak tanımlarız. Döndürülen fonksiyon, bir *closure* oluşturmak için `i` değişkenini kapsar.

```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

```go
func main() {
```

`intSeq`’i çağırıp sonucu (bir fonksiyon) `nextInt`’e atarız. Bu fonksiyon değeri, her `nextInt` çağrısında güncellenecek kendi `i` değerini yakalar.

```go
    nextInt := intSeq()
```

`nextInt`’i birkaç kez çağırarak *closure*’ın etkisini görün.

```go
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
```

Durumun (state) yalnızca o belirli fonksiyona özgü olduğunu doğrulamak için yenisini oluşturup test edin.

```go
    newInts := intSeq()
    fmt.Println(newInts())
}
```

```bash
$ go run closures.go
1
2
3
1
```

Şimdilik fonksiyonların bakacağımız son özelliği özyineleme (recursion).

Sonraki örnek: *Recursion*.

