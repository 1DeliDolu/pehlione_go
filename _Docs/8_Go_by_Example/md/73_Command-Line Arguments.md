
## 🧾 Go by Example: Komut Satırı Argümanları

Komut satırı argümanları, programların çalıştırılmasını parametrelemenin yaygın bir yoludur. Örneğin, `go run hello.go`, `go` programına `run` ve `hello.go` argümanlarını kullanır.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args, ham (raw) komut satırı argümanlarına erişim sağlar.
    // Bu slice’ın ilk değeri programın yolu (path) olduğuna dikkat edin;
    // os.Args[1:] ise programın argümanlarını tutar.
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // Normal indeksleme ile tek tek argümanlara erişebilirsiniz.
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
```

### 💻 CLI

Komut satırı argümanlarıyla deneme yapmak için önce `go build` ile bir binary üretmek en iyisidir.

```bash
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
```

### 🧩 Sonraki Adım

Şimdi `flag`’lerle daha gelişmiş komut satırı işlemesine bakacağız.

## ⏭️ Sonraki Örnek: Komut Satırı Flag’leri

