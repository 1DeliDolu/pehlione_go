
## 🧑‍🏭 Go by Example: İşçi Havuzları (Worker Pools)

Bu örnekte *goroutine*’ler ve *channel*’lar kullanarak bir *worker pool* (işçi havuzu) nasıl uygulanır, buna bakacağız.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "time"
)
```

---

## 👷 Worker Fonksiyonu

İşte worker (işçi) fonksiyonu; bunun birkaç eşzamanlı (concurrent) örneğini çalıştıracağız. Bu worker’lar `jobs` kanalından iş alacak ve karşılık gelen sonuçları `results` kanalına gönderecek. Pahalı bir işi simüle etmek için iş başına 1 saniye uyuyacağız.

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}
func main() {
```

---

## 📮 İş Gönderme ve Sonuç Toplama

Worker havuzumuzu kullanmak için onlara iş göndermemiz ve sonuçlarını toplamamız gerekir. Bunun için 2 kanal oluşturuyoruz.

```go
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
```

Bu, 3 worker başlatır; henüz iş olmadığı için başlangıçta bloklanırlar.

```go
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
```

Burada 5 işi gönderiyoruz ve ardından elimizdeki tüm işin bu olduğunu belirtmek için bu kanalı kapatıyoruz.

```go
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
```

Son olarak tüm işlerin sonuçlarını topluyoruz. Bu aynı zamanda worker goroutine’lerinin bittiğinden emin olmamızı sağlar. Birden çok goroutine’i beklemenin alternatif yolu `WaitGroup` kullanmaktır.

```go
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

---

## ⏱️ Çalışma Süresi ve Eşzamanlılık

Çalışan programımız, 5 işin çeşitli worker’lar tarafından yürütüldüğünü gösterir. Program, toplamda yaklaşık 5 saniyelik iş yapmasına rağmen yalnızca yaklaşık 2 saniye sürer; çünkü 3 worker eşzamanlı çalışır.

---

## 💻 CLI Çalıştırma ve Örnek Çıktı

```bash
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
real    0m2.358s
```

---

## ⏭️ Sonraki Örnek

Next example: WaitGroups.

