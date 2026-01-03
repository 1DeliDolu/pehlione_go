
## ⏱️ Timeout’lar

Timeout’lar, dış kaynaklara bağlanan veya yürütme süresini sınırlaması gereken programlar için önemlidir. Go’da timeout uygulamak, channel’lar ve `select` sayesinde kolay ve zariftir.

---

## ▶️ Çalıştır

```go
package main

import (
    "fmt"
    "time"
)
```

---

## 🧩 main Fonksiyonu

Örneğimizde, dış bir çağrının sonucu 2 saniye sonra `c1` channel’ı üzerinden döndürdüğünü varsayalım. Channel’ın buffer’lı olduğuna dikkat edin; böylece goroutine içindeki gönderim (*send*) bloklamaz. Bu, channel hiç okunmazsa goroutine sızıntılarını (*goroutine leaks*) önlemek için yaygın bir desendir.

```go
func main() {
    c1 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result 1"
    }()
```

İşte timeout’u uygulayan `select`. `res := <-c1` sonuç bekler; `<-time.After` ise 1 saniyelik timeout sonrasında gönderilecek bir değeri bekler. `select`, hazır olan ilk alma işlemiyle devam ettiği için, işlem izin verilen 1 saniyeden uzun sürerse timeout durumunu seçeriz.

```go
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }
```

Eğer 3 saniyelik daha uzun bir timeout’a izin verirsek, `c2` üzerinden alma işlemi başarılı olur ve sonucu yazdırırız.

```go
    c2 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
```

---

## 🖨️ Program Çıktısı

Bu programı çalıştırdığımızda, ilk işlemin timeout’a düştüğünü ve ikincisinin başarılı olduğunu görürüz.

```bash
$ go run timeouts.go
timeout 1
result 2
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Non-Blocking Channel Operations*.

