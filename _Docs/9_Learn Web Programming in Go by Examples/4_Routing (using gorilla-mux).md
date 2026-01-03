
## 🧭 Routing (gorilla/mux kullanarak)

## 🚀 Giriş

Go’nun `net/http` paketi, HTTP protokolü için pek çok işlevsellik sağlar. Ancak çok karmaşık istek yönlendirmeyi (*request routing*) pek iyi yapamaz; örneğin bir istek URL’sini tek tek parametrelere bölme gibi. Neyse ki bunun için çok popüler bir paket vardır ve Go topluluğunda iyi kod kalitesiyle bilinir. Bu örnekte, `gorilla/mux` paketini kullanarak isimlendirilmiş parametrelerle rotalar, GET/POST handler’ları ve alan adı (*domain*) kısıtlamaları oluşturmayı göreceksiniz.

---

## 📦 gorilla/mux Paketini Yükleme

`gorilla/mux`, Go’nun varsayılan HTTP router’ına uyum sağlayan bir pakettir. Web uygulamaları yazarken üretkenliği artıran birçok özellikle gelir. Ayrıca Go’nun varsayılan istek handler imzasıyla uyumludur: `func (w http.ResponseWriter, r *http.Request)`; bu sayede paket, *middleware* veya mevcut uygulamalar gibi diğer HTTP kütüphaneleriyle birlikte karıştırılıp kullanılabilir. Paketi GitHub üzerinden yüklemek için `go get` komutunu şu şekilde kullanın:

```bash
go get -u github.com/gorilla/mux
```

---

## 🧱 Yeni Bir Router Oluşturma

Önce yeni bir istek router’ı oluşturun. Router, web uygulamanızın ana router’ıdır ve daha sonra sunucuya parametre olarak verilecektir. Tüm HTTP bağlantılarını alır ve sizin kaydedeceğiniz istek handler’larına yönlendirir. Yeni bir router’ı şu şekilde oluşturabilirsiniz:

```go
r := mux.NewRouter()
```

---

## 🧩 Bir Request Handler Kaydetme

Yeni bir router oluşturduktan sonra, her zamanki gibi request handler’ları kaydedebilirsiniz. Tek fark, `http.HandleFunc(...)` çağırmak yerine router’ınız üzerinde `HandleFunc` çağırmanızdır; yani şöyle: `r.HandleFunc(...)`.

---

## 🧷 URL Parametreleri

`gorilla/mux` Router’ın en büyük gücü, istek URL’sinden parçaları (*segment*) çıkarabilmesidir. Örnek olarak, uygulamanızda şöyle bir URL olsun:

`/books/go-programming-blueprint/page/10`

Bu URL’de iki dinamik segment vardır:

* Kitap başlığı *slug*’ı (`go-programming-blueprint`)
* Sayfa (`10`)

Yukarıda belirtilen URL’yi eşleştirecek bir request handler yazmak için, dinamik segmentleri URL kalıbınızda (*pattern*) yer tutucularla değiştirirsiniz:

```go
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    // get the book
    // navigate to the page
})
```

Son adım, bu segmentlerden veriyi almaktır. Paket, `http.Request`’i parametre alan ve segmentleri bir *map* olarak döndüren `mux.Vars(r)` fonksiyonunu sağlar.

```go
func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    vars["title"] // the book title slug
    vars["page"] // the page
}
```

---

## 🔧 HTTP Sunucusunun Router’ını Ayarlama

`http.ListenAndServe(":80", nil)` içindeki `nil`’in ne anlama geldiğini hiç merak ettiniz mi? Bu, HTTP sunucusunun ana router parametresidir. Varsayılan olarak `nil`’dir; bu da `net/http` paketinin varsayılan router’ını kullanmak anlamına gelir. Kendi router’ınızı kullanmak için `nil` yerine router değişkeniniz `r`’yi koyun.

```go
http.ListenAndServe(":80", r)
```

---

## 📄 Kod (Kopyala/Yapıştır İçin)

Bu, bu örnekte öğrendiklerinizi denemek için kullanabileceğiniz tam koddur.

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}
```

---

## 🧰 gorilla/mux Router Özellikleri

## 🧪 Methods

Request handler’ı belirli HTTP method’larıyla sınırlandırın.

```go
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
```

---

## 🌐 Hostnames ve Subdomains

Request handler’ı belirli alan adları (*hostname*) veya alt alan adları (*subdomain*) ile sınırlandırın.

```go
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")
```

---

## 🔒 Schemes

Request handler’ı `http/https` ile sınırlandırın.

```go
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
```

---

## 🧱 Path Prefixes ve Subrouters

Request handler’ı belirli path prefix’leri ile sınırlandırın.

```go
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
```

