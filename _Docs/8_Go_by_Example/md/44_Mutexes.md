
## 🔒 Go by Example: Mutex’ler (Mutexes)

Önceki örnekte, basit sayaç durumunu (counter state) atomik işlemlerle nasıl yöneteceğimizi gördük. Daha karmaşık durumlar için, birden fazla *goroutine* arasında veriye güvenli şekilde erişmek amacıyla bir *mutex* kullanabiliriz.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "sync"
)
```

---

## 🧰 Container Yapısı ve Mutex ile Senkronizasyon

`Container`, sayaçlardan oluşan bir `map` tutar; bunu birden fazla goroutine’den eşzamanlı olarak güncellemek istediğimiz için erişimi senkronize etmek adına bir `Mutex` ekleriz. *Mutex*’lerin kopyalanmaması gerektiğini unutmayın; bu nedenle bu struct etrafta taşınacaksa pointer ile taşınmalıdır.

```go
type Container struct {
    mu       sync.Mutex
    counters map[string]int
}
```

---

## 🔐 Kilitleme ve Kilit Açma

`counters`’a erişmeden önce mutex’i kilitleyin; fonksiyon sonunda `defer` ifadesiyle kilidi açın.

```go
func (c *Container) inc(name string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.counters[name]++
}
```

*Mutex*’in zero value’su olduğu gibi kullanılabildiği için, burada ayrıca bir başlangıç (initialization) gerekmez.

---

## 🧵 Goroutine’lerle Eşzamanlı Artırım

```go
func main() {
    c := Container{
        counters: map[string]int{"a": 0, "b": 0},
    }
    var wg sync.WaitGroup
```

Bu fonksiyon, isim verilmiş bir sayacı döngü içinde artırır.

```go
    doIncrement := func(name string, n int) {
        for range n {
            c.inc(name)
        }
    }
```

Birden fazla goroutine’i eşzamanlı çalıştırın; hepsinin aynı `Container`’a eriştiğine ve ikisinin aynı sayaca eriştiğine dikkat edin.

```go
    wg.Go(func() {
        doIncrement("a", 10000)
    })
    wg.Go(func() {
        doIncrement("a", 10000)
    })
    wg.Go(func() {
        doIncrement("b", 10000)
    })
```

Goroutine’lerin bitmesini bekleyin.

```go
    wg.Wait()
    fmt.Println(c.counters)
}
```

---

## ✅ Çalıştırma Sonucu

Programı çalıştırmak, sayaçların beklendiği gibi güncellendiğini gösterir.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run mutexes.go
map[a:20000 b:10000]
```

---

## 🔁 Sonraki Konu

Sonraki olarak, aynı state yönetimi işini yalnızca *goroutine*’ler ve *channel*’lar kullanarak uygulamaya bakacağız.

---

## ⏭️ Sonraki Örnek

Next example: Stateful Goroutines.

