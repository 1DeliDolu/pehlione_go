
## 🚀 Giriş

Bu örnekte, Go’da temel bir HTTP sunucusunun nasıl oluşturulacağını öğreneceksiniz. Önce HTTP sunucumuzun neleri yapabilmesi gerektiğinden bahsedelim. Temel bir HTTP sunucusunun yerine getirmesi gereken birkaç önemli görev vardır.

* **Dinamik istekleri işleme:** Web sitesini gezen, hesaplarına giriş yapan veya resim yükleyen kullanıcıların gönderdiği gelen istekleri işler.
* **Statik varlıkları sunma:** Kullanıcıya dinamik bir deneyim sağlamak için tarayıcılara JavaScript, CSS ve görseller sunar.
* **Bağlantıları kabul etme:** İnternetten bağlantı kabul edebilmek için HTTP Sunucusu belirli bir portu dinlemelidir.

---

## ⚙️ Dinamik İstekleri İşleme

`net/http` paketi, istekleri kabul etmek ve dinamik olarak işlemek için gereken tüm yardımcı araçları içerir. `http.HandleFunc` fonksiyonu ile yeni bir *handler* kaydedebiliriz. İlk parametresi eşleştirilecek bir yolu (*path*), ikinci parametresi ise çalıştırılacak fonksiyonu alır. Bu örnekte: Birisi web sitenizi (`http://example.com/`) ziyaret ettiğinde, güzel bir mesajla karşılanacaktır.

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```

Dinamik kısım için, `http.Request` isteğe ve parametrelerine dair tüm bilgileri içerir. GET parametrelerini `r.URL.Query().Get("token")` ile veya POST parametrelerini (HTML formundan gelen alanları) `r.FormValue("email")` ile okuyabilirsiniz.

---

## 🧩 Statik Varlıkları Sunma

JavaScript, CSS ve görseller gibi statik varlıkları sunmak için, yerleşik `http.FileServer` kullanırız ve bunu bir URL yoluna yönlendiririz. Dosya sunucusunun doğru çalışabilmesi için dosyaları nereden sunacağını bilmesi gerekir. Bunu şu şekilde yapabiliriz:

```go
fs := http.FileServer(http.Dir("static/"))
```

Dosya sunucumuz hazır olduğunda, tıpkı dinamik isteklerde yaptığımız gibi bir URL yolunu ona yönlendirmemiz yeterlidir. Dikkat edilmesi gereken bir nokta: Dosyaları doğru sunabilmek için URL yolunun bir kısmını kaldırmamız gerekir. Genellikle bu kısım, dosyaların bulunduğu dizinin adıdır.

```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

---

## 🔌 Bağlantıları Kabul Etme

Temel HTTP sunucumuzu tamamlamak için son adım, internetten bağlantı kabul edebilmek adına bir portu dinlemektir. Tahmin edebileceğiniz gibi Go’da yerleşik bir HTTP sunucusu da vardır ve oldukça hızlı bir şekilde başlatabiliriz. Başladıktan sonra HTTP sunucunuzu tarayıcıda görüntüleyebilirsiniz.

```go
http.ListenAndServe(":80", nil)
```

---

## 📄 Kod (Kopyala/Yapıştır İçin)

Bu, bu örnekte öğrendiklerinizi denemek için kullanabileceğiniz tam koddur.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":80", nil)
}
```

