
## ⚠️ Go by Example: Hatalar

Go’da hataları iletmenin *idiomatik* yolu, açık ve ayrı bir dönüş değeri ile bunu belirtmektir. Bu yaklaşım; Java, Python ve Ruby gibi dillerde kullanılan *exception* yapılarından ve C’de bazen kullanılan “tek bir sonuç/hata değeri” şeklindeki aşırı yüklenmiş yaklaşımdan farklıdır. Go’nun yaklaşımı, hangi fonksiyonların hata döndürdüğünü kolayca görmeyi ve hataları, hata dışı işlerde de kullanılan aynı dil yapılarıyla ele almayı kolaylaştırır.

Ek ayrıntılar için `errors` paketinin dokümantasyonuna ve ilgili blog yazısına bakın.

---

## ▶️ Çalıştırma

```go
package main
import (
    "errors"
    "fmt"
)
```

---

## 🧾 Hatalar Ayrı Bir Return Değeri Olarak Döndürülür

Geleneksel olarak hatalar, son dönüş değeri olur ve yerleşik bir arayüz olan `error` tipindedir.

```go
func f(arg int) (int, error) {
    if arg == 42 {
```

`errors.New`, verilen hata mesajıyla temel bir hata değeri oluşturur.

```go
        return -1, errors.New("can't work with 42")
    }
```

`error` pozisyonunda `nil` olması, hata olmadığı anlamına gelir.

```go
    return arg + 3, nil
}
```

---

## 🧷 Sentinel Error Tanımı

*Sentinel error*, belirli bir hata durumunu ifade etmek için kullanılan, önceden tanımlanmış bir değişkendir.

```go
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")
func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {
```

---

## 🧩 Hata Sarma (Wrapping) ile Bağlam Eklemek

Daha üst seviye hatalarla, bağlam eklemek için hataları *wrap* edebiliriz. Bunun en basit yolu `fmt.Errorf` içinde `%w` fiilini kullanmaktır. Wrap edilmiş hatalar, (A, B’yi sarar; B, C’yi sarar vb.) mantıksal bir zincir oluşturur ve bu zincir `errors.Is` ve `errors.As` gibi fonksiyonlarla sorgulanabilir.

```go
        return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}
```

---

## 🧪 Hata Kontrolü ve İşleme

```go
func main() {
    for _, i := range []int{7, 42} {
```

`if` satırında *inline* hata kontrolü yapmak idiomatiktir.

```go
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }
```

```go
    for i := range 5 {
        if err := makeTea(i); err != nil {
```

`errors.Is`, verilen bir hatanın (veya hata zincirindeki herhangi bir hatanın) belirli bir hata değeriyle eşleşip eşleşmediğini kontrol eder. Bu, özellikle wrap edilmiş ya da iç içe geçmiş hatalarda faydalıdır; zincir içinde belirli hata tiplerini veya sentinel hataları tanımlamanızı sağlar.

```go
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }
        fmt.Println("Tea is ready!")
    }
}
```

---

## 🖥️ Çalıştırma Komutu ve Çıktı

```bash
$ go run errors.go
```

```text
f worked: 10
f failed: can't work with 42
Tea is ready!
Tea is ready!
We should buy new tea!
Tea is ready!
Now it is dark.
```

