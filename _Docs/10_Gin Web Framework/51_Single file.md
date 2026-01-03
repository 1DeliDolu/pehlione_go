
## 📎 Tek Dosya Yükleme

### 🧷 Notlar

* Bu bölüm *issue #774* ve detaylı örnek kodu referans alır.
* `file.Filename` **güvenilmemelidir**. *MDN’deki Content-Disposition* ve *#1693* konusuna bakın.
* Dosya adı her zaman opsiyoneldir ve uygulama tarafından körlemesine kullanılmamalıdır: **path bilgisi temizlenmeli** ve sunucu dosya sistemi kurallarına göre **dönüşüm yapılmalıdır**.

```go
func main() {
  router := gin.Default()
  // Multipart form’lar için daha düşük bir bellek limiti ayarla (varsayılan 32 MiB)
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // tek dosya
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    // Dosyayı belirtilen hedefe yükle.
    c.SaveUploadedFile(file, "./files/" + file.Filename)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
  })
  router.Run(":8080")
}
```

## 🧪 curl ile Nasıl Gönderilir

```bash
curl -X POST http://localhost:8080/upload \
  -F "file=@/Users/appleboy/test.zip" \
  -H "Content-Type: multipart/form-data"
```

