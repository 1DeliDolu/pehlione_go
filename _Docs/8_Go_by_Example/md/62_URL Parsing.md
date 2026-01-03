
## 🌐 Go by Example: URL Ayrıştırma

URL’ler, kaynakları konumlandırmak için tekdüze bir yol sağlar. Go’da URL’lerin nasıl ayrıştırılacağı aşağıdadır.

### ▶️ Çalıştırma

```go
package main

import (
    "fmt"
    "net"
    "net/url"
)

func main() {
    // Şema (scheme), kimlik doğrulama bilgileri, host, port, path, query parametreleri ve fragment içeren bu örnek URL’yi ayrıştıracağız.
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // URL’yi ayrıştırın ve hata olmadığından emin olun.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // Scheme’e erişmek doğrudandır.
    fmt.Println(u.Scheme)

    // User tüm kimlik doğrulama bilgisini içerir; tekil değerler için Username ve Password çağırın.
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // Host, varsa hem hostname’i hem de port’u içerir. Ayırmak için SplitHostPort kullanın.
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // Burada path’i ve # sonrası fragment’i çıkarıyoruz.
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // Query parametrelerini k=v formatındaki string olarak almak için RawQuery kullanın.
    // Query parametrelerini bir map’e de ayrıştırabilirsiniz.
    // Ayrıştırılan query parametre map’leri string -> []string şeklindedir; sadece ilk değeri istiyorsanız [0] indeksleyin.
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
```

### 💻 CLI

```bash
$ go run url-parsing.go 
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
```

### 🧾 Çıktı

URL ayrıştırma programımızı çalıştırmak, çıkardığımız tüm farklı parçaları gösterir.

## ⏭️ Sonraki Örnek: SHA256 Hash’leri

