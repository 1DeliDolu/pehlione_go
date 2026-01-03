
## 🔁 Channel’lar Üzerinde Range

Önceki bir örnekte `for` ve `range` ile temel veri yapılarını nasıl dolaştığımızı görmüştük. Aynı sözdizimini, bir channel’dan alınan değerler üzerinde iterasyon yapmak için de kullanabiliriz.

---

## ▶️ Çalıştır

```go
package main

import "fmt"
```

---

## 🧩 main Fonksiyonu

`queue` channel’ında kuyrukta olan 2 değer üzerinde iterasyon yapacağız.

```go
func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
```

Bu `range`, `queue` üzerinden alınan her bir eleman geldikçe iterasyon yapar. Yukarıda channel’ı kapattığımız için, iterasyon 2 elemanı aldıktan sonra sonlanır.

```go
    for elem := range queue {
        fmt.Println(elem)
    }
}
```

---

## 🖨️ CLI Çıktısı

```bash
$ go run range-over-channels.go
one
two
```

---

## 📝 Not

Bu örnek ayrıca, boş olmayan bir channel’ın kapatılmasının mümkün olduğunu ve yine de kalan değerlerin alınabildiğini gösterdi.

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Timers*.

