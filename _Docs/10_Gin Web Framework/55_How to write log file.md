
## 📝 Log Dosyası Nasıl Yazılır

```go
func main() {
    // Konsol rengini devre dışı bırak, logları dosyaya yazarken konsol rengine ihtiyacın yoktur.
    gin.DisableConsoleColor()

    // Logları bir dosyaya yaz.
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // Logları aynı anda hem dosyaya hem konsola yazman gerekiyorsa aşağıdaki kodu kullan.
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    router.Run(":8080")
}
```

