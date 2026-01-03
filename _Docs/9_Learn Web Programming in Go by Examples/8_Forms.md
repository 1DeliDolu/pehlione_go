
## 🧾 Formlar

Bu örnek, bir *iletişim formunu* simüle etmeyi ve mesajı bir **struct** içine parse etmeyi gösterir.

```go
// forms.go
package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // details ile bir şeyler yap
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}
```

```html
<!-- forms.html -->
{{if .Success}}
    <h1>Mesajınız için teşekkürler!</h1>
{{else}}
    <h1>İletişim</h1>
    <form method="POST">
        <label>E-posta:</label><br />
        <input type="text" name="email"><br />
        <label>Konu:</label><br />
        <input type="text" name="subject"><br />
        <label>Mesaj:</label><br />
        <textarea name="message"></textarea><br />
        <input type="submit">
    </form>
{{end}}
```

