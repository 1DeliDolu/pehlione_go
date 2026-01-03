
## 🧱 Go by Example: Struct’lar

Go’nun *struct*’ları, alanlardan (*field*) oluşan tipli koleksiyonlardır. Verileri bir araya getirip kayıtlar (*record*) oluşturmak için kullanışlıdır.

### ▶️ Çalıştır

```go
package main
import "fmt"
```

## 👤 Struct Tipi Tanımı

Bu `person` struct tipi, `name` ve `age` alanlarına sahiptir.

```go
type person struct {
    name string
    age  int
}
```

## 🏗️ Constructor Fonksiyonu

`newPerson`, verilen isimle yeni bir `person` struct’ı oluşturur.

```go
func newPerson(name string) *person {
```

Go, *garbage collected* bir dildir; yerel bir değişkene ait pointer’ı güvenle döndürebilirsiniz — yalnızca ona ait aktif referans kalmadığında *garbage collector* tarafından temizlenir.

```go
    p := person{name: name}
    p.age = 42
    return &p
}
```

## 🧪 Kullanım Örnekleri

```go
func main() {
```

Bu sözdizimi yeni bir struct oluşturur.

```go
    fmt.Println(person{"Bob", 20})
```

Struct başlatırken alan isimlerini verebilirsiniz.

```go
    fmt.Println(person{name: "Alice", age: 30})
```

Atlanan alanlar *zero value* alır.

```go
    fmt.Println(person{name: "Fred"})
```

Başına `&` koymak, struct’a bir pointer verir.

```go
    fmt.Println(&person{name: "Ann", age: 40})
```

Yeni struct oluşturmayı constructor fonksiyonlarında kapsüllemek *idiomatic* bir yaklaşımdır.

```go
    fmt.Println(newPerson("Jon"))
```

Struct alanlarına nokta (`.`) ile erişilir.

```go
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
```

Struct pointer’larıyla da nokta kullanabilirsiniz — pointer’lar otomatik olarak *dereference* edilir.

```go
    sp := &s
    fmt.Println(sp.age)
```

Struct’lar değiştirilebilirdir (*mutable*).

```go
    sp.age = 51
    fmt.Println(sp.age)
```

Bir struct tipi yalnızca tek bir değer için kullanılacaksa, ona isim vermek zorunda değiliz. Değer, anonim bir struct tipine sahip olabilir. Bu teknik, genellikle *table-driven tests* için kullanılır.

```go
    dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}
```

## 💻 CLI Çıktısı

```bash
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```

## ➡️ Sonraki Örnek

Next example: Methods.

