
## 👋 Go by Example: Hello World

İlk programımız klasik “hello world” mesajını yazdıracak. İşte tam kaynak kodu.

---

## 🧾 Kaynak Kod

```go
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

---

## ▶️ Programı Çalıştırma

Programı çalıştırmak için, kodu `hello-world.go` dosyasına koyun ve `go run` kullanın.

```bash
$ go run hello-world.go
hello world
```

---

## 🏗️ Binary Oluşturma

Bazen programlarımızı binary’lere derlemek isteyeceğiz. Bunu `go build` kullanarak yapabiliriz.

```bash
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

---

## ✅ Oluşturulan Binary’yi Çalıştırma

Daha sonra oluşturulan binary’yi doğrudan çalıştırabiliriz.

```bash
$ ./hello-world
hello world
```

---

## 📌 Sonraki Adım

Artık temel Go programlarını çalıştırıp derleyebildiğimize göre, dil hakkında daha fazlasını öğrenelim.

**Sonraki örnek:** *Values.*

