
## 🔁 Go by Example: Özyineleme (Recursion)

Go, *özyinelemeli (recursive)* fonksiyonları destekler. İşte klasik bir örnek.

### ▶️ Çalıştır

```go
package main
import "fmt"
```

Bu `fact` fonksiyonu, `fact(0)` *temel durumuna (base case)* ulaşana kadar kendisini çağırır.

```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
func main() {
    fmt.Println(fact(7))
```

## 🧩 Anonim Fonksiyonlarda Özyineleme

Anonim fonksiyonlar da özyinelemeli olabilir; ancak bunun için, fonksiyon tanımlanmadan önce onu saklamak üzere `var` ile bir değişkeni açıkça bildirmek gerekir.

```go
    var fib func(n int) int
    fib = func(n int) int {
        if n < 2 {
            return n
        }
```

`fib` daha önce `main` içinde bildirildiği için, Go burada `fib` ile hangi fonksiyonu çağıracağını bilir.

```go
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
```

## 💻 CLI Çıktısı

```bash
$ go run recursion.go 
5040
13
```

## ➡️ Sonraki Örnek

Next example: Range over Built-in Types.

