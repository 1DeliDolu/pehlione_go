
## 🔒 Channel Kapatma

Bir channel’ı kapatmak, artık o channel üzerinden başka değer gönderilmeyeceğini belirtir. Bu, channel alıcılarına tamamlanma bilgisini iletmek için faydalı olabilir.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 🧩 Örnek Senaryo

Bu örnekte, `main()` goroutine’inden bir worker goroutine’ine yapılacak işleri iletmek için bir `jobs` channel’ı kullanacağız. Worker için artık işimiz kalmadığında `jobs` channel’ını kapatacağız.

```go
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)
```

---

## 🧑‍🏭 Worker Goroutine

İşte worker goroutine’i. Sürekli olarak `jobs` channel’ından `j, more := <-jobs` ile okur. Bu özel 2-değerli alma biçiminde, `jobs` kapatılmışsa ve channel’daki tüm değerler zaten alınmışsa `more` değeri `false` olur. Bunu, tüm işlerimizi bitirdiğimizde `done` üzerinden bildirim yapmak için kullanırız.

```go
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
```

---

## 📤 İş Gönderme ve Channel’ı Kapatma

Bu kısım, `jobs` channel’ı üzerinden worker’a 3 iş gönderir, ardından channel’ı kapatır.

```go
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
```

Daha önce gördüğümüz senkronizasyon yaklaşımını kullanarak worker’ı bekleriz.

```go
    <-done
```

---

## 📥 Kapalı Channel’dan Okuma

Kapalı bir channel’dan okumak hemen başarılı olur ve ilgili tipin *zero value* değerini döndürür. Opsiyonel ikinci dönüş değeri, alınan değerin channel’a başarılı bir gönderim (*send*) işlemiyle gelip gelmediğini gösterir: Eğer değer başarılı bir gönderimle geldiyse `true`, channel kapalı ve boş olduğu için üretilmiş bir zero value ise `false` olur.

```go
    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
received more jobs: false
```

---

## 🔁 Sonraki Örneğe Geçiş

Kapalı channel fikri, bir sonraki örneğimize doğal olarak götürür: channel’lar üzerinde `range` kullanımı.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Range over Channels*.

