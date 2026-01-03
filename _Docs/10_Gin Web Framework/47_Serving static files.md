
## 🗂️ Statik Dosyaları Sunma

```go
func main() {
  router := gin.Default()
  router.Static("/assets", "./assets")
  router.StaticFS("/more_static", http.Dir("my_file_system"))
  router.StaticFile("/favicon.ico", "./resources/favicon.ico")

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

