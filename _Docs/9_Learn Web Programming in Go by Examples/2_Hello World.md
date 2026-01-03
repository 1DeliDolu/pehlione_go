
## 🧭 Giriş

Go, *bataryası dahil* (batteries included) bir programlama dilidir ve içinde zaten yerleşik bir web sunucusu barındırır. Standart kütüphanedeki `net/http` paketi, HTTP protokolüyle ilgili tüm işlevleri içerir. Buna (diğer birçok şeyin yanı sıra) bir HTTP istemcisi ve bir HTTP sunucusu dahildir. Bu örnekte, tarayıcınızda görüntüleyebileceğiniz bir web sunucusu oluşturmanın ne kadar basit olduğunu göreceksiniz.

## 🧩 İstek İşleyicisi Kaydetme

Öncelikle, tarayıcılardan, HTTP istemcilerinden veya API isteklerinden gelen tüm HTTP bağlantılarını alan bir *Handler* oluşturun. Go’da bir handler, şu imzaya sahip bir fonksiyondur:

```go
func (w http.ResponseWriter, r *http.Request)
```

Fonksiyon iki parametre alır:

* `http.ResponseWriter`: `text/html` yanıtınızı yazdığınız yerdir.
* `http.Request`: URL veya header alanları gibi şeyler dahil olmak üzere bu HTTP isteğiyle ilgili tüm bilgileri içerir.

Varsayılan HTTP sunucusuna bir istek işleyicisi kaydetmek şu kadar basittir:

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

## 🔌 HTTP Bağlantılarını Dinleme

Yalnızca istek işleyicisi, dışarıdan gelen HTTP bağlantılarını tek başına kabul edemez. Bir HTTP sunucusunun, bağlantıları istek işleyicisine iletmek için bir port üzerinde dinlemesi gerekir. Port `80`, çoğu durumda HTTP trafiği için varsayılan port olduğu için bu sunucu da onu dinleyecektir.

Aşağıdaki kod Go’nun varsayılan HTTP sunucusunu başlatır ve `80` portunda bağlantıları dinler. Tarayıcınızla `http://localhost/` adresine giderek sunucunuzun isteğinizi işlediğini görebilirsiniz.

```go
http.ListenAndServe(":80", nil)
```

## 📋 Kodun Tamamı

Bu, bu örnekte öğrendiklerinizi denemek için kullanabileceğiniz tam koddur (kopyala/yapıştır):

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", nil)
}
```

