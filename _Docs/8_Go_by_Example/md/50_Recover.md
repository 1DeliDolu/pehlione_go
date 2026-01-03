
## 🛟 Go by Example: Recover

Go, yerleşik `recover` fonksiyonunu kullanarak bir *panic*’ten kurtulmayı mümkün kılar. Bir `recover`, panic’in programı abort etmesini engelleyebilir ve bunun yerine yürütmenin devam etmesine izin verebilir.

Bunun faydalı olabileceği bir örnek: Bir sunucu, istemci bağlantılarından biri kritik bir hata gösterdiğinde çökmek istemez. Bunun yerine, o bağlantıyı kapatmak ve diğer istemcilere hizmet vermeye devam etmek ister. Aslında Go’nun `net/http` paketi HTTP sunucuları için varsayılan olarak bunu yapar.

---

## ▶️ Çalıştırma

```go
package main
import "fmt"
```

---

## 💥 Panic Üreten Fonksiyon

Bu fonksiyon panic oluşturur.

```go
func mayPanic() {
    panic("a problem")
}
```

---

## 🧷 Defer İçinde recover Kullanımı

`recover`, ertelenmiş (deferred) bir fonksiyon içinde çağrılmalıdır. Kapsayan fonksiyon panic yaptığında, `defer` devreye girer ve içindeki `recover` çağrısı panic’i yakalar.

```go
func main() {
```

`recover`’ın dönüş değeri, `panic` çağrısında yükseltilen hatadır.

```go
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()
    mayPanic()
```

Bu kod çalışmayacaktır, çünkü `mayPanic` panic yapar. `main` yürütmesi panic noktasında durur ve ertelenmiş closure içinde devam eder.

```go
    fmt.Println("After mayPanic()")
}
```

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run recover.go
Recovered. Error:
 a problem
```

---

## ⏭️ Sonraki Örnek

Next example: String Functions.

