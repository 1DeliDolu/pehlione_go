
## 🚪 Çıkış

Belirli bir durum koduyla (*status*) anında çıkmak için `os.Exit` kullanın.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "fmt"
    "os"
)
```

```go
func main() {
```

---

## ⛔ os.Exit ve defer

`os.Exit` kullanıldığında `defer`’lar çalıştırılmaz; bu yüzden aşağıdaki `fmt.Println` hiçbir zaman çağrılmayacaktır.

```go
    defer fmt.Println("!")
```

---

## 🔢 Durum Kodu ile Çıkış

3 durum kodu ile çıkın.

```go
    os.Exit(3)
}
```

---

## 📝 Not

C gibi dillerin aksine, Go `main` fonksiyonundan dönen bir *integer* dönüş değerini çıkış durumunu belirtmek için kullanmaz. Sıfır olmayan bir durumla çıkmak istiyorsanız `os.Exit` kullanmalısınız.

---

## 🧪 go run ile Çalıştırma

`exit.go` dosyasını `go run` ile çalıştırırsanız, çıkış durumu `go` tarafından yakalanır ve yazdırılır.

```bash
$ go run exit.go
exit status 3
```

---

## 🧪 Binary Derleyip Çalıştırma

Bir binary derleyip çalıştırarak durum kodunu terminalde görebilirsiniz.

```bash
$ go build exit.go
$ ./exit
$ echo $?
3
```

Programımızdaki `!` ifadesinin hiç yazdırılmadığına dikkat edin.

