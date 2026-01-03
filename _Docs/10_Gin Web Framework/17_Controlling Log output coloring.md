
## 🎨 Log Çıktısının Renklendirmesini Kontrol Etme

Varsayılan olarak, konsoldaki log çıktıları algılanan **TTY** durumuna göre renklendirilmelidir.

## 🚫 Logları Asla Renklendirme

```go
func main() {
    // Logların rengini devre dışı bırak
    gin.DisableConsoleColor()

    // Varsayılan middleware ile bir gin router oluşturur:
    // logger ve recovery (çökmesiz) middleware
    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
```

## ✅ Logları Her Zaman Renklendir

```go
func main() {
    // Logların rengini zorla
    gin.ForceConsoleColor()

    // Varsayılan middleware ile bir gin router oluşturur:
    // logger ve recovery (çökmesiz) middleware
    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
```

