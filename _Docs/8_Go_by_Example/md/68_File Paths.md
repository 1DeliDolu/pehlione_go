
## 🗂️ Go by Example: Dosya Yolları

`filepath` paketi, dosya yollarını işletim sistemleri arasında taşınabilir olacak şekilde ayrıştırmak ve oluşturmak için fonksiyonlar sağlar; örneğin Linux’ta `dir/file`, Windows’ta `dir\file` gibi.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "path/filepath"
    "strings"
)

func main() {
    // Join, yolları taşınabilir şekilde oluşturmak için kullanılmalıdır.
    // İstediğiniz sayıda argüman alır ve bunlardan hiyerarşik bir yol oluşturur.
    p := filepath.Join("dir1", "dir2", "filename")
    fmt.Println("p:", p)

    // / veya \ karakterlerini elle birleştirmek yerine her zaman Join kullanmalısınız.
    // Taşınabilirliğe ek olarak, Join gereksiz ayırıcıları ve dizin değişimlerini kaldırarak
    // yolları normalize eder.
    fmt.Println(filepath.Join("dir1//", "filename"))
    fmt.Println(filepath.Join("dir1/../dir1", "filename"))

    // Dir ve Base, bir yolu dizin ve dosya olarak ayırmak için kullanılabilir.
    // Alternatif olarak Split, ikisini tek çağrıda döndürür.
    fmt.Println("Dir(p):", filepath.Dir(p))
    fmt.Println("Base(p):", filepath.Base(p))

    // Bir yolun mutlak (absolute) olup olmadığını kontrol edebiliriz.
    fmt.Println(filepath.IsAbs("dir/file"))
    fmt.Println(filepath.IsAbs("/dir/file"))

    filename := "config.json"

    // Bazı dosya adlarının nokta sonrası bir uzantısı vardır.
    // Ext ile bu uzantıyı ayırabiliriz.
    ext := filepath.Ext(filename)
    fmt.Println(ext)

    // Uzantısı çıkarılmış dosya adını bulmak için strings.TrimSuffix kullanın.
    fmt.Println(strings.TrimSuffix(filename, ext))

    // Rel, bir base ile target arasında göreli (relative) bir yol bulur.
    // Target, base’e göre göreli hale getirilemiyorsa hata döndürür.
    rel, err := filepath.Rel("a/b", "a/b/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)

    rel, err = filepath.Rel("a/b", "a/c/t/file")
    if err != nil {
        panic(err)
    }
    fmt.Println(rel)
}
```

### 💻 CLI

```bash
$ go run file-paths.go
p: dir1/dir2/filename
dir1/filename
dir1/filename
Dir(p): dir1/dir2
Base(p): filename
false
true
.json
config
t/file
../c/t/file
```

## ⏭️ Sonraki Örnek: Dizinler

