
## 🧩 Özel Hatalar

`Error()` metodunu bir tip üzerinde uygulayarak özel hata (*custom error*) tipleri tanımlamak mümkündür. Aşağıda, bir argüman hatasını açıkça temsil etmek için özel bir tip kullanan, önceki örneğin bir varyantı yer alıyor.

---

## ▶️ Çalıştır

```go
package main

import (
    "errors"
    "fmt"
)
```

---

## 🏷️ Özel Hata Tipi Tanımlama

Özel bir hata tipinin adı genellikle `"Error"` son ekiyle biter.

```go
type argError struct {
    arg     int
    message string
}
```

---

## 🧠 Error() Metodu ile error Arayüzünü Uygulama

Bu `Error` metodunu eklemek, `argError` tipinin `error` arayüzünü (*interface*) uygulamasını sağlar.

```go
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}
```

---

## 🧪 Özel Hatayı Döndürme

```go
func f(arg int) (int, error) {
    if arg == 42 {
        // Özel hatamızı döndür.
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}
```

---

## 🔎 errors.As Kullanımı

`errors.As`, `errors.Is`’in daha gelişmiş bir sürümüdür. Verilen bir hatanın (veya zincirindeki herhangi bir hatanın) belirli bir hata tipiyle eşleşip eşleşmediğini kontrol eder ve o tipe dönüştürerek true döndürür. Eşleşme yoksa false döndürür.

```go
func main() {
    _, err := f(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}
```

---

## 💻 CLI Çıktısı

```bash
$ go run custom-errors.go
42
can't work with it
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: *Goroutines*.

