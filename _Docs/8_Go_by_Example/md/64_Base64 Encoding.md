
## 🧩 Go by Example: Base64 Kodlama

Go, base64 kodlama/kod çözme için yerleşik destek sağlar.

### ▶️ Çalıştırma

```go
package main

// Bu sözdizimi encoding/base64 paketini varsayılan base64 yerine b64 adıyla içe aktarır.
// Aşağıda biraz yer kazandırır.
import (
    b64 "encoding/base64"
    "fmt"
)

func main() {
    // Encode/decode edeceğimiz string.
    data := "abc123!?$*&()'-=@~"

    // Go hem standart hem de URL-uyumlu base64 destekler.
    // Standart encoder ile kodlama şöyle yapılır. Encoder bir []byte ister, bu yüzden string’imizi o tipe çeviririz.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Decode işlemi hata döndürebilir; girdinin iyi biçimlendirildiğini zaten bilmiyorsanız kontrol edebilirsiniz.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // Bu, URL-uyumlu base64 formatı ile encode/decode eder.
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
```

### 🧾 Not

String, standart ve URL base64 encoder’ları ile biraz farklı değerlere kodlanır (sondaki `+` yerine `-` gibi), fakat ikisi de istendiği gibi orijinal string’e geri decode edilir.

### 💻 CLI

```bash
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~
YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
```

## ⏭️ Sonraki Örnek: Dosya Okuma

