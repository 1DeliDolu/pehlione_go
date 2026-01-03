
## 🧵 Middleware İçinde Goroutine Kullanımı

Bir middleware veya handler içinde yeni Goroutine’ler başlatırken, Goroutine içinde **orijinal context’i** kullanmamalısınız; onun yerine **salt-okunur (read-only) bir kopya** kullanmalısınız.

```go
func main() {
  router := gin.Default()

  router.GET("/long_async", func(c *gin.Context) {
    // goroutine içinde kullanılmak üzere kopya oluştur
    cCp := c.Copy()
    go func() {
      // time.Sleep() ile uzun bir görevi simüle et. 5 saniye
      time.Sleep(5 * time.Second)

      // kopyalanmış context "cCp" kullanılıyor, ÖNEMLİ
      log.Println("Done! in path " + cCp.Request.URL.Path)
    }()
  })

  router.GET("/long_sync", func(c *gin.Context) {
    // time.Sleep() ile uzun bir görevi simüle et. 5 saniye
    time.Sleep(5 * time.Second)

    // goroutine KULLANMADIĞIMIZ için context'i kopyalamamız gerekmez
    log.Println("Done! in path " + c.Request.URL.Path)
  })

  // 0.0.0.0:8080 üzerinde dinle ve sun
  router.Run(":8080")
}
```

