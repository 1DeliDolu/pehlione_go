
## 🧾 Go by Example: Variables

Go’da değişkenler açıkça bildirilir ve derleyici tarafından örneğin fonksiyon çağrılarının tür doğruluğunu (*type-correctness*) denetlemek için kullanılır.

---

## 🧾 Kaynak Kod

```go
package main

import "fmt"

func main() {
    // var, 1 veya daha fazla değişken bildirir.
    var a = "initial"
    fmt.Println(a)

    // Birden fazla değişkeni aynı anda bildirebilirsiniz.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go, başlangıç değeri verilen değişkenlerin türünü çıkarır.
    var d = true
    fmt.Println(d)

    // Karşılık gelen bir başlangıç değeri olmadan bildirilen değişkenler sıfır değer (*zero value*) alır.
    // Örneğin int için sıfır değer 0’dır.
    var e int
    fmt.Println(e)

    // := sözdizimi, değişken bildirme ve başlatma için kısayoldur.
    // Bu örnekte, var f string = "apple" ifadesine denktir.
    // Bu sözdizimi yalnızca fonksiyonların içinde kullanılabilir.
    f := "apple"
    fmt.Println(f)
}
```

---

## ▶️ Çalıştırma

```bash
$ go run variables.go
initial
1 2
true
0
apple
```

---

## 📌 Sonraki Örnek

**Next example:** *Constants.*

