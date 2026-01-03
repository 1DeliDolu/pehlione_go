
## 🧩 Struct Embedding

Go, türlerin daha sorunsuz bir şekilde bileşim (*composition*) ile ifade edilebilmesi için struct ve arayüzlerin (*interfaces*) gömülmesini (*embedding*) destekler. Bu, Go sürüm 1.16+ ile gelen ve dosya/klasörleri uygulama ikilisine gömmek için kullanılan bir Go yönergesi olan `//go:embed` ile karıştırılmamalıdır.

**Run codeCopy code**

```go
package main
import "fmt"
```

## 🧱 base Yapısı

```go
type base struct {
    num int
}
func (b base) describe() string {
    return fmt.Sprintf("num=%v olan base", b.num)
}
```

## 📦 container İçinde base Embedding

Bir `container`, bir `base` gömer. Embedding, adı olmayan bir alan gibi görünür.

```go
type container struct {
    base
    str string
}
```

## 🚀 main Fonksiyonu ve Kullanım

Struct’ları literal ile oluştururken embedding’i açıkça başlatmamız gerekir; burada gömülü tür alan adı olarak görev yapar.

```go
func main() {
    co := container{
        base: base{
            num: 1,
        },
        str: "some name",
    }
```

`base`’in alanlarına `co` üzerinden doğrudan erişebiliriz; ör. `co.num`.

```go
    fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
```

Alternatif olarak, gömülü tür adını kullanarak tam yolu yazabiliriz.

```go
    fmt.Println("also num:", co.base.num)
```

`container`, `base`’i gömdüğü için `base`’in metotları `container`’ın metotları haline gelir. Burada `base`’ten gömülmüş bir metodu `co` üzerinde doğrudan çağırıyoruz.

```go
    fmt.Println("describe:", co.describe())
    type describer interface {
        describe() string
    }
```

Metotları olan struct’ların gömülmesi, başka struct’lara arayüz uygulamalarını kazandırmak için kullanılabilir. Burada `container` artık `base`’i gömdüğü için `describer` arayüzünü uygular.

```go
    var d describer = co
    fmt.Println("describer:", d.describe())
}
```

## 🖥️ Çalıştırma

```bash
$ go run struct-embedding.go
co={num: 1, str: some name}
also num: 1
describe: base with num=1
describer: base with num=1
```

## ⏭️ Sonraki Örnek

Sonraki örnek: *Generics*.

