
## 🔐 Parola Hash’leme

### 🧂 Parola Hash’leme (bcrypt)

Bu örnek, **bcrypt** kullanarak parolaların nasıl hash’leneceğini gösterir. Bunun için Go’nun bcrypt kütüphanesini şu şekilde indirmemiz gerekir:

```bash
$ go get golang.org/x/crypto/bcrypt
```

Bundan sonra yazdığımız her uygulama bu kütüphaneyi kullanabilir.

```go
// passwords.go
package main

import (
    "fmt"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
    password := "secret"
    hash, _ := HashPassword(password) // basitlik adına hata yok sayıldı

    fmt.Println("Parola:", password)
    fmt.Println("Hash:  ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Println("Eşleşme:", match)
}
```

```bash
$ go run passwords.go
Password: secret
Hash:     $2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK
Match:    true
```

