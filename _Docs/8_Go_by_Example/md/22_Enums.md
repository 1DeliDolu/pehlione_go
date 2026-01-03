
## 🧱 Enum’lar (Enums)

Numaralandırılmış türler (*enums*), *sum type*’ların özel bir durumudur. Bir *enum*, sabit sayıda olası değere sahip olan, her biri ayrı bir ada sahip bir türdür. Go’da ayrı bir dil özelliği olarak *enum* türü yoktur, ancak *enum*’ları mevcut dil kalıplarını kullanarak uygulamak kolaydır.

**Run codeCopy code**

```go
package main
import "fmt"
```

## 🖥️ ServerState Enum Türü

`ServerState` adlı enum türümüzün temelinde `int` türü vardır.

```go
type ServerState int
```

## 🔢 Olası Değerler ve iota

`ServerState` için olası değerler sabitler olarak tanımlanır. Özel anahtar kelime `iota` art arda gelen sabit değerleri otomatik olarak üretir; bu örnekte 0, 1, 2 vb.

```go
const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)
```

## 🧾 Stringer ile Metne Çevirme

`fmt.Stringer` arayüzünü uygulayarak, `ServerState` değerleri yazdırılabilir veya string’e dönüştürülebilir.

Bu, olası değer sayısı çoksa zahmetli hale gelebilir. Böyle durumlarda, süreci otomatikleştirmek için `stringer` aracı `go:generate` ile birlikte kullanılabilir. Daha uzun bir açıklama için bu yazıya bakın.

```go
var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}
func (ss ServerState) String() string {
    return stateName[ss]
}
```

## 🛡️ Derleme Zamanı Tür Güvenliği

Eğer elimizde `int` türünde bir değer varsa, bunu `transition` fonksiyonuna geçemeyiz — derleyici tür uyuşmazlığı nedeniyle hata verecektir. Bu, *enum*’lar için derleme zamanında belirli bir tür güvenliği sağlar.

```go
func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)
    ns2 := transition(ns)
    fmt.Println(ns2)
}
```

## 🔄 Durum Geçişi (transition)

`transition`, bir sunucu için durum geçişini taklit eder; mevcut durumu alır ve yeni bir durum döndürür.

```go
func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
```

Diyelim ki burada bir sonraki durumu belirlemek için bazı koşulları kontrol ediyoruz…

```go
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}
```

## 🖥️ Çalıştırma

```bash
$ go run enums.go
connected
idle
```

## ⏭️ Sonraki Örnek

Sonraki örnek: *Struct Embedding*.

