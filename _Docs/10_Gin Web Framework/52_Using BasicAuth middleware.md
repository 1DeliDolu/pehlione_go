
## 🔐 BasicAuth Middleware Kullanımı

```go
// bazı özel verileri simüle et
var secrets = gin.H{
  "foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
  "austin": gin.H{"email": "austin@example.com", "phone": "666"},
  "lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func main() {
  router := gin.Default()

  // gin.BasicAuth() middleware kullanarak grup oluştur
  // gin.Accounts, map[string]string için bir kısayoldur
  authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
    "foo":    "bar",
    "austin": "1234",
    "lena":   "hello2",
    "manu":   "4321",
  }))

  // /admin/secrets endpoint’i
  // "localhost:8080/admin/secrets" isteği at
  authorized.GET("/secrets", func(c *gin.Context) {
    // kullanıcıyı al, BasicAuth middleware tarafından set edilmiştir
    user := c.MustGet(gin.AuthUserKey).(string)
    if secret, ok := secrets[user]; ok {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
    } else {
      c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
    }
  })

  // 0.0.0.0:8080 üzerinde dinle ve servis et
  router.Run(":8080")
}
```

