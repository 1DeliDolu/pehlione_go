
## 📖 Go by Example: Dosya Okuma

Dosya okuma ve yazma, birçok Go programı için gereken temel görevlerdir. Önce dosya okuma ile ilgili bazı örneklere bakacağız.

### ▶️ Çalıştırma

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

// Dosya okumada çoğu çağrı için hata kontrolü gerekir.
// Bu yardımcı fonksiyon, aşağıdaki hata kontrollerimizi sadeleştirecek.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Belki de en temel dosya okuma işi, bir dosyanın tüm içeriğini belleğe tek seferde almaktır.
    path := filepath.Join(os.TempDir(), "dat")
    dat, err := os.ReadFile(path)
    check(err)
    fmt.Print(string(dat))

    // Çoğu zaman bir dosyanın nasıl ve hangi kısımlarının okunacağı üzerinde daha fazla kontrol istersiniz.
    // Bu işler için, bir os.File değeri elde etmek amacıyla dosyayı Open ile açarak başlayın.
    f, err := os.Open(path)
    check(err)

    // Dosyanın başından bazı byte’lar oku.
    // En fazla 5 byte okunmasına izin ver ama gerçekte kaç byte okunduğunu da not et.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // Dosyada bilinen bir konuma Seek yapıp oradan Read edebilirsiniz.
    o2, err := f.Seek(6, io.SeekStart)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    // Seek etmenin diğer yöntemleri, mevcut imleç (cursor) konumuna göre göreli olabilir,
    _, err = f.Seek(2, io.SeekCurrent)
    check(err)
    // ya da dosyanın sonuna göre göreli olabilir.
    _, err = f.Seek(-4, io.SeekEnd)
    check(err)

    // io paketi, dosya okuma için faydalı olabilecek bazı fonksiyonlar sağlar.
    // Örneğin, yukarıdaki gibi okumalar ReadAtLeast ile daha sağlam şekilde uygulanabilir.
    o3, err := f.Seek(6, io.SeekStart)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // Yerleşik bir rewind yoktur, ama Seek(0, io.SeekStart) bunu sağlar.
    _, err = f.Seek(0, io.SeekStart)
    check(err)

    // bufio paketi, pek çok küçük okuma için verimliliği ve sunduğu ek okuma metotları nedeniyle
    // faydalı olabilecek tamponlu (buffered) bir okuyucu uygular.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // İşiniz bitince dosyayı kapatın (genelde bu, Open’dan hemen sonra defer ile planlanır).
    f.Close()
}
```

### 💻 CLI

```bash
$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
```

### 🧩 Sonraki Adım

Şimdi dosya yazmaya bakacağız.

## ⏭️ Sonraki Örnek: Dosya Yazma

