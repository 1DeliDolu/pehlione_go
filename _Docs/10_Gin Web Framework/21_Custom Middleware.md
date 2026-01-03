
## 🧩 Özel Middleware

```go
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    t := time.Now()

    // Örnek değişkeni ayarla
    c.Set("example", "12345")

    // istek öncesi

    c.Next()

    // istek sonrası
    latency := time.Since(t)
    log.Print(latency)

    // gönderdiğimiz status koduna eriş
    status := c.Writer.Status()
    log.Println(status)
  }
}

func main() {
  r := gin.New()
  r.Use(Logger())

  r.GET("/test", func(c *gin.Context) {
    example := c.MustGet("example").(string)

    // şunu yazdırır: "12345"
    log.Println(example)
  })

  // 0.0.0.0:8080 üzerinde dinle ve sun
  r.Run(":8080")
}
```

