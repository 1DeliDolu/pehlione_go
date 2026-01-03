
## 🔃 Go by Example: Sıralama (Sorting)

Go’nun `slices` paketi, yerleşik (builtin) tipler ve kullanıcı tanımlı tipler için sıralama (sorting) uygular. Önce yerleşik tipler için sıralamaya bakacağız.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "slices"
)
func main() {
```

---

## 🧬 Yerleşik Tiplerde Sıralama

Sıralama fonksiyonları *generic*tir ve herhangi bir sıralanabilir (ordered) yerleşik tip için çalışır. Sıralanabilir tiplerin listesi için `cmp.Ordered`’a bakın.

```go
    strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)
```

---

## 🔢 int Dilimi Sıralama Örneği

```go
    ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)
```

---

## ✅ Zaten Sıralı mı Kontrol Etme

Ayrıca `slices` paketini kullanarak bir dilimin zaten sıralı olup olmadığını kontrol edebiliriz.

```go
    s := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s)
}
```

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run sorting.go
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
```

---

## ⏭️ Sonraki Örnek

Next example: Sorting by Functions.

