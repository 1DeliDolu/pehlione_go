
## 🧷 JSONP

Farklı bir domain’deki sunucudan veri istemek için **JSONP** kullanma. Sorgu parametresi `callback` mevcutsa, yanıt gövdesine callback ekler.

```go
func main() {
  router := gin.Default()

  router.GET("/JSONP?callback=x", func(c *gin.Context) {
    data := map[string]interface{}{
      "foo": "bar",
    }

    // GET /JSONP?callback=x
    // callback x'tir
    // Çıktı şu olur  :   x({\"foo\":\"bar\"})
    c.JSONP(http.StatusOK, data)
  })

  // 0.0.0.0:8080 üzerinde dinle ve sun
  router.Run(":8080")
}
```

