
## 🧾 HTML Render Etme

Gin, HTML render etmek için `html/template` paketini kullanır. Kullanım şekli ve mevcut placeholder’lar dahil daha fazla bilgi için `text/template` dokümantasyonuna bakın.

HTML dosyalarını yüklemek için `LoadHTMLGlob()` veya `LoadHTMLFiles()` kullanın.

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/*")
  //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
  router.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
      "title": "Main website",
    })
  })
  router.Run(":8080")
}
```

## 📄 templates/index.tmpl

```html
<html>
  <h1>
    {{ .title }}
  </h1>
</html>
```

## 🗂️ Farklı Dizinlerde Aynı İsimli Template Kullanımı

```go
func main() {
  router := gin.Default()
  router.LoadHTMLGlob("templates/**/*")
  router.GET("/posts/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
      "title": "Posts",
    })
  })
  router.GET("/users/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
      "title": "Users",
    })
  })
  router.Run(":8080")
}
```

> Not: HTML template’inizi `{{define <template-path>}}` `{{end}}` bloğu içine alın ve template dosyanızı göreli yol olan `<template-path>` ile tanımlayın. Aksi halde GIN template dosyalarını doğru şekilde parse edemez.

## 📄 templates/posts/index.tmpl

```html
{{ define "posts/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>Using posts/index.tmpl</p>
</html>
{{ end }}
```

## 📄 templates/users/index.tmpl

```html
{{ define "users/index.tmpl" }}
<html><h1>
  {{ .title }}
</h1>
<p>Using users/index.tmpl</p>
</html>
{{ end }}
```

## 🧩 http.FileSystem Üzerinden Template Yükleme (v1.11+)

Template’leriniz gömülü (*embedded*) ise veya bir `http.FileSystem` tarafından sağlanıyorsa `LoadHTMLFS` kullanın:

```go
import (
  "embed"
  "io/fs"
  "net/http"
  "github.com/gin-gonic/gin"
)

//go:embed templates
var tmplFS embed.FS

func main() {
  r := gin.Default()
  sub, _ := fs.Sub(tmplFS, "templates")
  r.LoadHTMLFS(http.FS(sub), "**/*.tmpl")
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "From FS"})
  })
  r.Run(":8080")
}
```

## 🛠️ Özel Template Renderer

Kendi HTML template renderer’ınızı da kullanabilirsiniz.

```go
import "html/template"

func main() {
  router := gin.Default()
  html := template.Must(template.ParseFiles("file1", "file2"))
  router.SetHTMLTemplate(html)
  router.Run(":8080")
}
```

## 🧷 Özel Delimiter’lar

Özel delimiter’lar (*delims*) kullanabilirsiniz.

```go
  router := gin.Default()
  router.Delims("{[{", "}]}")
  router.LoadHTMLGlob("/path/to/templates")
```

## 🧠 Özel Template Fonksiyonları

Detaylı örnek koda bakın.

## 📄 main.go

```go
import (
    "fmt"
    "html/template"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
    year, month, day := t.Date()
    return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
    router := gin.Default()
    router.Delims("{[{", "}]}")
    router.SetFuncMap(template.FuncMap{
        "formatAsDate": formatAsDate,
    })
    router.LoadHTMLFiles("./testdata/template/raw.tmpl")

    router.GET("/raw", func(c *gin.Context) {
        c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
            "now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
        })
    })

    router.Run(":8080")
}
```

## 📄 raw.tmpl

Terminal penceresi

```text
Date: {[{.now | formatAsDate}]}
```

## ✅ Sonuç

Terminal penceresi

```text
Date: 2017/07/01
```

