
## 🧩 Multitemplate

Gin varsayılan olarak yalnızca **tek bir** `html.Template` kullanmaya izin verir. Go 1.6 *block template* gibi özellikleri kullanmak için bir **multitemplate render**’a bakın. Bu, birden fazla `*template.Template` (yani tek template yerine **multi template**) desteği sağlayan özel bir HTML render’dır.

## ▶️ Kullanım

Kullanmaya başlayın.

## 📦 İndir ve Kur

```bash
go get github.com/gin-contrib/multitemplate
```

## 📥 Projene Import Et

```go
import "github.com/gin-contrib/multitemplate"
```

## 🧪 Basit Örnek

`example/simple/example.go` örneğine bakın.

```go
package main

import (
  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
  r := multitemplate.NewRenderer()
  r.AddFromFiles("index", "templates/base.html", "templates/index.html")
  r.AddFromFiles("article", "templates/base.html", "templates/index.html", "templates/article.html")
  return r
}

func main() {
  router := gin.Default()
  router.HTMLRender = createMyRender()
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index", gin.H{
      "title": "Html5 Template Engine",
    })
  })
  router.GET("/article", func(c *gin.Context) {
    c.HTML(200, "article", gin.H{
      "title": "Html5 Article Engine",
    })
  })
  router.Run(":8080")
}
```

## 🏗️ Gelişmiş Örnek

`html/template` *inheritance*’ını yaklaşık olarak modelleme.

`example/advanced/example.go` örneğine bakın.

```go
package main

import (
  "path/filepath"

  "github.com/gin-contrib/multitemplate"
  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.HTMLRender = loadTemplates("./templates")
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
      "title": "Welcome!",
    })
  })
  router.GET("/article", func(c *gin.Context) {
    c.HTML(200, "article.html", gin.H{
      "title": "Html5 Article Engine",
    })
  })

  router.Run(":8080")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
  r := multitemplate.NewRenderer()

  layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
  if err != nil {
    panic(err.Error())
  }

  includes, err := filepath.Glob(templatesDir + "/includes/*.html")
  if err != nil {
    panic(err.Error())
  }

  // layouts/ ve includes/ dizinlerinden template haritamızı oluştur
  for _, include := range includes {
    layoutCopy := make([]string, len(layouts))
    copy(layoutCopy, layouts)
    files := append(layoutCopy, include)
    r.AddFromFiles(filepath.Base(include), files...)
  }
  return r
}
```

