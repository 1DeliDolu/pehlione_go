
## 🧩 Go by Example: Fonksiyonlarla Sıralama (Sorting by Functions)

Bazen bir koleksiyonu, doğal sıralamasından (natural order) farklı bir ölçüte göre sıralamak isteriz. Örneğin, string’leri alfabetik olarak değil de uzunluklarına göre sıralamak isteyebiliriz. İşte Go’da özel (custom) sıralamalara bir örnek.

---

## ▶️ Çalıştırma

```go
package main
import (
    "cmp"
    "fmt"
    "slices"
)
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
```

---

## 📏 String Uzunluğuna Göre Karşılaştırma

String uzunlukları için bir karşılaştırma (comparison) fonksiyonu uygularız. `cmp.Compare` bu iş için faydalıdır.

```go
    lenCmp := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }
```

---

## 🧷 slices.SortFunc ile Özel Sıralama

Artık `slices.SortFunc`’u bu özel karşılaştırma fonksiyonuyla çağırarak `fruits`’i isim uzunluğuna göre sıralayabiliriz.

```go
    slices.SortFunc(fruits, lenCmp)
    fmt.Println(fruits)
```

---

## 👥 Yerleşik Olmayan Tiplerde Sıralama

Aynı tekniği, yerleşik (built-in) tip olmayan değerlerden oluşan bir dilimi sıralamak için de kullanabiliriz.

```go
    type Person struct {
        name string
        age  int
    }
    people := []Person{
        Person{name: "Jax", age: 37},
        Person{name: "TJ", age: 25},
        Person{name: "Alex", age: 72},
    }
```

---

## 🎂 Yaşa Göre Sıralama

`people`’ı yaşa göre `slices.SortFunc` kullanarak sıralayın.

Not: Eğer `Person` struct’ı büyükse, dilimin `*Person` içermesini isteyebilir ve sıralama fonksiyonunu buna göre ayarlayabilirsiniz. Emin değilseniz, benchmark alın!

```go
    slices.SortFunc(people,
        func(a, b Person) int {
            return cmp.Compare(a.age, b.age)
        })
    fmt.Println(people)
}
```

---

## 💻 CLI Çalıştırma ve Çıktı

```bash
$ go run sorting-by-functions.go 
[kiwi peach banana]
[{TJ 25} {Jax 37} {Alex 72}]
```

---

## ⏭️ Sonraki Örnek

Next example: Panic.

