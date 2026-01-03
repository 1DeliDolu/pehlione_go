
## ✍️ Go by Example: Dosya Yazma

Go’da dosya yazma, daha önce okuma için gördüğümüz kalıplara benzer biçimde ilerler.

### ▶️ Çalıştırma

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Başlangıç olarak, bir string’i (ya da yalnızca byte’ları) bir dosyaya yazmak şöyle yapılır.
    d1 := []byte("hello\ngo\n")
    path1 := filepath.Join(os.TempDir(), "dat1")
    err := os.WriteFile(path1, d1, 0644)
    check(err)

    // Daha ayrıntılı (granular) yazımlar için, yazma amaçlı bir dosya açın.
    path2 := filepath.Join(os.TempDir(), "dat2")
    f, err := os.Create(path2)
    check(err)

    // Bir dosyayı açtıktan hemen sonra Close’u defer etmek Go’da idiomatiktir.
    defer f.Close()

    // Byte slice’ları beklediğiniz gibi Write edebilirsiniz.
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // WriteString de mevcuttur.
    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    // Yazıları kalıcı depolamaya flush etmek için Sync çağırın.
    f.Sync()

    // bufio, daha önce gördüğümüz tamponlu okuyuculara ek olarak tamponlu yazıcılar da sağlar.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    // Flush, tamponlanan tüm işlemlerin alttaki writer’a uygulanmasını garanti eder.
    w.Flush()
}
```

### 💻 CLI

```bash
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes
```

### 📄 Yazılan Dosyaların İçeriğini Kontrol Etme

```bash
$ cat /tmp/dat1
hello
go

$ cat /tmp/dat2
some
writes
buffered
```

### 🧩 Sonraki Adım

Şimdi, az önce gördüğümüz dosya G/Ç (I/O) fikirlerinden bazılarını `stdin` ve `stdout` akışlarına uygulamaya bakacağız.

## ⏭️ Sonraki Örnek: Satır Filtreleri

