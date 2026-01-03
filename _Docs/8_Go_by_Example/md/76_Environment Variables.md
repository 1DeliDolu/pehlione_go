
## 🌿 Ortam Değişkenleri

Ortam değişkenleri (*environment variables*), Unix programlarına yapılandırma bilgisini iletmek için evrensel bir mekanizmadır. Ortam değişkenlerini nasıl ayarlayacağımıza, alacağımıza ve listeleyeceğimize bakalım.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "fmt"
    "os"
    "strings"
)
```

```go
func main() {
```

---

## 🧷 Ayarlama ve Okuma

Bir anahtar/değer (*key/value*) çifti ayarlamak için `os.Setenv` kullanın. Bir anahtarın değerini almak için `os.Getenv` kullanın. Anahtar ortamda mevcut değilse bu fonksiyon boş bir *string* döndürür.

```go
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))
```

---

## 📋 Ortam Değişkenlerini Listeleme

Ortamdaki tüm anahtar/değer çiftlerini listelemek için `os.Environ` kullanın. Bu, `KEY=value` biçiminde *string*’lerden oluşan bir slice döndürür. Anahtar ve değeri almak için `strings.SplitN` ile bölebilirsiniz. Burada tüm anahtarları yazdırıyoruz.

```go
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}
```

---

## 🧪 Program Çıktısı

Programı çalıştırmak, program içinde ayarladığımız `FOO` değerini aldığımızı; ancak `BAR`’ın boş olduğunu gösterir.

```bash
$ go run environment-variables.go
FOO: 1
BAR: 
```

Ortamdaki anahtarların listesi, kullandığınız makineye göre değişir.

```text
TERM_PROGRAM
PATH
SHELL
...
FOO
```

---

## 🧪 Ortamdan BAR Değeri Almak

Eğer önce ortamda `BAR`’ı ayarlarsak, çalışan program o değeri alır.

```bash
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Logging**.

