
## 🔐 Go by Example: SHA256 Hash’leri

SHA256 hash’leri, ikili (binary) ya da metin (text) blokları için kısa kimlikler hesaplamak amacıyla sıkça kullanılır. Örneğin, TLS/SSL sertifikaları bir sertifikanın imzasını hesaplamak için SHA256 kullanır. Go’da SHA256 hash’lerinin nasıl hesaplanacağı aşağıdadır.

### ▶️ Çalıştırma

```go
package main

// Go, çeşitli crypto/* paketlerinde birden fazla hash fonksiyonu uygular.
import (
    "crypto/sha256"
    "fmt"
)

func main() {
    s := "sha256 this string"

    // Burada yeni bir hash ile başlıyoruz.
    h := sha256.New()

    // Write byte bekler. Bir string s varsa, byte’a zorlamak için []byte(s) kullanın.
    h.Write([]byte(s))

    // Bu, nihai hash sonucunu bir byte slice olarak verir.
    // Sum’a verilen argüman, var olan bir byte slice’a ekleme yapmak için kullanılabilir; genellikle gerekmez.
    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
```

### 💻 CLI

```bash
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...
```

### 🧩 Ek Notlar

Diğer hash’leri de yukarıda gösterilen kalıba benzer şekilde hesaplayabilirsiniz. Örneğin, SHA512 hash’leri hesaplamak için `crypto/sha512` içe aktarın ve `sha512.New()` kullanın.

Kriptografik olarak güvenli hash’lere ihtiyacınız varsa, hash gücünü dikkatle araştırmanız gerektiğini unutmayın.

## ⏭️ Sonraki Örnek: Base64 Kodlama

