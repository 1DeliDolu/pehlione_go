
## 🧠 Go by Example: Durum Sahibi Goroutine’ler (Stateful Goroutines)

Önceki örnekte, birden fazla *goroutine* arasında paylaşılan duruma erişimi senkronize etmek için *mutex*’lerle açık kilitleme (explicit locking) kullandık. Başka bir seçenek de, aynı sonucu elde etmek için *goroutine* ve *channel*’ların yerleşik senkronizasyon özelliklerini kullanmaktır. Bu kanal tabanlı yaklaşım, Go’nun “iletişim kurarak belleği paylaşma” ve her veri parçasının tam olarak 1 goroutine tarafından sahiplenilmesi (owned) fikriyle uyumludur.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)
```

---

## 🧩 Okuma/Yazma İstek Yapıları

Bu örnekte durum (state) tek bir goroutine tarafından sahiplenilecek. Bu, verinin eşzamanlı erişimle (concurrent access) asla bozulmamasını garanti eder. Bu durumu okumak veya yazmak için, diğer goroutine’ler sahip olan goroutine’e mesaj gönderecek ve karşılık gelen yanıtları alacaktır. Bu `readOp` ve `writeOp` struct’ları, bu istekleri ve sahip goroutine’in yanıt verebilmesi için bir yolu kapsüller (encapsulate).

```go
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}
func main() {
```

---

## 📊 Operasyon Sayacı

Daha önce olduğu gibi, kaç operasyon yaptığımızı sayacağız.

```go
    var readOps uint64
    var writeOps uint64
```

`reads` ve `writes` kanalları, diğer goroutine’ler tarafından sırasıyla okuma ve yazma istekleri göndermek için kullanılacak.

```go
    reads := make(chan readOp)
    writes := make(chan writeOp)
```

---

## 🏛️ State’i Sahiplenen Goroutine

İşte state’i sahiplenen goroutine; önceki örnekteki gibi bir `map` ama artık stateful goroutine’e özel (private). Bu goroutine, `reads` ve `writes` kanalları üzerinde tekrar tekrar `select` yapar ve istekler geldikçe yanıtlar. Yanıt, önce istenen operasyonu yapıp sonra başarıyı belirtmek için `resp` yanıt kanalına bir değer göndermekle gerçekleşir (okumalarda istenen değer gönderilir).

```go
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()
```

---

## 📥 100 Okuyucu Goroutine

Bu, `reads` kanalı aracılığıyla state sahibi goroutine’e okuma isteği gönderen 100 goroutine başlatır. Her okuma; bir `readOp` oluşturmayı, bunu `reads` kanalı üzerinden göndermeyi ve sonra sağlanan `resp` kanalından sonucu almayı gerektirir.

```go
    for range 100 {
        go func() {
            for {
                read := readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
```

---

## 📝 10 Yazıcı Goroutine

Benzer bir yaklaşımla 10 yazma işlemi de başlatırız.

```go
    for range 10 {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
```

---

## ⏱️ 1 Saniye Çalıştırma

Goroutine’lerin 1 saniye çalışmasına izin verin.

```go
    time.Sleep(time.Second)
```

---

## 📌 Sonuçları Alma ve Raporlama

Son olarak operasyon sayılarını yakalayın ve raporlayın.

```go
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}
```

---

## ✅ Çalıştırma Sonucu

Programımızı çalıştırmak, goroutine tabanlı state yönetimi örneğinin toplamda yaklaşık 80.000 operasyonu tamamladığını gösterir.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run stateful-goroutines.go
readOps: 71708
writeOps: 7177
```

---

## 🧭 Yaklaşım Karşılaştırması

Bu özel durumda goroutine tabanlı yaklaşım, mutex tabanlı olana göre biraz daha zahmetliydi. Yine de bazı durumlarda faydalı olabilir; örneğin başka kanalların da devrede olduğu yerlerde veya birden fazla mutex’i yönetmenin hata yapmaya açık olduğu durumlarda. Özellikle programınızın doğruluğunu (correctness) anlamak açısından hangi yaklaşım daha doğal geliyorsa onu kullanmalısınız.

---

## ⏭️ Sonraki Örnek

Next example: Sorting.

