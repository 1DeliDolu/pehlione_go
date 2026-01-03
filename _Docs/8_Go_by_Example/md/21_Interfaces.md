
## 🧩 Arayüzler (Interfaces)

Arayüzler, *metot imzalarının* adlandırılmış koleksiyonlarıdır.

**Run codeCopy code**

```go
package main
import (
    "fmt"
    "math"
)
```

## 📐 Geometrik Şekiller İçin Temel Arayüz

İşte geometrik şekiller için temel bir arayüz.

```go
type geometry interface {
    area() float64
    perim() float64
}
```

## 🟦🟠 rect ve circle Türlerinde Uygulama

Örneğimizde bu arayüzü `rect` ve `circle` türlerinde uygulayacağız.

```go
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}
```

## ✅ Bir Arayüzü Uygulamak

Go’da bir arayüzü uygulamak için, arayüzdeki tüm metotları uygulamamız yeterlidir. Burada `geometry` arayüzünü `rect` üzerinde uyguluyoruz.

```go
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
```

## 🟠 Daireler İçin Uygulama

Daireler için uygulama.

```go
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
```

## 📏 Arayüz Türü ile Metot Çağrısı

Bir değişkenin türü bir arayüz ise, adlandırılmış arayüzde bulunan metotları çağırabiliriz. İşte bunu kullanarak herhangi bir `geometry` üzerinde çalışabilen genel bir `measure` fonksiyonu.

```go
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
```

## 🔍 Çalışma Zamanı Türünü Öğrenme (Type Assertion)

Bazen bir arayüz değerinin çalışma zamanındaki türünü bilmek faydalıdır. Bir seçenek, burada gösterildiği gibi *type assertion* kullanmaktır; bir diğeri ise *type switch*tir.

```go
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("yarıçapı olan daire", c.radius)
    }
}
```

## 🚀 main Fonksiyonu ve Kullanım

```go
func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
```

`circle` ve `rect` struct türleri `geometry` arayüzünü uygular, bu yüzden bu struct’ların örneklerini `measure` fonksiyonuna argüman olarak kullanabiliriz.

```go
    measure(r)
    measure(c)
    detectCircle(r)
    detectCircle(c)
}
```

## 🖥️ Çalıştırma

```bash
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
circle with radius 5
```

## 🧠 Daha Derin Anlamak

Go’nun arayüzlerinin perde arkasında nasıl çalıştığını anlamak için şu blog yazısına göz atın.

## ⏭️ Sonraki Örnek

Sonraki örnek: *Enums*.

