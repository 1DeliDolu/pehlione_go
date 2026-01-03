
## 🧠 Go by Example: Fonksiyonlar (Functions)

Fonksiyonlar Go’da merkezî bir yer tutar. Birkaç farklı örnekle fonksiyonları öğreneceğiz.

```go
package main

import "fmt"
```

İşte iki `int` alan ve toplamlarını `int` olarak döndüren bir fonksiyon.

```go
func plus(a int, b int) int {
```

Go, açık (explicit) `return` zorunlu kılar; yani son ifadenin değerini otomatik olarak döndürmez.

```go
    return a + b
}
```

Aynı türden birden fazla ardışık parametre olduğunda, son parametre türü belirtene kadar, aynı türdeki parametreler için tür adını yazmayabilirsiniz.

```go
func plusPlus(a, b, c int) int {
    return a + b + c
}
```

```go
func main() {
```

Bir fonksiyonu beklediğiniz gibi `name(args)` ile çağırırsınız.

```go
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
```

```bash
$ go run functions.go 
1+2 = 3
1+2+3 = 6
```

Go fonksiyonlarında başka birkaç özellik daha vardır. Bunlardan biri, birden fazla dönüş değeridir; bunu sıradaki örnekte göreceğiz.

Sonraki örnek: *Multiple Return Values*.

