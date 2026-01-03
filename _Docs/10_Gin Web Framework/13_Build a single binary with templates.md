
## 📦 Template’lerle Tek Bir Binary Build Etmek

### 🧰 Üçüncü Parti Paketi Kullanma

`go-assets` kullanarak, template’leri de içinde barındıran tek bir binary içine bir sunucu build etmek için üçüncü parti paketi kullanabilirsiniz.

```go
func main() {
  r := gin.New()

  t, err := loadTemplate()
  if err != nil {
    panic(err)
  }
  router.SetHTMLTemplate(t)

  router.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "/html/index.tmpl", nil)
  })
  router.Run(":8080")
}

// loadTemplate, go-assets-builder tarafından embed edilen template’leri yükler
func loadTemplate() (*template.Template, error) {
  t := template.New("")
  for name, file := range Assets.Files {
    if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
      continue
    }
    h, err := ioutil.ReadAll(file)
    if err != nil {
      return nil, err
    }
    t, err = t.New(name).Parse(string(h))
    if err != nil {
      return nil, err
    }
  }
  return t, nil
}
```

Tam bir örnek için `assets-in-binary/example01` dizinine bakın.

