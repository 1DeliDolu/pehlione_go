
## 🧹 Go by Example: Satır Filtreleri

Bir satır filtresi, `stdin` üzerinden girdi okuyan, bunu işleyen ve ardından türetilmiş bir sonucu `stdout`’a yazdıran yaygın bir program türüdür. `grep` ve `sed` yaygın satır filtreleridir.

Aşağıda, tüm girdi metninin büyük harfe çevrilmiş bir sürümünü yazdıran Go ile bir satır filtresi örneği yer alır. Kendi Go satır filtrelerinizi yazmak için bu deseni kullanabilirsiniz.

### ▶️ Çalıştırma

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // Tamponlanmamış os.Stdin’i tamponlu bir scanner ile sarmak, scanner’ı bir sonraki token’a ilerleten
    // kullanışlı bir Scan metodu sağlar; varsayılan scanner’da bu bir sonraki satırdır.
    scanner := bufio.NewScanner(os.Stdin)

    // Text, girdiden mevcut token’ı (burada bir sonraki satırı) döndürür.
    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())

        // Büyük harfe çevrilmiş satırı yazdır.
        fmt.Println(ucl)
    }

    // Scan sırasında hata olup olmadığını kontrol et.
    // Dosya sonu (EOF) beklenen bir durumdur ve Scan tarafından hata olarak raporlanmaz.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
```

### 🧪 Deneme

Önce birkaç küçük harfli satır içeren bir dosya oluşturun.

```bash
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines
```

Ardından satır filtresini kullanarak satırları büyük harfe çevirin.

```bash
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
```

## ⏭️ Sonraki Örnek: Dosya Yolları

