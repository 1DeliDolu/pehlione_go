
## 🧩 Go by Example: Metotlar (Methods)

Go, struct tipleri üzerinde tanımlanan *metotları (methods)* destekler.

### ▶️ Çalıştır

```go
package main
import "fmt"
type rect struct {
    width, height int
}
```

## 📐 Pointer Receiver ile Metot

Bu `area` metodu, alıcı (*receiver*) tipi olarak `*rect` kullanır.

```go
func (r *rect) area() int {
    return r.width * r.height
}
```

## 📏 Value Receiver ile Metot

Metotlar, *pointer receiver* veya *value receiver* tipleri için tanımlanabilir. İşte bir *value receiver* örneği.

```go
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}
```

## 🧪 main İçinde Kullanım

```go
func main() {
    r := rect{width: 10, height: 5}
```

Burada struct’ımız için tanımladığımız 2 metodu çağırıyoruz.

```go
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
```

Go, metot çağrıları için değerler ve pointer’lar arasında dönüşümü otomatik olarak yönetir. Metot çağrılarında kopyalamayı önlemek veya metotun alıcı struct’ı değiştirmesine (*mutate*) izin vermek için *pointer receiver* tipi kullanmak isteyebilirsiniz.

```go
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
```

## 💻 CLI Çıktısı

```bash
$ go run methods.go 
area:  50
perim: 30
area:  50
perim: 30
```

## 🧠 Sonraki Konu

Sonraki olarak, Go’nun ilgili metot kümelerini gruplama ve adlandırma mekanizmasına bakacağız: *interface*’ler.

## ➡️ Sonraki Örnek

Next example: Interfaces.

