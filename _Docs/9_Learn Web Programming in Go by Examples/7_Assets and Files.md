
## 🗂️ Assets ve Dosyalar

Bu örnek, **CSS**, **JavaScript** veya **görseller** gibi *statik dosyaları* belirli bir dizinden nasıl servis edeceğinizi gösterir.

```go
// static-files.go
package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("assets/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}
```

```bash
$ tree assets/
assets/
└── css
    └── styles.css
```

```bash
$ go run static-files.go
```

```bash
$ curl -s http://localhost:8080/static/css/styles.css
```

```css
body {
    background-color: black;
}
```

