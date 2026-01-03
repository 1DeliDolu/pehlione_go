
## 📡 Sinyaller

Bazen Go programlarımızın Unix sinyallerini akıllıca ele almasını isteriz. Örneğin, bir sunucunun `SIGTERM` aldığında zarifçe (*gracefully*) kapanmasını veya bir komut satırı aracının `SIGINT` aldığında girdi işlemeyi durdurmasını isteyebiliriz. Go’da sinyallerin kanallar (*channels*) ile nasıl ele alınacağı aşağıdadır.

---

## ▶️ Çalıştırma

```go
package main
```

```go
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)
```

```go
func main() {
```

---

## 📥 Sinyal Bildirimlerini Kanaldan Alma

Go’nun sinyal bildirim mekanizması, `os.Signal` değerlerini bir kanal üzerinden göndererek çalışır. Bu bildirimleri almak için bir kanal oluşturacağız. Bu kanalın *buffered* olması gerektiğine dikkat edin.

```go
    sigs := make(chan os.Signal, 1)
```

`signal.Notify`, verilen kanalı belirtilen sinyallerin bildirimlerini alacak şekilde kaydeder.

```go
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
```

Sinyalleri burada `main` fonksiyonunda `sigs` kanalından alabilirdik; ancak zarif kapanış (*graceful shutdown*) için daha gerçekçi bir senaryoyu göstermek amacıyla bunu ayrı bir goroutine içinde nasıl yapabileceğimize bakalım.

```go
    done := make(chan bool, 1)
```

Bu goroutine, sinyaller için bloklayan bir receive çalıştırır. Bir sinyal aldığında onu yazdırır ve ardından programın bitirebileceğini bildirir.

```go
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()
```

Program, beklenen sinyali alana kadar burada bekler (goroutine’in `done` kanalına bir değer göndermesiyle anlaşılır) ve ardından çıkar.

```go
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
```

---

## 🧪 Örnek Çalıştırma

Bu programı çalıştırdığımızda, bir sinyal bekleyerek bloklanır. `ctrl-C` yazarak (terminalin `^C` olarak gösterdiği), bir `SIGINT` sinyali gönderebiliriz; bu da programın `interrupt` yazdırıp çıkmasına neden olur.

```bash
$ go run signals.go
awaiting signal
^C
interrupt
exiting
```

---

## ➡️ Sonraki Örnek

Sonraki örnek: **Exit**.

