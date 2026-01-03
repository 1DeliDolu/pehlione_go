
## 🧭 Komut Satırı Alt Komutları

Bazı komut satırı araçları, `go` aracı veya `git` gibi, her birinin kendi bayrak (*flags*) kümesi olan birçok alt komuta (*subcommands*) sahiptir. Örneğin `go build` ve `go get`, `go` aracının iki farklı alt komutudur. `flag` paketi, kendi bayraklarına sahip basit alt komutları kolayca tanımlamamızı sağlar.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "flag"
    "fmt"
    "os"
)
```

```go
func main() {
```

---

## 🧩 Alt Komut Tanımlama

`NewFlagSet` fonksiyonunu kullanarak bir alt komut tanımlarız ve bu alt komuta özel yeni bayrakları tanımlamaya devam ederiz.

```go
    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")
```

Farklı bir alt komut için, farklı desteklenen bayraklar tanımlayabiliriz.

```go
    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")
```

---

## 🥇 İlk Argüman: Alt Komut

Alt komutun, programa verilen ilk argüman olması beklenir.

```go
    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
```

---

## 🔎 Hangi Alt Komutun Çağrıldığını Kontrol Etme

```go
    switch os.Args[1] {
```

Her alt komut için, kendi bayraklarını ayrıştırırız ve sondaki konumsal argümanlara (*positional arguments*) erişebiliriz.

```go
    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}
```

---

## 🛠️ Derleme

```bash
$ go build command-line-subcommands.go 
```

---

## 🧪 foo Alt Komutunu Çalıştırma

Önce `foo` alt komutunu çağırın.

```bash
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]
```

---

## 🧪 bar Alt Komutunu Çalıştırma

Şimdi `bar` alt komutunu deneyin.

```bash
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]
```

---

## ❌ bar, foo Bayraklarını Kabul Etmez

Ancak `bar`, `foo` bayraklarını kabul etmez.

```bash
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
        level
```

---

## ➡️ Sonraki Örnek

Sonraki olarak, programları parametrelemek için yaygın bir başka yöntem olan ortam değişkenlerine (*environment variables*) bakacağız.

Sonraki örnek: **Environment Variables**.

