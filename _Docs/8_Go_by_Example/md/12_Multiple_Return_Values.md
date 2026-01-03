
## 🔁 Go by Example: Çoklu Dönüş Değerleri (Multiple Return Values)

Go, birden fazla dönüş değerini yerleşik olarak destekler. Bu özellik, deyimsel (idiomatic) Go’da sıkça kullanılır; örneğin bir fonksiyondan hem sonuç hem de hata (*error*) değerlerini döndürmek için.

```go
package main

import "fmt"
```

Bu fonksiyon imzasındaki `(int, int)`, fonksiyonun **2 adet `int`** döndürdüğünü gösterir.

```go
func vals() (int, int) {
    return 3, 7
}
```

```go
func main() {
```

Burada çağrıdan gelen 2 farklı dönüş değerini çoklu atama (multiple assignment) ile kullanıyoruz.

```go
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
```

Eğer dönen değerlerin yalnızca bir kısmını istiyorsanız, boş tanımlayıcı `_` kullanın.

```go
    _, c := vals()
    fmt.Println(c)
}
```

```bash
$ go run multiple-return-values.go
3
7
7
```

Değişken sayıda argüman kabul etmek, Go fonksiyonlarının bir başka güzel özelliğidir; bunu sıradaki örnekte inceleyeceğiz.

Sonraki örnek: *Variadic Functions*.

