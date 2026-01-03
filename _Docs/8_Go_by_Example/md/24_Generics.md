
## 🧬 Generics

Sürüm 1.18’den itibaren Go, *type parameters* olarak da bilinen *generics* desteğini eklemiştir.

**Run codeCopy code**

```go
package main
import "fmt"
```

## 🧩 Genel (Generic) Fonksiyon Örneği: SlicesIndex

Genel bir fonksiyon örneği olarak `SlicesIndex`, herhangi bir *comparable* türden bir slice ve o türden bir eleman alır ve `v` değerinin `s` içinde ilk geçtiği konumun indeksini döndürür; eğer yoksa `-1` döndürür. `comparable` kısıtı, bu türün değerlerini `==` ve `!=` operatörleriyle karşılaştırabileceğimiz anlamına gelir. Bu tür imzasının daha ayrıntılı bir açıklaması için şu blog yazısına bakın. Bu fonksiyonun standart kütüphanede `slices.Index` olarak bulunduğunu unutmayın.

```go
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
    for i := range s {
        if v == s[i] {
            return i
        }
    }
    return -1
}
```

## 🔗 Genel (Generic) Tür Örneği: List

Genel bir tür örneği olarak `List`, herhangi bir türden değerlere sahip tek yönlü bağlı listedir.

```go
type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}
```

## 🧱 Genel Türlerde Metot Tanımlama: Push

Genel türlerde, normal türlerde yaptığımız gibi metotlar tanımlayabiliriz; ancak tür parametrelerini yerinde tutmamız gerekir. Tür `List[T]`’dir, `List` değildir.

```go
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
```

## 📦 Tüm Elemanları Slice Olarak Alma: AllElements

`AllElements`, `List` elemanlarının tamamını bir slice olarak döndürür. Sonraki örnekte, özel türlerin tüm elemanları üzerinde gezinmenin daha idiomatik bir yolunu göreceğiz.

```go
func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}
```

## 🚀 main Fonksiyonu ve Kullanım

```go
func main() {
    var s = []string{"foo", "bar", "zoo"}
```

Genel fonksiyonları çağırırken, çoğu zaman tür çıkarımına (*type inference*) güvenebiliriz. `SlicesIndex` çağrısında `S` ve `E` için türleri belirtmek zorunda değiliz — derleyici bunları otomatik olarak çıkarır.

```go
    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
```

… ancak istersek bunları açıkça da belirtebiliriz.

```go
    _ = SlicesIndex[[]string, string](s, "zoo")
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}
```

## 🖥️ Çalıştırma

```bash
$ go run generics.go
index of zoo: 2
list: [10 13 23]
```

## ⏭️ Sonraki Örnek

Sonraki örnek: *Range over Iterators*.

