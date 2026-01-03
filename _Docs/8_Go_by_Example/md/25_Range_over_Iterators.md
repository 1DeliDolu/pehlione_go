
## 🧭 Go by Example: Iterator’lar Üzerinde Range

Go **1.23** sürümünden itibaren Go, *iterator* desteği ekledi; bu sayede neredeyse her şey üzerinde `range` ile dolaşabiliriz.

---

## ▶️ Çalıştırma

```go
package main
import (
    "fmt"
    "iter"
    "slices"
)
```

---

## 🧩 Önceki Örnekteki List Tipine Yeniden Bakış

Önceki örnekte `List` tipi üzerinde, listedeki tüm elemanların bir *slice*’ını döndüren bir `AllElements` metodu vardı. Go *iterator*’ları ile bunu daha iyi yapabiliriz — aşağıdaki gibi.

```go
type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}
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

---

## 🔁 All Metodu Iterator Döndürür

`All`, bir iterator döndürür; Go’da iterator, özel bir imzaya sahip bir fonksiyondur.

```go
func (lst *List[T]) All() iter.Seq[T] {
    return func(yield func(T) bool) {
```

Iterator fonksiyonu, parametre olarak başka bir fonksiyon alır; bu fonksiyona geleneksel olarak `yield` denir (ama isim keyfi olabilir). Iterator, dolaşmak istediğimiz her eleman için `yield` fonksiyonunu çağırır ve erken sonlandırma ihtimali için `yield`’in dönüş değerini dikkate alır.

```go
        for e := lst.head; e != nil; e = e.next {
            if !yield(e.val) {
                return
            }
        }
    }
}
```

---

## ♾️ Iterator İçin Altyapı Veri Yapısı Zorunlu Değildir

Iterasyon bir altyapı veri yapısı gerektirmez ve hatta sonlu olmak zorunda bile değildir. İşte Fibonacci sayıları üzerinde bir iterator döndüren bir fonksiyon: `yield` true döndürdüğü sürece çalışmaya devam eder.

```go
func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 1, 1
        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}
```

---

## 🧪 Kullanım Örneği

```go
func main() {
    lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
```

`List.All` bir iterator döndürdüğü için, bunu normal bir `range` döngüsünde kullanabiliriz.

```go
    for e := range lst.All() {
        fmt.Println(e)
    }
```

`slices` gibi paketlerde iterator’larla çalışmak için bir dizi faydalı fonksiyon bulunur. Örneğin `Collect`, herhangi bir iterator’ı alır ve tüm değerlerini bir slice içine toplar.

```go
    all := slices.Collect(lst.All())
    fmt.Println("all:", all)
    for n := range genFib() {
```

Döngü `break` ya da erken bir `return` ile kesildiğinde, iterator’a geçirilen `yield` fonksiyonu `false` döndürecektir.

```go
        if n >= 10 {
            break
        }
        fmt.Println(n)
    }
}
```

---

## 🖥️ Çalıştırma Komutu ve Çıktı

```bash
$ go run range-over-iterators.go
```

```text
10
13
23
all: [10 13 23]
1
1
2
3
5
8
```

---

## ⏭️ Sonraki Örnek

Sonraki örnek: Errors.

