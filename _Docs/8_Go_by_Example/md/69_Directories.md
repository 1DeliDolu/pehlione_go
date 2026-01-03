
## 📁 Go by Example: Dizinler

Go, dosya sistemindeki dizinlerle çalışmak için birkaç faydalı fonksiyon sunar.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Geçerli çalışma dizininde yeni bir alt dizin oluştur.
    err := os.Mkdir("subdir", 0755)
    check(err)

    // Geçici dizinler oluştururken, bunların silinmesini defer etmek iyi bir pratiktir.
    // os.RemoveAll, tüm dizin ağacını siler (rm -rf benzeri).
    defer os.RemoveAll("subdir")

    // Yeni, boş bir dosya oluşturmak için yardımcı fonksiyon.
    createEmptyFile := func(name string) {
        d := []byte("")
        check(os.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")

    // MkdirAll ile ebeveynler dahil bir dizin hiyerarşisi oluşturabiliriz.
    // Bu, komut satırındaki mkdir -p komutuna benzer.
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)

    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    // ReadDir, dizin içeriğini listeler ve os.DirEntry nesnelerinden oluşan bir slice döndürür.
    c, err := os.ReadDir("subdir/parent")
    check(err)

    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // Chdir, cd’ye benzer şekilde mevcut çalışma dizinini değiştirmemizi sağlar.
    err = os.Chdir("subdir/parent/child")
    check(err)

    // Şimdi geçerli dizini listelerken subdir/parent/child içeriğini göreceğiz.
    c, err = os.ReadDir(".")
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    // Başladığımız yere geri cd.
    err = os.Chdir("../../..")
    check(err)

    // Bir dizini, tüm alt dizinleriyle birlikte özyinelemeli (recursive) olarak da gezebiliriz.
    // WalkDir, ziyaret edilen her dosya veya dizini ele almak için bir callback fonksiyon kabul eder.
    fmt.Println("Visiting subdir")
    err = filepath.WalkDir("subdir", visit)
}

// visit, filepath.WalkDir tarafından özyinelemeli olarak bulunan her dosya veya dizin için çağrılır.
func visit(path string, d fs.DirEntry, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", path, d.IsDir())
    return nil
}
```

### 💻 CLI

```bash
$ go run directories.go
Listing subdir/parent
  child true
  file2 false
  file3 false
Listing subdir/parent/child
  file4 false
Visiting subdir
  subdir true
  subdir/file1 false
  subdir/parent true
  subdir/parent/child true
  subdir/parent/child/file4 false
  subdir/parent/file2 false
  subdir/parent/file3 false
```

## ⏭️ Sonraki Örnek: Geçici Dosyalar ve Dizinler

