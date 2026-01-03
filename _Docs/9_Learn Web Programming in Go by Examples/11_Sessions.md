
## 🍪 Oturumlar

Bu örnek, Go’da popüler **`gorilla/sessions`** paketiyle *session cookie*’lerinde veri saklamayı gösterir.

Cookie’ler, kullanıcının tarayıcısında saklanan küçük veri parçalarıdır ve her istekte sunucumuza gönderilir. Bunların içinde örneğin bir kullanıcının web sitemize giriş yapıp yapmadığını saklayabilir ve (sistemimizde) gerçekte kim olduğunu anlayabiliriz.

Bu örnekte yalnızca kimliği doğrulanmış kullanıcıların **`/secret`** sayfasındaki gizli mesajı görmesine izin vereceğiz. Buna erişmek için önce **`/login`** sayfasını ziyaret edip geçerli bir session cookie alması gerekir; bu işlem onu giriş yapmış olarak işaretler. Ayrıca **`/logout`** sayfasını ziyaret ederek gizli mesaja erişimini iptal edebilir.

```go
// sessions.go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/sessions"
)

var (
    // key 16, 24 veya 32 bayt uzunluğunda olmalıdır (AES-128, AES-192 veya AES-256)
    key = []byte("super-secret-key")
    store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Kullanıcının kimliği doğrulanmış mı kontrol et
    if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }

    // Gizli mesajı yazdır
    fmt.Fprintln(w, "Kek bir yalan!")
}

func login(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Kimlik doğrulama burada yapılır
    // ...

    // Kullanıcıyı doğrulanmış olarak işaretle
    session.Values["authenticated"] = true
    session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "cookie-name")

    // Kullanıcının kimlik doğrulamasını iptal et
    session.Values["authenticated"] = false
    session.Save(r, w)
}

func main() {
    http.HandleFunc("/secret", secret)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)

    http.ListenAndServe(":8080", nil)
}
```

```bash
$ go run sessions.go
```

```bash
$ curl -s http://localhost:8080/secret
Forbidden
```

```bash
$ curl -s -I http://localhost:8080/login
Set-Cookie: cookie-name=MTQ4NzE5Mz...
```

```bash
$ curl -s --cookie "cookie-name=MTQ4NzE5Mz..." http://localhost:8080/secret
The cake is a lie!
```

