
## 🧾 JSON

Bu örnek, **`encoding/json`** paketi kullanarak JSON verisini nasıl **encode** ve **decode** edeceğinizi gösterir.

```go
// json.go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Age       int    `json:"age"`
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%s %s %d yaşında!", user.Firstname, user.Lastname, user.Age)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        peter := User{
            Firstname: "John",
            Lastname:  "Doe",
            Age:       25,
        }

        json.NewEncoder(w).Encode(peter)
    })

    http.ListenAndServe(":8080", nil)
}
```

```bash
$ go run json.go
```

```bash
$ curl -s -XPOST -d'{"firstname":"Elon","lastname":"Musk","age":48}' http://localhost:8080/decode
Elon Musk is 48 years old!
```

```bash
$ curl -s http://localhost:8080/encode
{"firstname":"John","lastname":"Doe","age":25}
```

