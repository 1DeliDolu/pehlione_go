
## 🧪 Go by Example: Geçici Dosyalar ve Dizinler

Program çalışması boyunca, çoğu zaman program sonlandığında artık ihtiyaç duyulmayan veriler oluşturmak isteriz. Geçici dosyalar ve dizinler bu amaç için kullanışlıdır; çünkü zamanla dosya sistemini kirletmezler.

### ▶️ Çalıştırma

```go
package main

import (
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
    // Geçici dosya oluşturmanın en kolay yolu os.CreateTemp çağırmaktır.
    // Bu, bir dosya oluşturur ve okuma/yazma için açar.
    // İlk argüman olarak "" veriyoruz; böylece os.CreateTemp dosyayı işletim sistemimizin varsayılan konumunda oluşturur.
    f, err := os.CreateTemp("", "sample")
    check(err)

    // Geçici dosyanın adını göster.
    // Unix tabanlı işletim sistemlerinde dizin muhtemelen /tmp olacaktır.
    // Dosya adı, os.CreateTemp’e ikinci argüman olarak verilen prefix ile başlar
    // ve geri kalanı eşzamanlı çağrıların her zaman farklı dosya adları oluşturmasını garanti etmek için otomatik seçilir.
    fmt.Println("Temp file name:", f.Name())

    // İşimiz bittiğinde dosyayı temizle.
    // İşletim sistemi muhtemelen geçici dosyaları bir süre sonra kendisi temizler,
    // ama bunu açıkça yapmak iyi bir pratiktir.
    defer os.Remove(f.Name())

    // Dosyaya biraz veri yazabiliriz.
    _, err = f.Write([]byte{1, 2, 3, 4})
    check(err)

    // Çok sayıda geçici dosya yazmayı planlıyorsak, geçici bir dizin oluşturmayı tercih edebiliriz.
    // os.MkdirTemp’in argümanları CreateTemp ile aynıdır, ancak açık bir dosya yerine bir dizin adı döndürür.
    dname, err := os.MkdirTemp("", "sampledir")
    check(err)

    fmt.Println("Temp dir name:", dname)
    defer os.RemoveAll(dname)

    // Artık geçici dizinimizi başa ekleyerek geçici dosya adları üretebiliriz.
    fname := filepath.Join(dname, "file1")
    err = os.WriteFile(fname, []byte{1, 2}, 0666)
    check(err)
}
```

### 💻 CLI

```bash
$ go run temporary-files-and-directories.go
Temp file name: /tmp/sample610887201
Temp dir name: /tmp/sampledir898854668
```

## ⏭️ Sonraki Örnek: Embed Direktifi

