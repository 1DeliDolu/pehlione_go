
## 📑 Birden Fazla Dosya Yükleme

### 🧩 Detaylı Örnek Kod

```go
func main() {
  router := gin.Default()
  // Multipart form’lar için daha düşük bir bellek limiti ayarla (varsayılan 32 MiB)
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["files"]

    for _, file := range files {
      log.Println(file.Filename)

      // Dosyayı belirtilen hedefe yükle.
      c.SaveUploadedFile(file, "./files/" + file.Filename)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
  router.Run(":8080")
}
```

### 🧪 curl ile Nasıl Gönderilir

```bash
curl -X POST http://localhost:8080/upload \
  -F "files=@/Users/appleboy/test1.zip" \
  -F "files=@/Users/appleboy/test2.zip" \
  -H "Content-Type: multipart/form-data"
```

