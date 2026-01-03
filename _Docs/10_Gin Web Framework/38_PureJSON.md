
## 🧼 PureJSON

Normalde JSON, özel HTML karakterlerini unicode entity’leri ile değiştirir; örneğin `<`, `\u003c` olur. Bu tür karakterleri *literal* (aynı haliyle) encode etmek istiyorsanız bunun yerine `PureJSON` kullanabilirsiniz. Bu özellik Go **1.6 ve altı** sürümlerde mevcut değildir.

```go
func main() {
  router := gin.Default()

  // Unicode entity’lerini döndürür
  router.GET("/json", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // Literal karakterleri döndürür
  router.GET("/purejson", func(c *gin.Context) {
    c.PureJSON(200, gin.H{
      "html": "<b>Hello, world!</b>",
    })
  })

  // PureJSON yanıtı ve status code ile erken abort et (v1.11+)
  router.GET("/abort_purejson", func(c *gin.Context) {
    c.AbortWithStatusPureJSON(403, gin.H{"error": "forbidden"})
  })

  // 0.0.0.0:8080 üzerinde dinle ve sun
  router.Run(":8080")
}
```

## ✏️ Sayfayı Düzenle

