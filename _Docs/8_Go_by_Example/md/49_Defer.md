
## 🧹 Go by Example: Defer

`defer`, bir fonksiyon çağrısının programın yürütülmesinde daha sonra gerçekleştirilmesini sağlamak için kullanılır; genellikle temizlik (cleanup) amaçlarıyla. `defer`, diğer dillerde örneğin `ensure` ve `finally`’ın kullanıldığı yerlere benzer şekilde sıkça kullanılır.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "os"
    "path/filepath"
)
```

---

## 📄 Dosya Oluşturma, Yazma ve Kapatma

Diyelim ki bir dosya oluşturmak, içine yazmak ve işimiz bitince kapatmak istiyoruz. İşte bunu `defer` ile nasıl yapabileceğimiz.

```go
func main() {
```

`createFile` ile bir dosya nesnesi aldıktan hemen sonra, `closeFile` ile o dosyanın kapatılmasını `defer` ederiz. Bu, `writeFile` bittikten sonra, kapsayan fonksiyonun (`main`) sonunda çalıştırılacaktır.

```go
    path := filepath.Join(os.TempDir(), "defer.txt")
    f := createFile(path)
    defer closeFile(f)
    writeFile(f)
}
```

---

## 🏗️ Dosya Oluşturma Fonksiyonu

```go
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
```

---

## ✍️ Dosyaya Yazma Fonksiyonu

```go
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
```

---

## 🔒 Dosya Kapatma ve Hata Kontrolü

Bir dosyayı kapatırken, ertelenmiş (deferred) bir fonksiyonda bile, hataları kontrol etmek önemlidir.

```go
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()
    if err != nil {
        panic(err)
    }
}
```

---

## ✅ Çalıştırma Sonucu

Programı çalıştırmak, dosyanın yazıldıktan sonra kapatıldığını doğrular.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run defer.go
creating
writing
closing
```

---

## ⏭️ Sonraki Örnek

Next example: Recover.

