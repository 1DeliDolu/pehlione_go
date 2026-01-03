
## ⚛️ Go by Example: Atomik Sayaçlar (Atomic Counters)

Go’da durumu (state) yönetmenin birincil mekanizması, *channel*’lar üzerinden iletişim kurmaktır. Bunu örneğin *worker pool*’larda gördük. Ancak state yönetimi için birkaç seçenek daha vardır. Burada, birden çok *goroutine* tarafından erişilen atomik sayaçlar için `sync/atomic` paketini kullanmaya bakacağız.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)
func main() {
```

---

## 🔢 Atomik Sayaç Tanımı

(her zaman pozitif olan) sayacımızı temsil etmek için atomik bir tamsayı tipi kullanacağız.

```go
    var ops atomic.Uint64
```

---

## 🧵 Goroutine’leri Beklemek: WaitGroup

Bir `WaitGroup`, tüm goroutine’lerin işini bitirmesini beklememize yardımcı olacak.

```go
    var wg sync.WaitGroup
```

---

## 🚀 50 Goroutine ile Sayaç Artırma

Her biri sayacı tam olarak 1000 kez artıran 50 goroutine başlatacağız.

```go
    for range 50 {
        wg.Go(func() {
            for range 1000 {
```

Sayacı atomik olarak artırmak için `Add` kullanırız.

```go
                ops.Add(1)
            }
        })
    }
```

---

## ✅ Tamamlanmayı Bekleme

Tüm goroutine’ler bitene kadar bekleyin.

```go
    wg.Wait()
```

Burada hiçbir goroutine `ops`’a yazmıyor; ancak `Load` kullanarak, diğer goroutine’ler (atomik olarak) güncelleme yapıyor olsa bile bir değeri atomik şekilde okumak güvenlidir.

```go
    fmt.Println("ops:", ops.Load())
}
```

---

## 🎯 Beklenen Sonuç

Tam olarak `50,000` işlem bekliyoruz. Atomik olmayan bir tamsayı kullanıp `ops++` ile artırmış olsaydık, büyük olasılıkla her çalıştırmada değişen farklı bir sayı elde ederdik; çünkü goroutine’ler birbirinin işine karışırdı. Dahası, `-race` bayrağıyla çalıştırırken *data race* hataları alırdık.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run atomic-counters.go
ops: 50000
```

---

## 🔒 Sonraki Konu

Sonraki olarak, state yönetimi için başka bir araç olan *mutex*’lere bakacağız.

---

## ⏭️ Sonraki Örnek

Next example: Mutexes.

