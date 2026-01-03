
## 🧮 Go by Example: Values

Go’da *string*, tamsayı (*integer*), kayan noktalı (*float*), boolean vb. çeşitli değer türleri vardır. İşte birkaç temel örnek.

---

## 🧾 Kaynak Kod

```go
package main

import "fmt"

func main() {
    // String’ler, + ile birleştirilebilir.
    fmt.Println("go" + "lang")

    // Tamsayılar ve kayan noktalılar.
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // Boolean’lar ve beklediğiniz boolean operatörleri.
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

---

## ▶️ Çalıştırma

```bash
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

---

## 📌 Sonraki Örnek

**Next example:** *Variables.*

