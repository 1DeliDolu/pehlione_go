
## 💥 Go by Example: Panic

*Panic* genellikle beklenmedik şekilde bir şeylerin ters gittiği anlamına gelir. Çoğunlukla, normal çalışma sırasında oluşmaması gereken hatalarda veya düzgün (gracefully) şekilde ele almaya hazır olmadığımız durumlarda hızlıca başarısız olmak (fail fast) için kullanırız.

---

## ▶️ Çalıştırma

```go
package main
import (
    "os"
    "path/filepath"
)
func main() {
```

---

## ⚠️ Beklenmedik Hataları Kontrol Etmek İçin Panic

Bu sitede beklenmedik hataları kontrol etmek için her yerde `panic` kullanacağız. Bu, sitede panic olacak şekilde tasarlanmış tek programdır.

```go
    panic("a problem")
```

---

## 🛑 Hata Döndüren Fonksiyonda Abort Etme

`panic`’in yaygın bir kullanımı, bir fonksiyon ele almayı bilmediğimiz (veya ele almak istemediğimiz) bir hata değeri döndürürse işlemi durdurmaktır (abort). İşte yeni bir dosya oluştururken beklenmedik bir hata alırsak panic yapmaya örnek.

```go
    path := filepath.Join(os.TempDir(), "file")
    _, err := os.Create(path)
    if err != nil {
        panic(err)
    }
}
```

---

## 🧾 Çalıştırınca Ne Olur?

Bu programı çalıştırmak panic’e neden olur; bir hata mesajı ve goroutine izlerini (traces) yazdırır ve sıfır olmayan (non-zero) bir status ile çıkar.

İlk `panic` `main` içinde tetiklendiğinde, program kodun geri kalanına ulaşmadan çıkar. Programın geçici (temp) bir dosya oluşturmayı denemesini görmek isterseniz, ilk panic’i yorum satırı (comment) yapın.

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run panic.go
panic: a problem
goroutine 1 [running]:
main.main()
    /.../panic.go:12 +0x47
...
exit status 2
```

---

## 📝 Not

Bazı dillerin birçok hatayı ele almak için istisnalar (exceptions) kullanmasının aksine, Go’da mümkün olduğunca hata belirten dönüş değerleri (*error-indicating return values*) kullanmak idiomatiktir.

---

## ⏭️ Sonraki Örnek

Next example: Defer.

