
## 🧮 Go by Example: Değişken Argümanlı Fonksiyonlar (Variadic Functions)

Değişken argümanlı (variadic) fonksiyonlar, sonda herhangi bir sayıda argümanla çağrılabilir. Örneğin `fmt.Println`, yaygın bir variadic fonksiyondur.

```go
package main

import "fmt"
```

İşte argüman olarak keyfî sayıda `int` alacak bir fonksiyon.

```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
```

Fonksiyon içinde `nums`’un türü `[]int` ile eşdeğerdir. `len(nums)` çağırabilir, `range` ile üzerinde dolaşabilir, vb.

```go
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

```go
func main() {
```

Variadic fonksiyonlar, tek tek argümanlarla normal şekilde çağrılabilir.

```go
    sum(1, 2)
    sum(1, 2, 3)
```

Eğer elinizde zaten bir *slice* içinde birden fazla argüman varsa, bunları `func(slice...)` şeklinde variadic fonksiyona uygulayın.

```go
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```

```bash
$ go run variadic-functions.go 
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

Go’da fonksiyonların bir diğer önemli yönü de *closure* oluşturabilmeleridir; bunu sıradaki örnekte inceleyeceğiz.

Sonraki örnek: *Closures*.

